# terraform-provider-alteon Provider

This provider interacts with Radware Alteon VA (Standalone) API, tested with version 32.6.1.0 (and Terraform v.0.13)

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
The following env variables should be set (IP address is an example only):

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

# alteon_real_server Alteon Real Servers

Maintaines Alteon Real Server table SlbNewCfgEnhRealServerTable.  

## Example Usage

```
resource "alteon_real_server" "LabServer" {
  index="LabServer1"
  items {
      ipaddr="1.1.1.1"
      name="description"
    }
}
```

## Argument Reference

Index - The real server number
IpAddr - IP address of the real server identified by the instance of * slbRealServerIndex.
Weight - The server weight.
MaxConns - The maximum number of connections that are allowed.
TimeOut - The maximum number of minutes an inactive connection remains open.
PingInterval - The interval between keep-alive (ping) attempts in number of * seconds. Zero means disabling ping attempt.
FailRetry - The number of failed attempts to declare this server DOWN.
SuccRetry - The number of successful attempts to declare a server UP.
State - Enable or disable the server and remove the existing sessions using disabled-with-fastage option.
Name - The name of the real server.

## Argument Values

"IpAddr": IpAddress
"Weight": integer
"MaxConns": integer
"TimeOut": integer
"PingInterval": integer
"FailRetry": integer
"SuccRetry": integer
"State": integer // {2=ENABLED, 3=DISABLED, 4=DISABLED_WITH_FASTAGE}
"Name": string

### Alteon tables implemented:

- SlbNewCfgEnhRealServerTable: Real Servers

### Alteon tables to be implemented:

- SlbNewCfgEnhGroupTable: Server Groups
- SlbNewCfgEnhVirtServerTable: Virtual Servers
- VrrpNewCfgVirtRtrTable: Virtual Routers


#### Notes:

How to compile/install provider:

```
go mod terraform-provider-alteon
go mod tidy
go mod vendor 
```
or just

```
make update-go-deps
make install
```

Provider uses [go client from the repo](https://github.com/IrekRomaniuk/alteon-client-go)



