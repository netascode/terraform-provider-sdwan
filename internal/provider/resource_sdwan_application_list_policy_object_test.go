// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanApplicationListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanApplicationListPolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_application_list_policy_object.test", "entries.0.application", "netflix"),
				),
			},
		},
	})
}

func testAccSdwanApplicationListPolicyObjectConfig_all() string {
	return `
	resource "sdwan_application_list_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			application = "netflix"
		}]
	}
	`
}
