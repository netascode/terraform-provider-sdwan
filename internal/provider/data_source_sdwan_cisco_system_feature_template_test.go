// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanCiscoSystemFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoSystemFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "timezone", "UTC"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "hostname", "Router1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "system_description", "My Description"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "location", "Building 1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "latitude", "40"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "longitude", "50"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing_range", "1000"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing_sms", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "geo_fencing_sms_phone_numbers.0.number", "+1234567"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "system_ip", "5.5.5.5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "overlay_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "site_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "port_offset", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "port_hopping", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "control_session_pps", "300"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "track_transport", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "track_interface_tag", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "console_baud_rate", "115200"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "max_omp_sessions", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "multi_tenant", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "track_default_gateway", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "admin_tech_on_failure", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "idle_timeout", "100"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.name", "tracker1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.endpoint_ip", "5.6.7.8"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.transport_endpoint_ip", "5.6.7.8"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.transport_endpoint_protocol", "tcp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.transport_endpoint_port", "500"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.endpoint_dns_name", "abc.com"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.endpoint_api_url", "https://1.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.boolean", "or"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.threshold", "300"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.interval", "60"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.multiplier", "3"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "trackers.0.type", "interface"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.object_number", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.interface", "e1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.sig", "sig1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.ip", "6.6.6.6"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.mask", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.group_tracks_ids.0.track_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "object_trackers.0.boolean", "and"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "on_demand_tunnel", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "on_demand_tunnel_idle_timeout", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "affinity_group_number", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "transport_gateway", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "enable_mrf_migration", "enabled"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_system_feature_template.test", "migration_bgp_community", "100"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoSystemFeatureTemplateConfig = `

resource "sdwan_cisco_system_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  timezone = "UTC"
  hostname = "Router1"
  system_description = "My Description"
  location = "Building 1"
  latitude = 40
  longitude = 50
  geo_fencing = true
  geo_fencing_range = 1000
  geo_fencing_sms = true
  geo_fencing_sms_phone_numbers = [{
    number = "+1234567"
  }]
  device_groups = ["group1"]
  controller_group_list = ["1"]
  system_ip = "5.5.5.5"
  overlay_id = 1
  site_id = 1
  port_offset = 1
  port_hopping = true
  control_session_pps = 300
  track_transport = true
  track_interface_tag = 1
  console_baud_rate = "115200"
  max_omp_sessions = 5
  multi_tenant = true
  track_default_gateway = true
  admin_tech_on_failure = true
  idle_timeout = 100
  trackers = [{
    name = "tracker1"
    endpoint_ip = "5.6.7.8"
    transport_endpoint_ip = "5.6.7.8"
    transport_endpoint_protocol = "tcp"
    transport_endpoint_port = 500
    endpoint_dns_name = "abc.com"
    endpoint_api_url = "https://1.1.1.1"
    elements = ["abc", "def"]
    boolean = "or"
    threshold = 300
    interval = 60
    multiplier = 3
    type = "interface"
  }]
  object_trackers = [{
    object_number = 1
    interface = "e1"
    sig = "sig1"
    ip = "6.6.6.6"
    mask = "0.0.0.0"
    vpn_id = 1
	group_tracks_ids = [{
		track_id = 1
	}]
    boolean = "and"
  }]
  on_demand_tunnel = true
  on_demand_tunnel_idle_timeout = 10
  affinity_group_number = 5
  affinity_group_preference = [1]
  transport_gateway = true
  enable_mrf_migration = "enabled"
  migration_bgp_community = 100
}

data "sdwan_cisco_system_feature_template" "test" {
  id = sdwan_cisco_system_feature_template.test.id
}
`
