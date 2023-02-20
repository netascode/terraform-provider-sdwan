---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_sig_credentials_feature_template Resource - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This resource can manage a Cisco SIG Credentials feature template.
    - Minimum vManage version: 15.0.0
---

# sdwan_cisco_sig_credentials_feature_template (Resource)

This resource can manage a Cisco SIG Credentials feature template.
  - Minimum vManage version: `15.0.0`

## Example Usage

```terraform
resource "sdwan_cisco_sig_credentials_feature_template" "example" {
  name                     = "Example"
  description              = "My Example"
  device_types             = ["vedge-C8000V"]
  zscaler_organization     = "org1"
  zscaler_partner_base_uri = "abc"
  zscaler_username         = "user1"
  zscaler_password         = "password123"
  zscaler_cloud_name       = 1
  zscaler_partner_username = "partner1"
  zscaler_partner_password = "password123"
  zscaler_partner_api_key  = "key123"
  umbrella_api_key         = "key123"
  umbrella_api_secret      = "secret123"
  umbrella_organization_id = "org1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the feature template
- `device_types` (List of String) List of supported device types
  - Choices: `vedge-C8000V`, `vedge-C8300-1N1S-4T2X`, `vedge-C8300-1N1S-6T`, `vedge-C8300-2N2S-6T`, `vedge-C8300-2N2S-4T2X`, `vedge-C8500-12X4QC`, `vedge-C8500-12X`, `vedge-C8500-20X6C`, `vedge-C8500L-8S4X`, `vedge-C8200-1N-4T`, `vedge-C8200L-1N-4T`
- `name` (String) The name of the feature template
- `umbrella_api_key` (String) API Key
- `umbrella_api_secret` (String) API Secret
- `umbrella_organization_id` (String) Ord ID
- `zscaler_cloud_name` (Number) Third Party Cloud Name
  - Range: `0`-`255`
- `zscaler_organization` (String) Organization Name
- `zscaler_partner_api_key` (String) Partner API Key
- `zscaler_partner_base_uri` (String) Partner Base URI to be used in REST calls
- `zscaler_partner_password` (String) Partner Password
- `zscaler_partner_username` (String) Partner User Name
- `zscaler_password` (String) Password of Zscaler partner account
- `zscaler_username` (String) Username of Zscaler partner account

### Optional

- `umbrella_api_key_variable` (String) Variable name
- `umbrella_api_secret_variable` (String) Variable name
- `umbrella_organization_id_variable` (String) Variable name
- `zscaler_cloud_name_variable` (String) Variable name
- `zscaler_organization_variable` (String) Variable name
- `zscaler_partner_api_key_variable` (String) Variable name
- `zscaler_partner_base_uri_variable` (String) Variable name
- `zscaler_partner_password_variable` (String) Variable name
- `zscaler_partner_username_variable` (String) Variable name
- `zscaler_password_variable` (String) Variable name
- `zscaler_username_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the feature template
- `template_type` (String) The template type

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_cisco_sig_credentials_feature_template.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```