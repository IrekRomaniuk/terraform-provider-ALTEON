# alteon_real_server Alteon Real Servers

Maintaines Alteon Real Server table SlbNewCfgEnhRealServerTable.  

## Example Usage

```
data "alteon_real_server" "LabServer" {
  index="LabServer1"
}

output "LabServer" {
  value = data.alteon_real_server.LabServer
}
```

## Argument Reference

Index - The real server number

