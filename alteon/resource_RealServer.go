package alteon

import (
  "context"
  //"strconv"
  //"fmt"
  //"encoding/json"
  "time"
  ac "github.com/irekromaniuk/alteon-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRealServer() *schema.Resource {
  return &schema.Resource{
    CreateContext: resourceRealServerCreate,
    ReadContext:   resourceRealServerRead,
    UpdateContext: resourceRealServerUpdate,
    DeleteContext: resourceRealServerDelete,
    Schema: map[string]*schema.Schema{
		"index": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		  },
		"last_updated": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		  },  
		"items": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"ipaddr": &schema.Schema{
						Type:     schema.TypeString,
						Required: true,
					},
					"weight": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 1,
						},
					"maxconns": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
					},
					"timeout": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 10,
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
						Default: 3,
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
						Default: 2,
					},  
					"excludestr": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 2,
						},
					"submac": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 2,
					},
					"idsport": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
					},
					"ipver": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 1,
						},
					"ipv6addr": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true, 
					},
					"nxtrportidx": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 1,
					},  
					"nxtbuddyidx": &schema.Schema{
						Type:     schema.TypeInt,
						Optional: true, 
						Default: 1,
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
						Default: 1,
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

func resourceRealServerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	RealServerID := d.Get("index").(string)
	items := d.Get("items").([]interface{})
	rss := []ac.RealServerItem{}

	
	for _, item := range items {
	  i := item.(map[string]interface{})
	  rsi := ac.RealServerItem{
		IpAddr: i["ipaddr"].(string),
		Name: i["name"].(string),
		Weight: i["weight"].(int),
		TimeOut: i["timeout"].(int),
		State: i["state"].(int),
		MaxConns: i["maxconns"].(int),
		PingInterval: i["pinginterval"].(int),
		FailRetry: i["failretry"].(int),
		SuccRetry: i["succretry"].(int),
	  }
	  rss = append(rss, rsi)
	}
	/*prettyJSON, _ := json.MarshalIndent(rss, "", "    ")
	  diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "rss",
		Detail:   fmt.Sprint(string(prettyJSON)),
	  })
	return diags*/ 
	rs, err := c.CreateRealServer(rss, RealServerID)
	/*diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "rs",
		Detail:   rs.Status,
	  })*/
	if err != nil {
	  return diag.FromErr(err)
	}
	if rs.Status == "ok" {
		d.SetId(RealServerID)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Status",
			Detail:   rs.Status, //add message or testerr
		  })
	}
	
	resourceRealServerRead(ctx, d, m)

	return diags
}

func resourceRealServerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	RealServerID := d.Id()
  
	RealServer, err := c.GetRealServer(RealServerID)
	if err != nil {
	  return diag.FromErr(err)
	}
  
	RealServerItems := flattenRealServerItems(&RealServer.Items)
	if err := d.Set("items", RealServerItems); err != nil {
	  return diag.FromErr(err)
	}
  
	return diags
}

func resourceRealServerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    c := m.(*ac.Client)

	RealServerID := d.Id()

	if d.HasChange("items") {
		items := d.Get("items").([]interface{})
		rss := []ac.RealServerItem{}

		for _, item := range items {
			i := item.(map[string]interface{})

			rsi := ac.RealServerItem{
				IpAddr: i["ipaddr"].(string),
				Name: i["name"].(string),
				Weight: i["weight"].(int),
				TimeOut: i["timeout"].(int),
				State: i["state"].(int),
				MaxConns: i["maxconns"].(int),
				PingInterval: i["pinginterval"].(int),
				FailRetry: i["failretry"].(int),
				SuccRetry: i["succretry"].(int),
			}
			rss = append(rss, rsi)
		}

		_, err := c.UpdateRealServer(rss, RealServerID)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceRealServerRead(ctx, d, m)
	
}

func resourceRealServerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  // Warning or errors can be collected in a slice type
  c := m.(*ac.Client)

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  RealServerID := d.Id()

  err := c.DeleteRealServer(RealServerID)
  if err != nil {
    return diag.FromErr(err)
  }

  // d.SetId("") is automatically called assuming delete returns no errors, but
  // it is added here for explicitness.
  d.SetId("")

  return diags
}

func flattenRealServerItems(RealServerItems *[]ac.RealServerItem) []interface{} {
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
