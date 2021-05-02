provider "google" {

  credentials= file("practicasa-312404-90169553f50b.json")
  project = "practicasa-312404"
  region = "us-west1"
  
}

resource "random_id" "instance_id" {
  byte_length = 8
  
}
//creando maquina virtual

resource "google_compute_instance" "default" {
  name= "practicasa-${random_id.instance_id.hex}"
  machine_type = "f1-micro"
  zone = "us-west1-a"

  boot_disk {
    initialize_params {
      image= "imagenpractica12sa201404104"
    }
  }

  metadata_startup_script = "sudo apt-get update; sudo apt-get install -yq build-essential python-pip rsync; pip install flask"

  network_interface {
    network = "default"

    access_config {
      //leer la ip publica
    }
  }

  metadata = {
    ssh-keys = "nicte:${file("~/.ssh/id_rsa.pub")}"

  }
}


resource "google_compute_firewall" "default" {
  name = "demoterraform-python5000-clone"
  network = "default"
  allow {
    protocol = "tcp"
    ports = ["5000"]
  }
}


output "ip"{
  value = google_compute_instance.default.network_interface.0.access_config.0.nat_ip
}