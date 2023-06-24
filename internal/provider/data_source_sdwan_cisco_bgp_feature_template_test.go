// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanCiscoBGPFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoBGPFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "as_number", "65000"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "shutdown", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "router_id", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "propagate_aspath", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "propagate_community", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_route_targets.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_route_targets.0.export.0.asn_ip", "10:100"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_route_targets.0.import.0.asn_ip", "10:100"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_route_targets.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_route_targets.0.export.0.asn_ip", "10:100"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_route_targets.0.import.0.asn_ip", "10:100"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "mpls_interfaces.0.interface_name", "GigabitEthernet0"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "distance_external", "30"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "distance_internal", "210"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "distance_local", "30"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "keepalive", "90"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "holdtime", "220"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "always_compare_med", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "deterministic_med", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "missing_med_worst", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "compare_router_id", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "multipath_relax", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.family_type", "ipv4-unicast"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_aggregate_addresses.0.prefix", "10.0.0.0/8"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_aggregate_addresses.0.as_set_path", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_aggregate_addresses.0.summary_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv6_aggregate_addresses.0.prefix", "10.0.0.0/8"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv6_aggregate_addresses.0.as_set_path", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv6_aggregate_addresses.0.summary_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv4_networks.0.prefix", "10.2.2.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.ipv6_networks.0.prefix", "10.2.2.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.maximum_paths", "8"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.default_information_originate", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.table_map_policy", "MAP1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.table_map_filter", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.redistribute_routes.0.protocol", "ospf"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "address_families.0.redistribute_routes.0.route_policy", "POLICY1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address", "10.2.2.2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.description", "My neighbor"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.shutdown", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.remote_as", "65001"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.keepalive", "30"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.holdtime", "90"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.source_interface", "GigabitEthernet1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.next_hop_self", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_community", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_ext_community", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.ebgp_multihop", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.password", "cisco123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_label", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.send_label_explicit", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.as_override", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.allow_as_in", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.family_type", "ipv4-unicast"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes", "10000"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes_threshold", "80"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes_restart", "180"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.maximum_prefixes_warning_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.route_policies.0.direction", "in"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv4_neighbors.0.address_families.0.route_policies.0.policy_name", "POLICY1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address", "2001:1::1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.description", "My neighbor"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.shutdown", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.remote_as", "65001"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.keepalive", "30"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.holdtime", "90"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.source_interface", "GigabitEthernet1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.next_hop_self", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_community", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_ext_community", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.ebgp_multihop", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.password", "cisco123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_label", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.send_label_explicit", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.as_override", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.allow_as_in", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.family_type", "ipv6-unicast"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes", "10000"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes_threshold", "80"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes_restart", "180"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.maximum_prefixes_warning_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.route_policies.0.direction", "in"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_bgp_feature_template.test", "ipv6_neighbors.0.address_families.0.route_policies.0.policy_name", "POLICY1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoBGPFeatureTemplateConfig = `

resource "sdwan_cisco_bgp_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  as_number = "65000"
  shutdown = true
  router_id = "1.2.3.4"
  propagate_aspath = true
  propagate_community = true
  ipv4_route_targets = [{
    vpn_id = 1
	export = [{
		asn_ip = "10:100"
	}]
	import = [{
		asn_ip = "10:100"
	}]
  }]
  ipv6_route_targets = [{
    vpn_id = 1
	export = [{
		asn_ip = "10:100"
	}]
	import = [{
		asn_ip = "10:100"
	}]
  }]
  mpls_interfaces = [{
    interface_name = "GigabitEthernet0"
  }]
  distance_external = 30
  distance_internal = 210
  distance_local = 30
  keepalive = 90
  holdtime = 220
  always_compare_med = true
  deterministic_med = true
  missing_med_worst = true
  compare_router_id = true
  multipath_relax = true
  address_families = [{
    family_type = "ipv4-unicast"
	ipv4_aggregate_addresses = [{
		prefix = "10.0.0.0/8"
		as_set_path = true
		summary_only = true
	}]
	ipv6_aggregate_addresses = [{
		prefix = "10.0.0.0/8"
		as_set_path = true
		summary_only = true
	}]
	ipv4_networks = [{
		prefix = "10.2.2.0/24"
	}]
	ipv6_networks = [{
		prefix = "10.2.2.0/24"
	}]
    maximum_paths = 8
    default_information_originate = true
    table_map_policy = "MAP1"
    table_map_filter = true
	redistribute_routes = [{
		protocol = "ospf"
		route_policy = "POLICY1"
	}]
  }]
  ipv4_neighbors = [{
    address = "10.2.2.2"
    description = "My neighbor"
    shutdown = true
    remote_as = "65001"
    keepalive = 30
    holdtime = 90
    source_interface = "GigabitEthernet1"
    next_hop_self = true
    send_community = false
    send_ext_community = false
    ebgp_multihop = 10
    password = "cisco123"
    send_label = true
    send_label_explicit = true
    as_override = true
    allow_as_in = 5
	address_families = [{
		family_type = "ipv4-unicast"
		maximum_prefixes = 10000
		maximum_prefixes_threshold = 80
		maximum_prefixes_restart = 180
		maximum_prefixes_warning_only = true
		route_policies = [{
			direction = "in"
			policy_name = "POLICY1"
		}]
	}]
  }]
  ipv6_neighbors = [{
    address = "2001:1::1"
    description = "My neighbor"
    shutdown = true
    remote_as = "65001"
    keepalive = 30
    holdtime = 90
    source_interface = "GigabitEthernet1"
    next_hop_self = true
    send_community = false
    send_ext_community = false
    ebgp_multihop = 10
    password = "cisco123"
    send_label = true
    send_label_explicit = true
    as_override = true
    allow_as_in = 5
	address_families = [{
		family_type = "ipv6-unicast"
		maximum_prefixes = 10000
		maximum_prefixes_threshold = 80
		maximum_prefixes_restart = 180
		maximum_prefixes_warning_only = true
		route_policies = [{
			direction = "in"
			policy_name = "POLICY1"
		}]
	}]
  }]
}

data "sdwan_cisco_bgp_feature_template" "test" {
  id = sdwan_cisco_bgp_feature_template.test.id
}
`
