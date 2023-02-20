// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceSdwanCEdgeGlobalFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCEdgeGlobalFeatureTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "nat64_udp_timeout", "300"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "nat64_tcp_timeout", "3600"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "http_authentication", "local"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "ssh_version", "2"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "http_server", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "https_server", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "source_interface", "e1"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "ip_source_routing", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "arp_proxy", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "ftp_passive", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "rsh_rcp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "bootp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "domain_lookup", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "tcp_keepalives_out", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "tcp_keepalives_in", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "tcp_small_servers", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "udp_small_servers", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "lldp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "cdp", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "snmp_ifindex_persist", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "console_logging", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "vty_logging", "true"),
					resource.TestCheckResourceAttr("data.sdwan_cedge_global_feature_template.test", "line_vty", "true"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCEdgeGlobalFeatureTemplateConfig = `

resource "sdwan_cedge_global_feature_template" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
  device_types = ["vedge-C8000V"]
  nat64_udp_timeout = 300
  nat64_tcp_timeout = 3600
  http_authentication = "local"
  ssh_version = 2
  http_server = true
  https_server = true
  source_interface = "e1"
  ip_source_routing = true
  arp_proxy = true
  ftp_passive = true
  rsh_rcp = true
  bootp = true
  domain_lookup = true
  tcp_keepalives_out = true
  tcp_keepalives_in = true
  tcp_small_servers = true
  udp_small_servers = true
  lldp = true
  cdp = true
  snmp_ifindex_persist = true
  console_logging = true
  vty_logging = true
  line_vty = true
}

data "sdwan_cedge_global_feature_template" "test" {
  id = sdwan_cedge_global_feature_template.test.id
}
`