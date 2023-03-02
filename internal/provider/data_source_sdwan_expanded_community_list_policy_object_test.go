// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanExpandedCommunityListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanExpandedCommunityListPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_expanded_community_list_policy_object.test", "entries.0.community", "100:1000"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanExpandedCommunityListPolicyObjectConfig = `

resource "sdwan_expanded_community_list_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    community = "100:1000"
  }]
}

data "sdwan_expanded_community_list_policy_object" "test" {
  id = sdwan_expanded_community_list_policy_object.test.id
}
`