// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanSLAClassPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanSLAClassPolicyObjectConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.jitter", "100"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.latency", "10"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.loss", "1"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.fallback_best_tunnel_criteria", "jitter-loss-latency"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.fallback_best_tunnel_jitter", "100"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.fallback_best_tunnel_latency", "10"),
					resource.TestCheckResourceAttr("data.sdwan_sla_class_policy_object.test", "entries.0.fallback_best_tunnel_loss", "1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanSLAClassPolicyObjectConfig = `

resource "sdwan_sla_class_policy_object" "test" {
  name = "TF_TEST_MIN"
  entries = [{
    jitter = 100
    latency = 10
    loss = 1
    fallback_best_tunnel_criteria = "jitter-loss-latency"
    fallback_best_tunnel_jitter = 100
    fallback_best_tunnel_latency = 10
    fallback_best_tunnel_loss = 1
  }]
}

data "sdwan_sla_class_policy_object" "test" {
  id = sdwan_sla_class_policy_object.test.id
}
`
