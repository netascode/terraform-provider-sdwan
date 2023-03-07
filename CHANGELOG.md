## 0.1.4 (unreleased)

- Fix error when pushing feature template attributes without type `ignore`
- BREAKING CHANGE: Rename `enhance_ecnp_keying` attribte of `sdwan_cisco_vpn_feature_template` resource and data source to `enhance_ecmp_keying`
- Fix `sdwan_cisco_vpn_feature_template` schema inconsistency, where DNS address cannot be configured as a device specific variable

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
