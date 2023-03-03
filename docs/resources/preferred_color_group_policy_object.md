---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_preferred_color_group_policy_object Resource - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This resource can manage a Preferred Color Group policy object.
---

# sdwan_preferred_color_group_policy_object (Resource)

This resource can manage a Preferred Color Group policy object.

## Example Usage

```terraform
resource "sdwan_preferred_color_group_policy_object" "example" {
  name = "Example"
  entries = [
    {
      primary_color_preference   = "blue bronze"
      primary_path_preference    = "direct-path"
      secondary_color_preference = "3g"
      secondary_path_preference  = "multi-hop-path"
      tertiary_color_preference  = "custom1"
      tertiary_path_preference   = "all-paths"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the policy object

### Optional

- `entries` (Attributes List) List of entries, only 1 entry supported (see [below for nested schema](#nestedatt--entries))

### Read-Only

- `id` (String) The id of the policy object

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Optional:

- `primary_color_preference` (String) Color or space separated list of colors
- `primary_path_preference` (String) Path preference
  - Choices: `direct-path`, `multi-hop-path`, `all-paths`
- `secondary_color_preference` (String) Color or space separated list of colors
- `secondary_path_preference` (String) Path preference
  - Choices: `direct-path`, `multi-hop-path`, `all-paths`
- `tertiary_color_preference` (String) Color or space separated list of colors
- `tertiary_path_preference` (String) Path preference
  - Choices: `direct-path`, `multi-hop-path`, `all-paths`

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_preferred_color_group_policy_object.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```