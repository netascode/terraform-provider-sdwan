// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanCiscoSIGCredentialsFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoSIGCredentialsFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_organization", "org1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_base_uri", "abc"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_username", "user1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_password", "password123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_cloud_name", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_username", "partner1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_password", "password123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_api_key", "key123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_key", "key123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_secret", "secret123"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_sig_credentials_feature_template.test", "umbrella_organization_id", "org1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoSIGCredentialsFeatureTemplateConfig = `

resource "sdwan_cisco_sig_credentials_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  zscaler_organization = "org1"
  zscaler_partner_base_uri = "abc"
  zscaler_username = "user1"
  zscaler_password = "password123"
  zscaler_cloud_name = 1
  zscaler_partner_username = "partner1"
  zscaler_partner_password = "password123"
  zscaler_partner_api_key = "key123"
  umbrella_api_key = "key123"
  umbrella_api_secret = "secret123"
  umbrella_organization_id = "org1"
}

data "sdwan_cisco_sig_credentials_feature_template" "test" {
  id = sdwan_cisco_sig_credentials_feature_template.test.id
}
`
