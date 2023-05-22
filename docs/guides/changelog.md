---
subcategory: "Guides"
page_title: "Changelog"
description: |-
    Changelog
---

# Changelog

## 0.1.9 (unreleased)

- Fix payload issue with `sdwan_cisco_thousandeyes_feature_template` resource
- Make `sequences` attribute of `sdwan_route_policy_definition` resource optional
- Do not bump the version number of feature templates if only the description changes

## 0.1.8

- Add `sdwan_cisco_thousandeyes_feature_template` resource and data source
- Add `sdwan_cisco_dhcp_server_feature_template` resource and data source

## 0.1.7

- Add `sdwan_qos_map_policy_definition` resource and data source
- Add `version` attribute to policy objects
- Add `sdwan_rewrite_rule_policy_definition` resource and data source
- Add `sdwan_acl_policy_definition` resource and data source
- Add `sdwan_device_acl_policy_definition` resource and data source
- Add `sdwan_route_policy_definition` resource and data source
- Add `sdwan_localized_policy` resource and data source
- Add `policy_version` attribute to `sdwan_feature_device_template` resource

## 0.1.6

- Return error if attach/detach action fails
- Add `version` attribute to feature templates
- Add `version` attribute to feature device template
- Add `version` attribute to feature templates references of device templates
- Add `version` attribute to device template reference of `sdwan_attach_feature_device_template` resource

## 0.1.5

- Fix empty tunnel-interface config of `sdwan_cisco_vpn_interface_feature_template`
- Introduce `optional` attribute to designate list items as optional

## 0.1.4

- Fix error when pushing feature template attributes without type `ignore`
- BREAKING CHANGE: Rename `enhance_ecnp_keying` attribte of `sdwan_cisco_vpn_feature_template` resource and data source to `enhance_ecmp_keying`
- Fix `sdwan_cisco_vpn_feature_template` schema inconsistency, where DNS address cannot be configured as a device specific variable
- BREAKING CHANGE: Rename `tunnel_interface_restrict` attribte of `sdwan_cisco_vpn_interface_feature_template` resource and data source to `tunnel_interface_color_restrict`
- Remove specific ignored feature template attributes from payload
- Add `sdwan_cisco_bgp_feature_template` resource and data source

## 0.1.3

- Fix string length validation of encrypted passhprases in feature templates
- Fix parsing of empty lists in feature template payloads

## 0.1.2

- Add `sdwan_application_list_policy_object` resource and data source
- Add `sdwan_color_list_policy_object` resource and data source
- Add `sdwan_site_list_policy_object` resource and data source
- Add `sdwan_app_probe_class_policy_object` resource and data source
- Add `sdwan_sla_class_policy_object` resource and data source
- Add `sdwan_tloc_list_policy_object` resource and data source
- Add `sdwan_region_list_policy_object` resource and data source
- Add `sdwan_preferred_color_group_policy_object` resource and data source

## 0.1.1

- Fix device template configuration of sub-templates
- Fix mandatory feature template attributes when using device-specific variables
- Add `sdwan_attach_feature_device_template` resource
- Add `sdwan_vpn_list_policy_object` resource and data source
- Add `sdwan_ipv4_prefix_list_policy_object` resource and data source
- Add `sdwan_ipv6_prefix_list_policy_object` resource and data source
- Add `sdwan_policer_policy_object` resource and data source
- Add `sdwan_mirror_policy_object` resource and data source
- Add `sdwan_class_map_policy_object` resource and data source
- Add `sdwan_extended_community_list_policy_object` resource and data source
- Add `sdwan_data_ipv4_prefix_list_policy_object` resource and data source
- Add `sdwan_data_ipv6_prefix_list_policy_object` resource and data source
- Add `sdwan_as_path_list_policy_object` resource and data source
- Add `sdwan_standard_community_list_policy_object` resource and data source
- Add `sdwan_expanded_community_list_policy_object` resource and data source

## 0.1.0

- Initial release

