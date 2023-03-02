// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanClassMapPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanClassMapPolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_class_map_policy_object.test", "entries.0.queue", "2"),
				),
			},
		},
	})
}

func testAccSdwanClassMapPolicyObjectConfig_all() string {
	return `
	resource "sdwan_class_map_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			queue = 2
		}]
	}
	`
}