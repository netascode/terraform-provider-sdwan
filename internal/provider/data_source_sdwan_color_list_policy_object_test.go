// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanColorListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanColorListPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_color_list_policy_object.test", "entries.0.color", "blue"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanColorListPolicyObjectConfig = `

resource "sdwan_color_list_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    color = "blue"
  }]
}

data "sdwan_color_list_policy_object" "test" {
  id = sdwan_color_list_policy_object.test.id
}
`