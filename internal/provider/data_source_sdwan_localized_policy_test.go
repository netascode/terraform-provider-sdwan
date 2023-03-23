package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanLocalizedPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanLocalizedPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "name", "TF_TEST_ALL"),
					resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "description", "Terraform integration test"),
					resource.TestCheckResourceAttr("data.sdwan_localized_policy.test", "flow_visibility_ipv4", "true"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanLocalizedPolicyConfig = `

resource "sdwan_localized_policy" "test" {
	name = "TF_TEST_ALL"
	description = "Terraform integration test"
	flow_visibility_ipv4 = true
}

data "sdwan_localized_policy" "test" {
  id = sdwan_localized_policy.test.id
}
`
