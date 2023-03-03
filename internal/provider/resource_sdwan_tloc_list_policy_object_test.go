// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanTLOCListPolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTLOCListPolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.color", "blue"),
					resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.encapsulation", "gre"),
					resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.preference", "10"),
					resource.TestCheckResourceAttr("sdwan_tloc_list_policy_object.test", "entries.0.tloc_ip", "1.1.1.2"),
				),
			},
		},
	})
}

func testAccSdwanTLOCListPolicyObjectConfig_all() string {
	return `
	resource "sdwan_tloc_list_policy_object" "test" {
		name = "TF_TEST_ALL"
		entries = [{
			color = "blue"
			encapsulation = "gre"
			preference = 10
			tloc_ip = "1.1.1.2"
		}]
	}
	`
}