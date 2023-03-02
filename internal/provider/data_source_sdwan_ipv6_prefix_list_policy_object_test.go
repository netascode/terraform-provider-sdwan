// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanIPv6PrefixListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanIPv6PrefixListPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_ipv6_prefix_list_policy_object.test", "entries.0.prefix", "2001:1:1:2::/64"),
					resource.TestCheckResourceAttr("data.sdwan_ipv6_prefix_list_policy_object.test", "entries.0.le", "80"),
					resource.TestCheckResourceAttr("data.sdwan_ipv6_prefix_list_policy_object.test", "entries.0.ge", "128"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanIPv6PrefixListPolicyObjectConfig = `

resource "sdwan_ipv6_prefix_list_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    prefix = "2001:1:1:2::/64"
    le = 80
    ge = 128
  }]
}

data "sdwan_ipv6_prefix_list_policy_object" "test" {
  id = sdwan_ipv6_prefix_list_policy_object.test.id
}
`
