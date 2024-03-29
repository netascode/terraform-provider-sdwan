// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanStandardCommunityListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanStandardCommunityListPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_standard_community_list_policy_object.test", "entries.0.community", "100:1000"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanStandardCommunityListPolicyObjectConfig = `

resource "sdwan_standard_community_list_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    community = "100:1000"
  }]
}

data "sdwan_standard_community_list_policy_object" "test" {
  id = sdwan_standard_community_list_policy_object.test.id
}
`
