# terraform-provider-alteon (version 0.1)

Tested with Alteon VA (Standalone) Version 32.6.1.0 and Terraform v.0.13

### Alteon tables implemented:

- SlbNewCfgEnhRealServerTable: Real Servers

### Alteon tables to be implemented:

- SlbNewCfgEnhGroupTable: Server Groups
- SlbNewCfgEnhVirtServerTable: Virtual Servers
- VrrpNewCfgVirtRtrTable: Virtual Routers

### Resources implemented:

Example of main.tf

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
      ipaddr="1.1.1.1"
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
then set env variables

```
export ALTEON_USERNAME=admin
export ALTEON_PASSWORD=
export ALTEON_URI=https://13.92.134.158:8443/config

echo $ALTEON_USERNAME
echo $ALTEON_PASSWORD
echo $ALTEON_URI

cd examples
terraform init && terraform apply --auto-approve
terraform providers schema -json
```

### Notes:

```
go mod terraform-provider-alteon
go clean --modcache (optional !?)
go mod tidy
go mod vendor 
```
or just

```
make update-go-deps
make install
```



