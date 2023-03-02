// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanExpandedCommunityListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanExpandedCommunityListPolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_expanded_community_list_policy_object.test", "entries.0.community", "100:1000"),
				),
			},
		},
	})
}

func testAccSdwanExpandedCommunityListPolicyObjectConfig_all() string {
	return `
	resource "sdwan_expanded_community_list_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			community = "100:1000"
		}]
	}
	`
}