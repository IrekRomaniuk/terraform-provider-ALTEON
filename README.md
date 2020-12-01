![release](https://github.com/IrekRomaniuk/terraform-provider-ALTEON/workflows/release/badge.svg)
# Alteon Provider

This provider interacts with Radware Alteon VA (Standalone) API, tested with version 32.6.1.0 (and Terraform v.0.13). Provider uses [go client from the repo](https://github.com/IrekRomaniuk/alteon-client-go)

## Example Usage

Look for example of main.tf in the example directory

```

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

resource "alteon_real_server" "LabServer" {
  index="LabServer1"
  items {
      ipaddr="1.2.3.4"
      name="description"
    }
}

data "alteon_real_server" "LabServer" {
  index="LabServer1"
}

output "LabServer" {
  value = data.alteon_real_server.LabServer
}

```







