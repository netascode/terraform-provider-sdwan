package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanCLIDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCLIDeviceTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "device_type", "vedge-ISR-4331"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "cli_type", "device"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "cli_configuration", " system\n host-name             {{test}}-ISR4331-1200-1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCLIDeviceTemplateConfig = `

resource "sdwan_cli_device_template" "test" {
  name = "TF_TEST_ALL"
  description = "Terraform integration test"
  device_type = "vedge-ISR-4331"
  cli_type = "device"
  cli_configuration = " system\n host-name             {{test}}-ISR4331-1200-1"
}

data "sdwan_cli_device_template" "test" {
  id = sdwan_cli_device_template.test.id
}
`
