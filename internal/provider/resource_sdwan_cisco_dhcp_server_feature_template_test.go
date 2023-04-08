// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanCiscoDHCPServerFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoDHCPServerFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "address_pool", "10.1.1.0/24"),
				),
			},
			{
				Config: testAccSdwanCiscoDHCPServerFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "address_pool", "10.1.1.0/24"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "lease_time", "600"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "interface_mtu", "1500"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "domain_name", "cisco.com"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "default_gateway", "10.1.1.254"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "static_leases.0.mac_address", "11:11:11:11:11:11"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "static_leases.0.ip_address", "10.1.1.10"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "static_leases.0.hostname", "HOST1"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "options.0.option_code", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_dhcp_server_feature_template.test", "options.0.ascii", "abc"),
				),
			},
		},
	})
}

func testAccSdwanCiscoDHCPServerFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_dhcp_server_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		address_pool = "10.1.1.0/24"
	}
	`
}

func testAccSdwanCiscoDHCPServerFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_dhcp_server_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		address_pool = "10.1.1.0/24"
		exclude_addresses = ["10.1.1.1-10.1.1.5", "10.1.1.254"]
		lease_time = 600
		interface_mtu = 1500
		domain_name = "cisco.com"
		default_gateway = "10.1.1.254"
		dns_servers = ["1.2.3.4"]
		tftp_servers = ["1.2.3.4"]
		static_leases = [{
			mac_address = "11:11:11:11:11:11"
			ip_address = "10.1.1.10"
			hostname = "HOST1"
		}]
		options = [{
			option_code = 10
			ascii = "abc"
		}]
	}
	`
}