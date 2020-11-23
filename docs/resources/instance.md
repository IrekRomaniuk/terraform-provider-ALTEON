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