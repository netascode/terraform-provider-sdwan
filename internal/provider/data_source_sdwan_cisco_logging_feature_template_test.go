// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanCiscoLoggingFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCiscoLoggingFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "disk_logging", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "max_size", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "log_rotations", "10"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.name", "PROF1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.version", "TLSv1.2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.authentication_type", "Server"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "tls_profiles.0.ciphersuite_list", "aes-128-cbc-sha"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.hostname_ip", "2.2.2.2"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.source_interface", "e1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.logging_level", "information"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.enable_tls", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.custom_profile", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv4_servers.0.profile", "PROF1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.hostname_ip", "2001::1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.source_interface", "e1"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.logging_level", "information"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.enable_tls", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.custom_profile", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cisco_logging_feature_template.test", "ipv6_servers.0.profile", "PROF1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCiscoLoggingFeatureTemplateConfig = `

resource "sdwan_cisco_logging_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  disk_logging = true
  max_size = 10
  log_rotations = 10
  tls_profiles = [{
    name = "PROF1"
    version = "TLSv1.2"
    authentication_type = "Server"
    ciphersuite_list = "aes-128-cbc-sha"
  }]
  ipv4_servers = [{
    hostname_ip = "2.2.2.2"
    vpn_id = 1
    source_interface = "e1"
    logging_level = "information"
    enable_tls = true
    custom_profile = true
    profile = "PROF1"
  }]
  ipv6_servers = [{
    hostname_ip = "2001::1"
    vpn_id = 1
    source_interface = "e1"
    logging_level = "information"
    enable_tls = true
    custom_profile = true
    profile = "PROF1"
  }]
}

data "sdwan_cisco_logging_feature_template" "test" {
  id = sdwan_cisco_logging_feature_template.test.id
}
`
