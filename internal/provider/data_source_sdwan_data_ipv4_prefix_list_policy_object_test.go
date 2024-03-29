// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanDataIPv4PrefixListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanDataIPv4PrefixListPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_data_ipv4_prefix_list_policy_object.test", "entries.0.prefix", "10.0.0.0/12"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanDataIPv4PrefixListPolicyObjectConfig = `

resource "sdwan_data_ipv4_prefix_list_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    prefix = "10.0.0.0/12"
  }]
}

data "sdwan_data_ipv4_prefix_list_policy_object" "test" {
  id = sdwan_data_ipv4_prefix_list_policy_object.test.id
}
`
