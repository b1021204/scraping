terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.51.0"
    }
  }
}

provider "google" {
  project = "hakodate-ar-2023-dev"
  region  = "asia-northeast1"
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-nsysk"
  auto_create_subnetworks = false

}

resource "google_compute_address" "static" {
  name = "ipv4-address"
}

resource "google_compute_instance" "default" {
  name         = "nsysk"
  machine_type = "f1-micro"
  zone         = "asia-northeast1-b"
 tags = ["buildserver", "jenkins", "central", "terraformer", "http-server", "allow-ssh"]
  boot_disk {
    auto_delete = true
    initialize_params {
      image = "ubuntu-2004-focal-arm64-v20240607"
    }
  }
  network_interface {
    network = "default"
    access_config {

            nat_ip = google_compute_address.static.address

    }
  }

  /*
  provisioner "file" {
    source      = "script.sh"
    destination = "/tmp/script.sh"
  connection {
    type        = "ssh"
    user        = "root"
    private_key = file("key/terraform")
    host = google_compute_address.static.address
  }
  }*/
  /*
  provisioner "remote-exec" {
      connection {
    type        = "ssh"
    user        = "root"
    password =  ""
    private_key = file("key/terraform")
    host = google_compute_address.static.address
  }
    inline = [
      "apt update",
      "apt install sudo",
      "sudo apt-get install nginx -y",
    ]
  }*/
}

resource "google_compute_firewall" "allow_ssh" {
  name = "allow-ssh"
  network = google_compute_network.vpc_network.name



  target_tags   = ["allow-ssh"] // this targets our tagged VM
  source_ranges = ["0.0.0.0/0"]

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }
}

output "name" {
  value = google_compute_address.static.address
  
}



