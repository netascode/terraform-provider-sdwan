package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanFeatureDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanFeatureDeviceTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_feature_device_template.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("sdwan_feature_device_template.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("sdwan_feature_device_template.test", "device_type", "vedge-ISR-4331"),
					resource.TestCheckResourceAttr("sdwan_feature_device_template.test", "general_templates.0.type", "cisco_system"),
				),
			},
		},
	})
}

func testAccSdwanFeatureDeviceTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_system_feature_template" "system" {
		name = "TF_SYSTEM_1"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		hostname = "Router1"
		system_ip = "5.5.5.5"
		site_id = 1
		console_baud_rate = "115200"
		multi_tenant = true
	}

	resource "sdwan_feature_device_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_type = "vedge-ISR-4331"
		general_templates = [{
			id = sdwan_cisco_system_feature_template.system.id
			type = sdwan_cisco_system_feature_template.system.template_type
		}]
	}
	`
}
