package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanAttachFeatureDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanAttachFeatureDeviceTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "variables.0.var_site_id", "1001"),
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "variables.0.var_system_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "variables.0.var_hostname", "router1"),
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "variables.0.vpn_if_name_Default_vEdge_DHCP_Tunnel_Interface", "GigabitEthernet1"),
				),
			},
		},
	})
}

func testAccSdwanAttachFeatureDeviceTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_system_feature_template" "system" {
		name = "TF_SYSTEM_1"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		hostname_variable = "var_hostname"
		system_ip_variable = "var_system_ip"
		site_id_variable = "var_site_id"
		console_baud_rate = "115200"
	}

	resource "sdwan_feature_device_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_type = "vedge-C8000V"
		general_templates = [{
			id = sdwan_cisco_system_feature_template.system.id
			type = sdwan_cisco_system_feature_template.system.template_type
		}]
	}

	resource "sdwan_attach_feature_device_template" "test" {
		id = sdwan_feature_device_template.test.id
		devices = [{
			id = "C8K-CC678D1C-8EDF-3966-4F51-ABFAB64F5ABE"
			variables = {
			  var_site_id                                     = "1001"
			  var_system_ip                                   = "1.1.1.1"
			  var_hostname                                    = "router1"
			  vpn_if_name_Default_vEdge_DHCP_Tunnel_Interface = "GigabitEthernet1"
			}
		}]
	}
	`
}
