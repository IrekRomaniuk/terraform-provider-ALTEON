package alteon

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "context"
  "github.com/irekromaniuk/alteon-client-go"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Provider -
func Provider() *schema.Provider {
  return &schema.Provider{
	Schema: map[string]*schema.Schema{
		"username": &schema.Schema{
		  Type:        schema.TypeString,
		  Optional:    true,
		  DefaultFunc: schema.EnvDefaultFunc("ALTEON_USERNAME", nil),
		},
		"password": &schema.Schema{
		  Type:        schema.TypeString,
		  Optional:    true,
		  Sensitive:   true,
		  DefaultFunc: schema.EnvDefaultFunc("ALTEON_PASSWORD", nil),
		},
		"uri": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			DefaultFunc: schema.EnvDefaultFunc("ALTEON_URI", nil),
		  },
	  },  
    ResourcesMap: map[string]*schema.Resource{},
	DataSourcesMap: map[string]*schema.Resource{
		"real_server":     dataSourceRealServer(),
	},
	ConfigureContextFunc: providerConfigure,
  }
}  

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	uri := d.Get("uri").(string)
  
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Warning Message Summary",
		Detail:   "This is the detailed warning message from providerConfigure",
	})
  
	if (username != "") && (password != "") {
	  c, err := alteon.NewClient(&uri, &username, &password)
	  if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   "Unable to auth user for authenticated HashiCups client",
			})
			return nil, diags
	  }
  
	  return c, diags
	}
  
	c, err := alteon.NewClient(nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create HashiCups client",
			Detail:   "Unable to auth user for unauthenticated HashiCups client",
		})
		return nil, diags
	}
  
	return c, diags
  }
  
