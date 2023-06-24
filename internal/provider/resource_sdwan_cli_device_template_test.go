package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanCLIDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCLIDeviceTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "device_type", "vedge-ISR-4331"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "cli_type", "device"),
					resource.TestCheckResourceAttr("sdwan_cli_device_template.test", "cli_configuration", " system\n host-name             {{test}}-ISR4331-1200-1"),
				),
			},
		},
	})
}

func testAccSdwanCLIDeviceTemplateConfig_all() string {
	return `
	resource "sdwan_cli_device_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_type = "vedge-ISR-4331"
		cli_type = "device"
		cli_configuration = " system\n host-name             {{test}}-ISR4331-1200-1"
	}
	`
}
