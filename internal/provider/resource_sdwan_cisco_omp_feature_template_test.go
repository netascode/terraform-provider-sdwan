// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanCiscoOMPFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoOMPFeatureTemplateConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccSdwanCiscoOMPFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "graceful_restart", "true"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "overlay_as", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "send_path_limit", "4"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "ecmp_limit", "4"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "omp_admin_distance_ipv4", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "omp_admin_distance_ipv6", "10"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertisement_interval", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "graceful_restart_timer", "43200"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "eor_timer", "300"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "holdtime", "60"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "ignore_region_path_length", "false"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "transport_gateway", "prefer"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertise_ipv4_routes.0.protocol", "ospf"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertise_ipv4_routes.0.advertise_external_ospf", "external"),
					resource.TestCheckResourceAttr("sdwan_cisco_omp_feature_template.test", "advertise_ipv6_routes.0.protocol", "ospf"),
				),
			},
		},
	})
}

func testAccSdwanCiscoOMPFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_omp_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
	}
	`
}

func testAccSdwanCiscoOMPFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_omp_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		graceful_restart = true
		overlay_as = 1
		send_path_limit = 4
		ecmp_limit = 4
		shutdown = false
		omp_admin_distance_ipv4 = 10
		omp_admin_distance_ipv6 = 10
		advertisement_interval = 1
		graceful_restart_timer = 43200
		eor_timer = 300
		holdtime = 60
		ignore_region_path_length = false
		transport_gateway = "prefer"
		advertise_ipv4_routes = [{
			protocol = "ospf"
			advertise_external_ospf = "external"
		}]
		advertise_ipv6_routes = [{
			protocol = "ospf"
		}]
	}
	`
}
