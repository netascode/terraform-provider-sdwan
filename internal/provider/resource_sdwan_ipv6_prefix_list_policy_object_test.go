// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanIPv6PrefixListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanIPv6PrefixListPolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_ipv6_prefix_list_policy_object.test", "entries.0.prefix", "2001:1:1:2::/64"),
					resource.TestCheckResourceAttr("sdwan_ipv6_prefix_list_policy_object.test", "entries.0.le", "80"),
					resource.TestCheckResourceAttr("sdwan_ipv6_prefix_list_policy_object.test", "entries.0.ge", "128"),
				),
			},
		},
	})
}

func testAccSdwanIPv6PrefixListPolicyObjectConfig_all() string {
	return `
	resource "sdwan_ipv6_prefix_list_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			prefix = "2001:1:1:2::/64"
			le = 80
			ge = 128
		}]
	}
	`
}
