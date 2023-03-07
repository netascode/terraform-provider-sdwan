// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanCiscoVPNFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoVPNFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "vpn_name", "VPN1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "tenant_vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "organization_name", "org1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_admin_distance_ipv4", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_admin_distance_ipv6", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "enhance_ecmp_keying", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "dns_ipv4_servers.0.address", "9.9.9.9"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "dns_ipv4_servers.0.role", "primary"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "dns_ipv6_servers.0.address", "2001::9"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "dns_ipv6_servers.0.role", "primary"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "dns_hosts.0.hostname", "abc1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "services.0.service_types", "FW"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "services.0.interface", "e1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "services.0.track_enable", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_service_routes.0.prefix", "2.2.2.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_service_routes.0.vpn_id", "2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_service_routes.0.service", "sig"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.prefix", "3.3.3.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.null0", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.distance", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.vpn_id", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.dhcp", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.next_hops.0.address", "11.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.next_hops.0.distance", "20"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.track_next_hops.0.address", "12.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.track_next_hops.0.distance", "20"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_routes.0.track_next_hops.0.tracker", "tracker1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.prefix", "2001::/48"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.null0", "false"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.vpn_id", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.nat", "NAT64"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.next_hops.0.address", "2001::11"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv6_static_routes.0.next_hops.0.distance", "20"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_gre_routes.0.prefix", "3.3.3.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_gre_routes.0.vpn_id", "2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_ipsec_routes.0.prefix", "4.4.4.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "ipv4_static_ipsec_routes.0.vpn_id", "2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.protocol", "bgp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.route_policy", "rp1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.prefixes.0.prefix_entry", "1.1.1.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv4_routes.0.prefixes.0.aggregate_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.protocol", "bgp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.route_policy", "rp1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.prefixes.0.prefix_entry", "2001:2::/48"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "omp_advertise_ipv6_routes.0.prefixes.0.aggregate_only", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.name", "POOL1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.start_address", "100.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.end_address", "100.1.2.255"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.overload", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.leak_from_global", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.leak_from_global_protocol", "rip"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat64_pools.0.leak_to_global", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.name", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.prefix_length", "24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.range_start", "101.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.range_end", "101.1.2.255"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.overload", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.direction", "inside"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "nat_pools.0.tracker_id", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.pool_name", "POOL1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.source_ip", "10.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.translate_ip", "105.1.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.static_nat_direction", "inside"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_rules.0.tracker_id", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.source_ip_subnet", "10.2.1.0"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.translate_ip_subnet", "105.2.1.0"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.prefix_length", "24"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.static_nat_direction", "inside"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "static_nat_subnet_rules.0.tracker_id", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.pool_name", "POOL2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.source_port", "5000"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.translate_port", "6000"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.source_ip", "10.3.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.translate_ip", "120.3.1.1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "port_forward_rules.0.protocol", "tcp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.protocol", "ospf"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.route_policy", "policy1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.redistributes.0.protocol", "bgp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_imports.0.redistributes.0.route_policy", "policy1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.source_vpn_id", "5"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.protocol", "ospf"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.route_policy", "policy1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.redistributes.0.protocol", "bgp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_vpn_imports.0.redistributes.0.route_policy", "policy1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.protocol", "ospf"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.route_policy", "policy1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.redistributes.0.protocol", "bgp"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_vpn_feature_template.test", "route_global_exports.0.redistributes.0.route_policy", "policy1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoVPNFeatureTemplateConfig = `

resource "sdwan_cisco_vpn_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  vpn_id = 1
  vpn_name = "VPN1"
  tenant_vpn_id = 1
  organization_name = "org1"
  omp_admin_distance_ipv4 = 10
  omp_admin_distance_ipv6 = 10
  enhance_ecmp_keying = true
  dns_ipv4_servers = [{
    address = "9.9.9.9"
    role = "primary"
  }]
  dns_ipv6_servers = [{
    address = "2001::9"
    role = "primary"
  }]
  dns_hosts = [{
    hostname = "abc1"
    ip = ["7.7.7.7"]
  }]
  services = [{
    service_types = "FW"
    address = ["8.8.8.8"]
    interface = "e1"
    track_enable = true
  }]
  ipv4_static_service_routes = [{
    prefix = "2.2.2.0/24"
    vpn_id = 2
    service = "sig"
  }]
  ipv4_static_routes = [{
    prefix = "3.3.3.0/24"
    null0 = false
    distance = 10
    vpn_id = 5
    dhcp = false
	next_hops = [{
		address = "11.1.1.1"
		distance = 20
	}]
	track_next_hops = [{
		address = "12.1.1.1"
		distance = 20
		tracker = "tracker1"
	}]
  }]
  ipv6_static_routes = [{
    prefix = "2001::/48"
    null0 = false
    vpn_id = 5
    nat = "NAT64"
	next_hops = [{
		address = "2001::11"
		distance = 20
	}]
  }]
  ipv4_static_gre_routes = [{
    prefix = "3.3.3.0/24"
    vpn_id = 2
    interface = ["e1"]
  }]
  ipv4_static_ipsec_routes = [{
    prefix = "4.4.4.0/24"
    vpn_id = 2
    interface = ["e1"]
  }]
  omp_advertise_ipv4_routes = [{
    protocol = "bgp"
    route_policy = "rp1"
    protocol_sub_type = ["external"]
	prefixes = [{
		prefix_entry = "1.1.1.0/24"
		aggregate_only = true
	}]
  }]
  omp_advertise_ipv6_routes = [{
    protocol = "bgp"
    route_policy = "rp1"
    protocol_sub_type = ["external"]
	prefixes = [{
		prefix_entry = "2001:2::/48"
		aggregate_only = true
	}]
  }]
  nat64_pools = [{
    name = "POOL1"
    start_address = "100.1.1.1"
    end_address = "100.1.2.255"
    overload = true
    leak_from_global = true
    leak_from_global_protocol = "rip"
    leak_to_global = true
  }]
  nat_pools = [{
    name = 1
    prefix_length = 24
    range_start = "101.1.1.1"
    range_end = "101.1.2.255"
    overload = true
    direction = "inside"
    tracker_id = 10
  }]
  static_nat_rules = [{
    pool_name = "POOL1"
    source_ip = "10.1.1.1"
    translate_ip = "105.1.1.1"
    static_nat_direction = "inside"
    tracker_id = 10
  }]
  static_nat_subnet_rules = [{
    source_ip_subnet = "10.2.1.0"
    translate_ip_subnet = "105.2.1.0"
    prefix_length = 24
    static_nat_direction = "inside"
    tracker_id = 10
  }]
  port_forward_rules = [{
    pool_name = "POOL2"
    source_port = 5000
    translate_port = 6000
    source_ip = "10.3.1.1"
    translate_ip = "120.3.1.1"
    protocol = "tcp"
  }]
  route_global_imports = [{
    protocol = "ospf"
    protocol_sub_type = ["external"]
    route_policy = "policy1"
	redistributes = [{
		protocol = "bgp"
		route_policy = "policy1"
	}]
  }]
  route_vpn_imports = [{
    source_vpn_id = 5
    protocol = "ospf"
    protocol_sub_type = ["external"]
    route_policy = "policy1"
	redistributes = [{
		protocol = "bgp"
		route_policy = "policy1"
	}]
  }]
  route_global_exports = [{
    protocol = "ospf"
    protocol_sub_type = ["external"]
    route_policy = "policy1"
	redistributes = [{
		protocol = "bgp"
		route_policy = "policy1"
	}]
  }]
}

data "sdwan_cisco_vpn_feature_template" "test" {
  id = sdwan_cisco_vpn_feature_template.test.id
}
`
