package alteon

import (
	"testing"
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
  testAccProvider = Provider().(*schema.Provider)
  testAccProviders = map[string]*schema.Provider{
    "alteon": testAccProvider,
  }
}

func TestAccProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

/*var providerFactories = map[string]func() (*schema.Provider, error){
	"scaffolding": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestAccProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}*/

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.
	if v := os.Getenv("ALTEON_USERNAME"); v == "" {
		t.Fatal("ALTEON_USERNAME must be set for acceptance tests")
	  }
	  if v := os.Getenv("ALTEON_PASSWORD"); v == "" {
		t.Fatal("ALTEON_PASSWORD must be set for acceptance tests")
	  }
	  if v := os.Getenv("ALTEON_URI"); v == "" {
		t.Fatal("ALTEON_URI must be set for acceptance tests")
	  }
}