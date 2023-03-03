// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanAppProbeClassPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanAppProbeClassPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "entries.0.forwarding_class", "FC1"),
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "entries.0.mappings.0.color", "blue"),
					resource.TestCheckResourceAttr("data.sdwan_app_probe_class_policy_object.test", "entries.0.mappings.0.dscp", "8"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanAppProbeClassPolicyObjectConfig = `

resource "sdwan_app_probe_class_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    forwarding_class = "FC1"
	mappings = [{
		color = "blue"
		dscp = 8
	}]
  }]
}

data "sdwan_app_probe_class_policy_object" "test" {
  id = sdwan_app_probe_class_policy_object.test.id
}
`
