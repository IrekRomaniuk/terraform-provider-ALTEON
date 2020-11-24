# alteon Provider

This provider interacts with Radware Alteon VA (Standalone, next-generation application delivery controller - ADC) API, tested with version 32.6.1.0 (and Terraform v.0.13). Provider uses [go client from the repo](https://github.com/IrekRomaniuk/alteon-client-go)

## Example Usage

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
    username="admin"
    password=""
    uri="https://13.92.134.158:8443/config"
}


```

## Argument Reference

- `username` - alteon VA username
- `password` - alteon VA password
- `uri` - alteon VA uri, i.e. "https://13.92.134.158:8443/config"