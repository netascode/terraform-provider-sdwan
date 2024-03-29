// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanSiteListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSiteListPolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_site_list_policy_object.test", "entries.0.site_id", "100-200"),
				),
			},
		},
	})
}

func testAccSdwanSiteListPolicyObjectConfig_all() string {
	return `
	resource "sdwan_site_list_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			site_id = "100-200"
		}]
	}
	`
}
