package alteon

import (
	"testing"
	//"os"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ac "github.com/irekromaniuk/alteon-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRealServer_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:     func() { testAccPreCheck(t) },
        Providers:    testAccProviders,
        CheckDestroy: testAccRealServerDestroy,
        Steps: []resource.TestStep{
            {
                Config: testAccRealServerConfig(),
                Check: resource.ComposeTestCheckFunc(
                    testAccCheckRealServerExists("alteon_real_server.TestServer.index"),
                    resource.TestCheckResourceAttr("alteon_real_server.TestServer", "index", "TestServer"),
                ),
            },
        },
    })
}


// testAccCheckRealServerExists uses the Example SDK directly to retrieve 
func testAccCheckRealServerExists(resourceName string) resource.TestCheckFunc {
    return func(s *terraform.State) error {
        // retrieve the resource by name from state
        rs, ok := s.RootModule().Resources[resourceName]
        if !ok {
            return fmt.Errorf("Not found: %s", resourceName)
        }
        if rs.Primary.ID == "" {
            return fmt.Errorf("Resource ID is not set")
        }
        // retrieve the client from the test provider
        client := testAccProvider.Meta().(*ac.Client)
		Table  := "SlbNewCfgEnhRealServerTable"
		RealServer, err := client.GetItem(Table, resourceName)
        if err != nil {
            return fmt.Errorf("error fetching item with resource %s. %s", resourceName, err)
        }
        return nil
    }
}

func testAccRealServerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ac.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "alteon_real_server" {
			continue
		}
		Table  := "SlbNewCfgEnhRealServerTable"
		RealServer, err := client.GetItem(Table, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Client can't connect")
		}
		Items := RealServer[Table]
		helper, err := json.Marshal(Items)
		if err != nil {
			return fmt.Errorf("Client response can't be decoded")
		}
		var Item []ac.RealServerItem
		json.Unmarshal(helper, &Item)
		/*{
			"SlbNewCfgEnhRealServerTable": []
		}*/
		/*notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}*/
		if len(Item)>0 {
			return fmt.Errorf("Resource still exists")
		}
	}

	return nil
}

func testAccRealServerConfig() string {
	return fmt.Sprintf(`
		resource "alteon_real_server" "TestServer" {
		index        = "TestServer"
		items {
			ipaddr="1.1.1.2"
			name="description2"
			state=2
			}
		}
	`)
}
