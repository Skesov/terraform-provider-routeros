package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testIpServiceWebserver = "routeros_ip_service_webserver.test"

func TestAccIpServiceWebserverTest_basic(t *testing.T) {
	if !testCheckMinVersion(t, "7.25") {
		t.Logf("Test skipped, the minimum required version is 7.25")
		return
	}

	for _, name := range testNames {
		t.Run(name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					testAccPreCheck(t)
					testSetTransportEnv(t, name)
				},
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccIpServiceWebserverConfig(false),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testIpServiceWebserver),
							resource.TestCheckResourceAttr(testIpServiceWebserver, "graphs_plain", "false"),
							resource.TestCheckResourceAttr(testIpServiceWebserver, "scep_plain", "false"),
						),
					},
					{
						Config: testAccIpServiceWebserverConfig(true),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testIpServiceWebserver, "graphs_plain", "true"),
							resource.TestCheckResourceAttr(testIpServiceWebserver, "scep_plain", "true"),
						),
					},
				},
			})
		})
	}
}

func testAccIpServiceWebserverConfig(b bool) string {
	return fmt.Sprintf(`%v

resource "routeros_ip_service_webserver" "test" {
	graphs_plain = %v
	scep_plain   = %v
}
`, providerConfig, b, b)
}
