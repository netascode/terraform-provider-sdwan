---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_vpn_feature_template Data Source - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This data source can read the Cisco VPN feature template.
---

# sdwan_cisco_vpn_feature_template (Data Source)

This data source can read the Cisco VPN feature template.

## Example Usage

```terraform
data "sdwan_cisco_vpn_feature_template" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the feature template

### Read-Only

- `description` (String) The description of the feature template
- `device_types` (List of String) List of supported device types
- `dns_hosts` (Attributes List) Static DNS mapping (see [below for nested schema](#nestedatt--dns_hosts))
- `dns_ipv4_servers` (Attributes List) DNS (see [below for nested schema](#nestedatt--dns_ipv4_servers))
- `dns_ipv6_servers` (Attributes List) DNS (see [below for nested schema](#nestedatt--dns_ipv6_servers))
- `enhance_ecnp_keying` (Boolean) Optional packet fields for ECMP keying
- `enhance_ecnp_keying_variable` (String) Variable name
- `ipv4_static_gre_routes` (Attributes List) Configure routes pointing to a GRE tunnel (see [below for nested schema](#nestedatt--ipv4_static_gre_routes))
- `ipv4_static_ipsec_routes` (Attributes List) Configure routes pointing to a IPSEC tunnel (see [below for nested schema](#nestedatt--ipv4_static_ipsec_routes))
- `ipv4_static_routes` (Attributes List) Configure IPv4 Static Routes (see [below for nested schema](#nestedatt--ipv4_static_routes))
- `ipv4_static_service_routes` (Attributes List) Configure IPv4 Static Service Routes (see [below for nested schema](#nestedatt--ipv4_static_service_routes))
- `ipv6_static_routes` (Attributes List) Configure IPv6 Static Routes (see [below for nested schema](#nestedatt--ipv6_static_routes))
- `name` (String) The name of the feature template
- `nat64_pools` (Attributes List) Set NAT64 v4 pool range (see [below for nested schema](#nestedatt--nat64_pools))
- `nat_pools` (Attributes List) Configure NAT Pool entries (see [below for nested schema](#nestedatt--nat_pools))
- `omp_admin_distance_ipv4` (Number) omp-admin-distance-ipv4
- `omp_admin_distance_ipv4_variable` (String) Variable name
- `omp_admin_distance_ipv6` (Number) omp-admin-distance-ipv6
- `omp_admin_distance_ipv6_variable` (String) Variable name
- `omp_advertise_ipv4_routes` (Attributes List) Advertise routes to OMP (see [below for nested schema](#nestedatt--omp_advertise_ipv4_routes))
- `omp_advertise_ipv6_routes` (Attributes List) Advertise routes to OMP (see [below for nested schema](#nestedatt--omp_advertise_ipv6_routes))
- `organization_name` (String) Org Name selected
- `port_forward_rules` (Attributes List) Configure Port Forward entries (see [below for nested schema](#nestedatt--port_forward_rules))
- `route_global_exports` (Attributes List) Enable route leaking to Global VPN from this Service VPN (see [below for nested schema](#nestedatt--route_global_exports))
- `route_global_imports` (Attributes List) Enable route leaking from Global VPN to this Service VPN (see [below for nested schema](#nestedatt--route_global_imports))
- `route_vpn_imports` (Attributes List) Enable route leak from Service VPN to current VPN (see [below for nested schema](#nestedatt--route_vpn_imports))
- `services` (Attributes List) Configure services (see [below for nested schema](#nestedatt--services))
- `static_nat_rules` (Attributes List) Configure static NAT entries (see [below for nested schema](#nestedatt--static_nat_rules))
- `static_nat_subnet_rules` (Attributes List) Configure static NAT Subnet entries (see [below for nested schema](#nestedatt--static_nat_subnet_rules))
- `template_type` (String) The template type
- `tenant_vpn_id` (Number) Tenant VPN
- `vpn_id` (Number) List of VPN instances
- `vpn_name` (String) Name
- `vpn_name_variable` (String) Variable name

<a id="nestedatt--dns_hosts"></a>
### Nested Schema for `dns_hosts`

Read-Only:

- `hostname` (String) Hostname
- `hostname_variable` (String) Variable name
- `ip` (List of String) List of IP
- `ip_variable` (String) Variable name


<a id="nestedatt--dns_ipv4_servers"></a>
### Nested Schema for `dns_ipv4_servers`

Read-Only:

- `address` (String) DNS Address
- `role` (String) Role
- `role_variable` (String) Variable name


<a id="nestedatt--dns_ipv6_servers"></a>
### Nested Schema for `dns_ipv6_servers`

Read-Only:

- `address` (String) DNS Address
- `role` (String) Role
- `role_variable` (String) Variable name


<a id="nestedatt--ipv4_static_gre_routes"></a>
### Nested Schema for `ipv4_static_gre_routes`

Read-Only:

- `interface` (List of String) List of GRE Interfaces
- `interface_variable` (String) Variable name
- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name
- `vpn_id` (Number) Destination VPN to resolve the prefix


<a id="nestedatt--ipv4_static_ipsec_routes"></a>
### Nested Schema for `ipv4_static_ipsec_routes`

Read-Only:

- `interface` (List of String) List of IPSEC Interfaces (Separated by commas)
- `interface_variable` (String) Variable name
- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name
- `vpn_id` (Number) Destination VPN to resolve the prefix


<a id="nestedatt--ipv4_static_routes"></a>
### Nested Schema for `ipv4_static_routes`

Read-Only:

- `dhcp` (Boolean) Default Gateway obtained from DHCP
- `dhcp_variable` (String) Variable name
- `distance` (Number) Administrative distance
- `distance_variable` (String) Variable name
- `next_hops` (Attributes List) IP gateway address (see [below for nested schema](#nestedatt--ipv4_static_routes--next_hops))
- `null0` (Boolean) null0
- `null0_variable` (String) Variable name
- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name
- `track_next_hops` (Attributes List) IP gateway address (see [below for nested schema](#nestedatt--ipv4_static_routes--track_next_hops))
- `vpn_id` (Number) Destination VPN(!=0 or !=512) to resolve the prefix
- `vpn_id_variable` (String) Variable name

<a id="nestedatt--ipv4_static_routes--next_hops"></a>
### Nested Schema for `ipv4_static_routes.next_hops`

Read-Only:

- `address` (String) IP Address
- `address_variable` (String) Variable name
- `distance` (Number) Administrative distance
- `distance_variable` (String) Variable name


<a id="nestedatt--ipv4_static_routes--track_next_hops"></a>
### Nested Schema for `ipv4_static_routes.track_next_hops`

Read-Only:

- `address` (String) IP Address
- `address_variable` (String) Variable name
- `distance` (Number) Administrative distance
- `distance_variable` (String) Variable name
- `tracker` (String) Static route tracker
- `tracker_variable` (String) Variable name



<a id="nestedatt--ipv4_static_service_routes"></a>
### Nested Schema for `ipv4_static_service_routes`

Read-Only:

- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name
- `service` (String) Service
- `vpn_id` (Number) Destination VPN to resolve the prefix


<a id="nestedatt--ipv6_static_routes"></a>
### Nested Schema for `ipv6_static_routes`

Read-Only:

- `nat` (String) NAT
- `nat_variable` (String) Variable name
- `next_hops` (Attributes List) IP gateway address (see [below for nested schema](#nestedatt--ipv6_static_routes--next_hops))
- `null0` (Boolean) null0
- `null0_variable` (String) Variable name
- `prefix` (String) Prefix
- `prefix_variable` (String) Variable name
- `vpn_id` (Number) Destination VPN(!=0 or !=512) to resolve the prefix
- `vpn_id_variable` (String) Variable name

<a id="nestedatt--ipv6_static_routes--next_hops"></a>
### Nested Schema for `ipv6_static_routes.next_hops`

Read-Only:

- `address` (String) IP Address
- `address_variable` (String) Variable name
- `distance` (Number) Administrative distance
- `distance_variable` (String) Variable name



<a id="nestedatt--nat64_pools"></a>
### Nested Schema for `nat64_pools`

Read-Only:

- `end_address` (String) Ending IP address of NAT pool range
- `end_address_variable` (String) Variable name
- `leak_from_global` (Boolean) Enable Route Leaking from Global VPN to this Service VPN
- `leak_from_global_protocol` (String) Select protocol for route leaking
- `leak_to_global` (Boolean) Enable Route Leaking from this Service VPN to Global VPN
- `name` (String) NAT64 Pool name
- `overload` (Boolean) NAT 64 Overload Option
- `overload_variable` (String) Variable name
- `start_address` (String) Starting IP address of NAT pool range
- `start_address_variable` (String) Variable name


<a id="nestedatt--nat_pools"></a>
### Nested Schema for `nat_pools`

Read-Only:

- `direction` (String) Direction of NAT translation
- `direction_variable` (String) Variable name
- `name` (Number) NAT Pool Name, natpool1..31
- `name_variable` (String) Variable name
- `overload` (Boolean) Enable port translation(PAT)
- `overload_variable` (String) Variable name
- `prefix_length` (Number) Ending IP address of NAT Pool Prefix Length
- `prefix_length_variable` (String) Variable name
- `range_end` (String) Ending IP address of NAT pool range
- `range_end_variable` (String) Variable name
- `range_start` (String) Starting IP address of NAT pool range
- `range_start_variable` (String) Variable name
- `tracker_id` (Number) Add Object/Object Group Tracker
- `tracker_id_variable` (String) Variable name


<a id="nestedatt--omp_advertise_ipv4_routes"></a>
### Nested Schema for `omp_advertise_ipv4_routes`

Read-Only:

- `prefixes` (Attributes List) (see [below for nested schema](#nestedatt--omp_advertise_ipv4_routes--prefixes))
- `protocol` (String) Advertised routes protocol
- `protocol_sub_type` (List of String)
- `protocol_sub_type_variable` (String) Variable name
- `protocol_variable` (String) Variable name
- `route_policy` (String) Set Route Policy to OMP
- `route_policy_variable` (String) Variable name

<a id="nestedatt--omp_advertise_ipv4_routes--prefixes"></a>
### Nested Schema for `omp_advertise_ipv4_routes.prefixes`

Read-Only:

- `aggregate_only` (Boolean) Aggregate Only
- `aggregate_only_variable` (String) Variable name
- `prefix_entry` (String) Prefix
- `prefix_entry_variable` (String) Variable name



<a id="nestedatt--omp_advertise_ipv6_routes"></a>
### Nested Schema for `omp_advertise_ipv6_routes`

Read-Only:

- `prefixes` (Attributes List) (see [below for nested schema](#nestedatt--omp_advertise_ipv6_routes--prefixes))
- `protocol` (String) Advertised routes protocol
- `protocol_sub_type` (List of String)
- `protocol_sub_type_variable` (String) Variable name
- `protocol_variable` (String) Variable name
- `route_policy` (String)
- `route_policy_variable` (String) Variable name

<a id="nestedatt--omp_advertise_ipv6_routes--prefixes"></a>
### Nested Schema for `omp_advertise_ipv6_routes.prefixes`

Read-Only:

- `aggregate_only` (Boolean) Aggregate Only
- `aggregate_only_variable` (String) Variable name
- `prefix_entry` (String) Prefix
- `prefix_entry_variable` (String) Variable name



<a id="nestedatt--port_forward_rules"></a>
### Nested Schema for `port_forward_rules`

Read-Only:

- `pool_name` (String) NAT Pool Name, natpool1..31
- `pool_name_variable` (String) Variable name
- `protocol` (String) Protocol
- `protocol_variable` (String) Variable name
- `source_ip` (String) Source IP address to be translated
- `source_ip_variable` (String) Variable name
- `source_port` (Number) Source Port
- `source_port_variable` (String) Variable name
- `translate_ip` (String) Statically translated source IP address
- `translate_ip_variable` (String) Variable name
- `translate_port` (Number) Translate Port
- `translate_port_variable` (String) Variable name


<a id="nestedatt--route_global_exports"></a>
### Nested Schema for `route_global_exports`

Read-Only:

- `protocol` (String) Select a Route Protocol to enable route leaking from this Service VPN to Global VPN
- `protocol_sub_type` (List of String)
- `protocol_sub_type_variable` (String) Variable name
- `protocol_variable` (String) Variable name
- `redistributes` (Attributes List) Enable redistribution of replicated route protocol (see [below for nested schema](#nestedatt--route_global_exports--redistributes))
- `route_policy` (String) Select a Route Policy to enable route leaking from this Service VPN to Global VPN

<a id="nestedatt--route_global_exports--redistributes"></a>
### Nested Schema for `route_global_exports.redistributes`

Read-Only:

- `protocol` (String) Select a Route Protocol to enable redistribution
- `protocol_variable` (String) Variable name
- `route_policy` (String) Select a Route Policy to enable redistribution



<a id="nestedatt--route_global_imports"></a>
### Nested Schema for `route_global_imports`

Read-Only:

- `protocol` (String) Select a Route Protocol to enable route leaking from Global VPN to this Service VPN
- `protocol_sub_type` (List of String)
- `protocol_sub_type_variable` (String) Variable name
- `protocol_variable` (String) Variable name
- `redistributes` (Attributes List) Enable redistribution of replicated route protocol (see [below for nested schema](#nestedatt--route_global_imports--redistributes))
- `route_policy` (String) Select a Route Policy to enable route leaking from Global VPN to this Service VPN

<a id="nestedatt--route_global_imports--redistributes"></a>
### Nested Schema for `route_global_imports.redistributes`

Read-Only:

- `protocol` (String) Select a Route Protocol to enable redistribution
- `protocol_variable` (String) Variable name
- `route_policy` (String) Select a Route Policy to enable redistribution



<a id="nestedatt--route_vpn_imports"></a>
### Nested Schema for `route_vpn_imports`

Read-Only:

- `protocol` (String) Select a Route Protocol to enable route leaking to current VPN
- `protocol_sub_type` (List of String)
- `protocol_sub_type_variable` (String) Variable name
- `protocol_variable` (String) Variable name
- `redistributes` (Attributes List) Enable redistribution of replicated route protocol (see [below for nested schema](#nestedatt--route_vpn_imports--redistributes))
- `route_policy` (String) Select a Route Policy to enable route leaking to current VPN
- `route_policy_variable` (String) Variable name
- `source_vpn_id` (Number) Select a Source VPN where route leaks from
- `source_vpn_id_variable` (String) Variable name

<a id="nestedatt--route_vpn_imports--redistributes"></a>
### Nested Schema for `route_vpn_imports.redistributes`

Read-Only:

- `protocol` (String) Select a Route Protocol to enable redistribution
- `protocol_variable` (String) Variable name
- `route_policy` (String) Select a Route Policy to enable redistribution
- `route_policy_variable` (String) Variable name



<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-Only:

- `address` (List of String) List of IPv4 address
- `address_variable` (String) Variable name
- `interface` (String) Tracking Service
- `interface_variable` (String) Variable name
- `service_types` (String) Service Type
- `track_enable` (Boolean) Tracking Service
- `track_enable_variable` (String) Variable name


<a id="nestedatt--static_nat_rules"></a>
### Nested Schema for `static_nat_rules`

Read-Only:

- `pool_name` (String) NAT Pool Name, natpool1..31
- `pool_name_variable` (String) Variable name
- `source_ip` (String) Source IP address to be translated
- `source_ip_variable` (String) Variable name
- `static_nat_direction` (String) Direction of static NAT translation
- `static_nat_direction_variable` (String) Variable name
- `tracker_id` (Number) Add Object/Object Group Tracker
- `tracker_id_variable` (String) Variable name
- `translate_ip` (String) Statically translated source IP address
- `translate_ip_variable` (String) Variable name


<a id="nestedatt--static_nat_subnet_rules"></a>
### Nested Schema for `static_nat_subnet_rules`

Read-Only:

- `prefix_length` (Number) Network Prefix Length
- `prefix_length_variable` (String) Variable name
- `source_ip_subnet` (String) Source IP Subnet to be translated
- `source_ip_subnet_variable` (String) Variable name
- `static_nat_direction` (String) Direction of static NAT translation
- `static_nat_direction_variable` (String) Variable name
- `tracker_id` (Number) Add Object/Object Group Tracker
- `tracker_id_variable` (String) Variable name
- `translate_ip_subnet` (String) Statically translated source IP Subnet
- `translate_ip_subnet_variable` (String) Variable name

