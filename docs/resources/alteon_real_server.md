# alteon_real_server Resource

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

- `Index` - String, the real server number, 

- `IpAddr` - IPAddress, IP address of the real server identified by the instance of * slbRealServerIndex.

- `Weight` - integer, the server weight.

- `MaxConns` - integer, the maximum number of connections that are allowed.

- `TimeOut` - integer, the maximum number of minutes an inactive connection remains open.

- `PingInterval` - integer, the interval between keep-alive (ping) attempts in number of * seconds. Zero means disabling ping attempt.

- `FailRetry` - integer, the number of failed attempts to declare this server DOWN.

- `SuccRetry` - integer, the number of successful attempts to declare a server UP.

- `State` - integer, enable or disable the server and remove the existing sessions using disabled-with-fastage option. // {2=ENABLED, 3=DISABLED, 4=DISABLED_WITH_FASTAGE}

- `Name` - string, the name of the real server, string
