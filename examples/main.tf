terraform {
  required_providers {
    alteon = {
      versions = ["0.1"]
      source = "github.com/irekromaniuk/alteon"
    }
  }
}

provider "alteon" {
}

/*data "alteon_real_server" "LabServer" {
  index="LabServer1"
}

output "LabServer" {
  value = data.alteon_real_server.LabServer
}*/

resource "alteon_real_server" "LabServer" {
  index="LabServer22"
  items {
    	ipaddr="10.2.3.4"
      name="description"
      weight=1
      timeout=2
    }
}