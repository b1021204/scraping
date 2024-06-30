variable "hostname" { default = "tf-vm" }
variable "domain" { default = "illumination-as-code.com" }
variable "machine_num" { default = 3 }
variable "cpu" {
  type    = list(number)
  default = [1, 1, 1]
}

terraform {
  required_version = ">= 0.13"
  required_providers {
    libvirt = {
      version = ">= 0.7.6"
      source  = "dmacvicar/libvirt"
    }
  }
}

provider "libvirt" {
  uri = "qemu:///system"
}

resource "libvirt_volume" "os_image" {
  count = var.machine_num
  name   = "${var.hostname}${count.index}-vm-os_image"
  pool   = "default"
  source = "./jammy-server-cloudimg-amd64.img"
  format = "qcow2"
}

resource "libvirt_cloudinit_disk" "commoninit" {
  count = var.machine_num
  name           = "${var.hostname}${count.index}-init.iso"
  pool           = "default"
  user_data      = <<-EOF
  #cloud-config
  hostname: ${var.hostname}${count.index}
  fqdn: ${var.hostname}${count.index}.${var.domain}
  manage_etc_hosts: true
  users:
    - name: ubuntu
      sudo: ALL=(ALL) NOPASSWD:ALL
      groups: users, admin
      home: /home/ubuntu
      shell: /bin/bash
      lock_passwd: false
      ssh-authorized-keys:
        - ${file("id_ed25519.pub")}
  # only cert auth via ssh (console access can still login)
  ssh_pwauth: false
  disable_root: false
  chpasswd:
    list: |
       ubuntu:linux
    expire: False
  packages:
   - qemu-guest-agent
  EOF
  network_config = <<-EOF
  version: 2
  ethernets:
    ens3:
      addresses:
      - "192.168.${count.index+2}.22/24"
      gateway4: 192.168.${count.index+2}.1
      nameservers:
        addresses:
          - 8.8.8.8
          - 8.8.4.4
  EOF
}

resource "libvirt_domain" "vm" {
  for_each = { for idx, cpu in var.cpu : idx => cpu if cpu > 0 }

  name   = "${var.hostname}-${each.key+1}"
  memory = 2048
  vcpu   = each.value

  disk {
    volume_id = libvirt_volume.os_image[each.key].id
  }

  network_interface {
    bridge = "br${each.key+2}"
    addresses = ["192.168.${each.key+2}.22"]
  }

  cloudinit = libvirt_cloudinit_disk.commoninit[each.key].id

  console {
    type        = "pty"
    target_port = "0"
    target_type = "serial"
  }

  graphics {
    type        = "spice"
    listen_type = "address"
    autoport    = "true"
  }

  provisioner "file" {
    source      = "compute/compute_pi"
    destination = "/tmp/compute_pi"
    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = file("id_ed25519")
      host        = format("192.168.%d.22", each.key + 2)
    }
  }

  provisioner "file" {
    source      = format("compute/pi%d.service", each.key + 1)
    destination = format("/tmp/pi%d.service", each.key + 1)
    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = file("id_ed25519")
      host        = format("192.168.%d.22", each.key + 2)
    }
  }

  provisioner "remote-exec" {
    inline = [
      format("sudo mv /tmp/pi%d.service /etc/systemd/system/pi%d.service", each.key + 1, each.key + 1),
      "sudo chmod 777 /tmp/compute_pi",
      "sudo systemctl daemon-reload",
      format("sudo systemctl start pi%d.service", each.key + 1),
    ]
    connection {
      type        = "ssh"
      user        = "ubuntu"
      private_key = file("id_ed25519")
      host        = format("192.168.%d.22", each.key + 2)
    }
  }
}