package alteon

import (
  "context"
  //"strconv"
  "encoding/json"
  ac "github.com/irekromaniuk/alteon-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRealServer() *schema.Resource {
	return &schema.Resource{
	  ReadContext: dataSourceRealServerRead,
	  Schema: map[string]*schema.Schema{
		"index": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		  },
		"items": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ipaddr": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"weight": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"maxconns": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"timeout": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},  
					"pinginterval": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"failretry": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"succretry": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"state": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"deletestatus": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"type": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},  
					"name": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
						},
					"addurl": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"remurl": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"cookie": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},  
					"excludestr": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"submac": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"idsport": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"ipver": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"ipv6addr": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"nxtrportidx": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},  
					"nxtbuddyidx": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"llbtype": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"copy": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"portsingress": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
						},
					"portsegress": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},  
					"addportsingress": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"remportsingress": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"addportsegress": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"remportsegress": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},  
					"vlaningress": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"vlanegress": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"egressif": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
					"sectype": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
						},
					"ingressif": &schema.Schema{
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
	 	},
	},
  }
}

  func dataSourceRealServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "HostUrl and Token",
		Detail:   "HostURL:" + c.HostURL + " Token:" + c.Token,
	  })
	// return diags 
	RealServerID := d.Get("index").(string)

	RealServer, err := c.GetRealServer(RealServerID)
	if err != nil {
		return diag.FromErr(err)
	}
	prettyJSON, _ := json.MarshalIndent(RealServer, "", "    ")
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "RealServer",
		Detail:   string(prettyJSON),
	  })
	// return diags 
	RealServerItems, diags := flattenRealServerItemsData(&RealServer.Items)
	prettyJSON, _ = json.MarshalIndent(RealServerItems[0], "", "    ")
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "RealServerItems",
		Detail:    string(prettyJSON),
	  })
	// return diags   
	if err := d.Set("items", RealServerItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(RealServerID)

	return diags
  }

  func flattenRealServerItemsData(RealServerItems *[]ac.RealServerItem) ([]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	
	   
	if RealServerItems != nil {
	  rss := make([]interface{}, len(*RealServerItems), len(*RealServerItems))
  
	  for i, RealServerItem := range *RealServerItems {
		rsi := make(map[string]interface{})
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "flatten IpAddr",
			Detail:   RealServerItem.IpAddr,
		  })
		rsi["ipaddr"] = RealServerItem.IpAddr
		rsi["weight"] = RealServerItem.Weight
		rsi["maxconns"] = RealServerItem.MaxConns
		rsi["timeout"] = RealServerItem.TimeOut
		rsi["pinginterval"] = RealServerItem.PingInterval
		rsi["failretry"] = RealServerItem.FailRetry
		rsi["succretry"] = RealServerItem.SuccRetry
  
		rss[i] = rsi
	  }
  
	  return rss, diags
	}
  
	return make([]interface{}, 0), diags
  }
  