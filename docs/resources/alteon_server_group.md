# alteon_server_group Resource

Maintaines Alteon Server Group table SlbNewCfgEnhGroupTable.  

## Example Usage

```
resource "alteon_server_group" "LabServers" {
    index="LabServers"
    items {
        name="Group description"
        addserver="TestServer"
        healthcheckurl="tcp-443"
    }
}
```

## Argument Reference

- `Index` - The group alphanumeric index for which the information pertains.
- `AddServer` - The real server to be added to the group. When read, 0 is returned.
- `RemoveServer` - The real server to be removed from the group. When read, 0 is returned.
- `HealthCheckUrl` - The specific content which is examined during health checks. * The content depends on the type of health check.
- `Name` - The name of the real server group.
