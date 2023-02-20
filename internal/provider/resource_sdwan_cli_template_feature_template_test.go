// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSdwanCLITemplateFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCLITemplateFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cli_template_feature_template.test", "cli_config", "! Enable new BGP community format\nip bgp-community new-format\n"),
				),
			},
			{
				Config: testAccSdwanCLITemplateFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cli_template_feature_template.test", "cli_config", "! Enable new BGP community format\nip bgp-community new-format\n"),
				),
			},
		},
	})
}

func testAccSdwanCLITemplateFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cli_template_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		cli_config = "! Enable new BGP community format\nip bgp-community new-format\n"
	}
	`
}

func testAccSdwanCLITemplateFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cli_template_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		cli_config = "! Enable new BGP community format\nip bgp-community new-format\n"
	}
	`
}
