terraform {
    required_providers {
        google = {
            source = "hashicorp/google"
            version = "4.51.0"
        }
    }
}

provider "google" {

    project = "hakodate-ar-2023-dev"
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-network"
}


resource "google_compute_instance" "vm_instance" {
    name = "terraform-instance"
    machine_type = "f1-micro"
    zone = "asia-northeast1-a"
   // zone = "asia-northeast1-a"
    boot_disk {
        initialize_params {
          image = "debian-cloud/debian-11"
      }
    }

    network_interface {
        network = google_compute_network.vpc_network.name
        access_config {
          
        }
    }
  
}