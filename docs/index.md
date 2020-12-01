# alteon Provider

This provider interacts with Radware Alteon VA (Standalone, next-generation application delivery controller - ADC) API, tested with version 32.6.1.0 (and Terraform v.0.13). Provider uses [go client from the repo](https://github.com/IrekRomaniuk/alteon-client-go)

## Example Usage

```

terraform {
  required_providers {
    alteon = {
      versions = ["0.2"]
      source = "irekromaniuk/alteon"
    }
  }
}

provider "alteon" {
    username="admin"
    password=""
    uri="https://13.92.134.159:8443/config"
}


```

## Argument Reference

- `username` - alteon VA username (env "ALTEON_USERNAME")
- `password` - alteon VA password (env "ALTEON_PASSWORD")
- `uri` - alteon VA uri, i.e. "https://13.92.134.159:8443/config" (env "ALTEON_URI")
