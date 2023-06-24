package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanLocalizedPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanLocalizedPolicyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_localized_policy.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("sdwan_localized_policy.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("sdwan_localized_policy.test", "flow_visibility_ipv4", "true"),
				),
			},
		},
	})
}

func testAccSdwanLocalizedPolicyConfig_all() string {
	return `
	resource "sdwan_localized_policy" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		flow_visibility_ipv4 = true
	}
	`
}
