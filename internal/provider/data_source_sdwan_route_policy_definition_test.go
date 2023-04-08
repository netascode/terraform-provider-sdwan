// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanRoutePolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanRoutePolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "default_action", "reject"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.id", "10"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.ip_type", "ipv4"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.name", "Sequence 10"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.base_action", "accept"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.match_entries.0.type", "metric"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.match_entries.0.metric", "100"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.action_entries.0.type", "aggregator"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.action_entries.0.aggregator", "10"),
					resource.TestCheckResourceAttr("data.sdwan_route_policy_definition.test", "sequences.0.action_entries.0.aggregator_ip_address", "10.1.2.3"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanRoutePolicyDefinitionConfig = `

resource "sdwan_route_policy_definition" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  default_action = "reject"
  sequences = [{
    id = 10
    ip_type = "ipv4"
    name = "Sequence 10"
    base_action = "accept"
	match_entries = [{
		type = "metric"
		metric = 100
	}]
	action_entries = [{
		type = "aggregator"
		aggregator = 10
		aggregator_ip_address = "10.1.2.3"
	}]
  }]
}

data "sdwan_route_policy_definition" "test" {
  id = sdwan_route_policy_definition.test.id
}
`