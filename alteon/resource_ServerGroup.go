package alteon

import (
  "context"
  //"strconv"
  //"fmt"
  "encoding/json"
  "time"
  ac "github.com/irekromaniuk/alteon-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	Table string = "SlbNewCfgEnhGroupTable"
)

func resourceServerGroup() *schema.Resource {
  return &schema.Resource{
    CreateContext: resourceServerGroupCreate,
    ReadContext:   resourceServerGroupRead,
    UpdateContext: resourceServerGroupUpdate,
    DeleteContext: resourceServerGroupDelete,
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
					"addserver": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
					},
					"removeserver": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true, 
						},
					"healthcheckurl": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true, 
					},
					"name": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true, 
					},  
				},
			},
	 	},
	},
  }
}

func resourceServerGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	ServerGroupID := d.Get("index").(string)
	
	items := d.Get("items").([]interface{})
	rss := []ac.ServerGroupItem{}

	
	for _, item := range items {
	  i := item.(map[string]interface{})
	  rsi := ac.ServerGroupItem{
		AddServer: i["addserver"].(string),
		RemoveServer: i["removeserver"].(string),
		HealthCheckUrl: i["healthcheckurl"].(string),
		Name : i["name"].(string),
	  }
	  rss = append(rss, rsi)
	}
	
	brss, err :=json.MarshalIndent(rss[0], "", "    ")
	//prettyJSON, _ := json.MarshalIndent(rss, "", "    ")
	/*  diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "brss",
		Detail:   fmt.Sprint(string(brss)),
	  })
	return diags */
	if err != nil {
		diag.FromErr(err)
	}
	rs, err := c.CreateItem(brss, Table, ServerGroupID)
	/*diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "rs",
		Detail:   rs.Status,
	  })*/
	if err != nil {
	  return diag.FromErr(err)
	}
	if rs.Status == "ok" {
		d.SetId(ServerGroupID)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Status",
			Detail:   rs.Status, //add message or testerr
		  })
	}
	
	resourceServerGroupRead(ctx, d, m)

	return diags
}

func resourceServerGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ac.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
  
	ServerGroupID := d.Id()
  
	ServerGroup, err := c.GetItem(Table, ServerGroupID)
	if err != nil {
	  return diag.FromErr(err)
	}
	Items := ServerGroup[Table]
	helper, err := json.Marshal(Items)
	if err != nil {
		return diag.FromErr(err)
	}
	var Item []ac.ServerGroupItem
	json.Unmarshal(helper, &Item)
	ServerGroupItems := flattenServerGroupItems(&Item) //&ServerGroup.Items
	if err := d.Set("items", ServerGroupItems); err != nil {
	  return diag.FromErr(err)
	}
  
	return diags
}

func resourceServerGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
    c := m.(*ac.Client)

	ServerGroupID := d.Id()

	if d.HasChange("items") {
		items := d.Get("items").([]interface{})
		rss := []ac.ServerGroupItem{}

		for _, item := range items {
			i := item.(map[string]interface{})

			rsi := ac.ServerGroupItem{
				AddServer: i["addserver"].(string),
				RemoveServer: i["removeserver"].(string),
				HealthCheckUrl: i["healthcheckurl"].(string),
				Name : i["name"].(string),
			}
			rss = append(rss, rsi)
		}
		brss, err :=json.MarshalIndent(rss[0], "", "    ")
		if err != nil {
			return diag.FromErr(err)
		}
		_, err = c.UpdateItem(brss, Table, ServerGroupID)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceServerGroupRead(ctx, d, m)
	
}

func resourceServerGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  // Warning or errors can be collected in a slice type
  c := m.(*ac.Client)

  // Warning or errors can be collected in a slice type
  var diags diag.Diagnostics

  ServerGroupID := d.Id()

  err := c.DeleteItem(Table, ServerGroupID)
  if err != nil {
    return diag.FromErr(err)
  }

  // d.SetId("") is automatically called assuming delete returns no errors, but
  // it is added here for explicitness.
  d.SetId("")

  return diags
}

func flattenServerGroupItems(ServerGroupItems *[]ac.ServerGroupItem) []interface{} {
	if ServerGroupItems != nil {
	  rss := make([]interface{}, len(*ServerGroupItems), len(*ServerGroupItems))
  
	  for i, ServerGroupItem := range *ServerGroupItems {
		rsi := make(map[string]interface{})
		
		rsi["addserver"] = ServerGroupItem.AddServer
		rsi["removeserver"] = ServerGroupItem.RemoveServer
		rsi["name"] = ServerGroupItem.Name
		rsi["healthcheckurl"] = ServerGroupItem.HealthCheckUrl
		rss[i] = rsi
	  }
  
	  return rss
	}
  
	return make([]interface{}, 0)
  }
