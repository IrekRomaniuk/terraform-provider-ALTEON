# terraform-provider-alteon

Tested with Alteon VA (Standalone) Version 32.6.1.0

### Tables implemented:

- SlbNewCfgEnhRealServerTable: RealServerTable

### Resources implemented:

main.tf

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

data "alteon_real_server" "LabServer" {
  index="LabServer1"
}

output "LabServer" {
  value = data.alteon_real_server.LabServer
}

resource "alteon_real_server" "LabServer" {
  index="LabServer1"
  items {
      ipaddr="1.1.1.1"
      name="description"
    }
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
make make update-go-deps
make install
```



