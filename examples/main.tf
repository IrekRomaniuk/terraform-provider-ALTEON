terraform {
  required_providers {
    alteon = {
      versions = ["0.1"]
      source = "github.com/irekromaniuk/alteon"
    }
  }
}

provider "alteon" {}

data "real_server" "LabServer" {
  Index="LabServer1"
}

/*output "LabServer" {
  value = data.real_server.LabServer
}*/