---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_class_map_policy_object Resource - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This resource can manage a Class Map policy object.
---

# sdwan_class_map_policy_object (Resource)

This resource can manage a Class Map policy object.

## Example Usage

```terraform
resource "sdwan_class_map_policy_object" "example" {
  name = "Example"
  entries = [
    {
      queue = 2
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

- `queue` (Number) Queue
  - Range: `0`-`7`

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_class_map_policy_object.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
