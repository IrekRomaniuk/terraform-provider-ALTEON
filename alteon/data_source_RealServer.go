package alteon

import (
  "context"
  "strconv"

  ac "github.com/irekromaniuk/alteon-client-go
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRealServer() *schema.Resource {
	return &schema.Resource{
	  ReadContext: dataSourceOrderRead,
	  Schema: map[string]*schema.Schema{
		"IpAddr": &schema.Schema{
		  Type:     schema.TypeString,
		  Required: true,
		},
		"Weight": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		  },
		"MaxConns": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		},
		"TimeOut": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},  
		"PingInterval": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		  },
		"FailRetry": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		},
		"SuccRetry": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"State": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		  },
		"DeleteStatus": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		},
		"Type": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},  
		"Name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		  },
		"AddUrl": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		},
		"RemUrl": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"Cookie": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},  
		"ExcludeStr": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		  },
		"Submac": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"Idsport": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"IPVer": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		  },
		"Ipv6Addr": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		},
		"NxtRportIdx": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},  
		"NxtBuddyIdx": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		  },
		"LLBType": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"Copy": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"PortsIngress": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			},
		"PortsEgress": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},  
		"AddPortsIngress": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
			},
		"RemPortsIngress": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"AddPortsEgress": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"RemPortsEgress": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},  
		"VlanIngress": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
			},
		"VlanEgress": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		},
		"EgressIf": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"SecType": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
			},
		"IngressIf": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
		},
	  },
	}
  }