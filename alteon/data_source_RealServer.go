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
			Computed: true ,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ipaddr": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"weight": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"maxconns": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"timeout": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},  
					"pinginterval": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"failretry": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"succretry": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"state": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"deletestatus": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"type": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},  
					"name": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						},
					"addurl": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"remurl": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"cookie": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},  
					"excludestr": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"submac": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"idsport": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"ipver": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"ipv6addr": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
					"nxtrportidx": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},  
					"nxtbuddyidx": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"llbtype": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"copy": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
					"portsingress": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
						},
					"portsegress": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},  
					"addportsingress": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"remportsingress": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"addportsegress": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"remportsegress": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},  
					"vlaningress": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"vlanegress": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"egressif": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
					},
					"sectype": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
						},
					"ingressif": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true,
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
	/*diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "HostUrl and Token",
		Detail:   "HostURL:" + c.HostURL + " Token:" + c.Token,
	  })*/
 
	RealServerID := d.Get("index").(string)
	Table  := "SlbNewCfgEnhRealServerTable"

	RealServer, err := c.GetItem(Table, RealServerID)
	if err != nil {
		return diag.FromErr(err)
	}
	Items := RealServer[Table]
	helper, err := json.Marshal(Items)
	if err != nil {
		return diag.FromErr(err)
	}
	var Item []ac.RealServerItem
	json.Unmarshal(helper, &Item)
	RealServerItems := flattenRealServerItemsData(&Item) //&RealServer.Items
	  
	if err := d.Set("items", RealServerItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(RealServerID)

	return diags
  }

  func flattenRealServerItemsData(RealServerItems *[]ac.RealServerItem) ([]interface{}) {
	 
	if RealServerItems != nil {
	  rss := make([]interface{}, len(*RealServerItems), len(*RealServerItems))
  
	  for i, RealServerItem := range *RealServerItems {
		rsi := make(map[string]interface{})
		
		rsi["ipaddr"] = RealServerItem.IpAddr
		rsi["name"] = RealServerItem.Name
		rsi["weight"] = RealServerItem.Weight
		rsi["maxconns"] = RealServerItem.MaxConns
		rsi["timeout"] = RealServerItem.TimeOut
		rsi["pinginterval"] = RealServerItem.PingInterval
		rsi["failretry"] = RealServerItem.FailRetry
		rsi["succretry"] = RealServerItem.SuccRetry
		rsi["state"] = RealServerItem.State
		
		rss[i] = rsi
	  }
  
	  return rss
	}
  
	return make([]interface{}, 0)
  }
  