package routeros

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testSystemWatchdog = "routeros_system_watchdog.test"

func TestAccSystemWatchdogTest_basic(t *testing.T) {
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
						Config: fmt.Sprintf(`%v

resource "routeros_system_watchdog" "test" {
}
`, providerConfig),
						Check: resource.ComposeTestCheckFunc(
							testResourcePrimaryInstanceId(testSystemWatchdog),
						),
					},
				},
			})
		})
	}
}
