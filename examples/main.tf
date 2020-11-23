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

/*resource "alteon_real_server" "LabServer" {
  index="LabServer1"
  items {
    	ipaddr="1.1.1.2"
      name="description2"
      //weight=1
      //timeout=2
      state=2
    }
}*/

resource "alteon_real_server" "LabServer2" {
  index="LabServer2"
  items {
    	ipaddr="2.2.2.2"
      name="description2"
      state=2
    }
}

/*resource "alteon_server_group" "LabServers" {
  
}*/