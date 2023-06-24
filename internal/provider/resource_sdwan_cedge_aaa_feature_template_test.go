// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanCEdgeAAAFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCEdgeAAAFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "server_groups_priority_order", "100"),
				),
			},
			{
				Config: testAccSdwanCEdgeAAAFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "dot1x_authentication", "true"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "dot1x_accounting", "true"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "server_groups_priority_order", "100"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.name", "user1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.password", "password123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.secret", "secret123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.privilege_level", "15"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.ssh_pubkeys.0.key_string", "abc123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "users.0.ssh_pubkeys.0.key_type", "rsa"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.group_name", "GROUP1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.source_interface", "e1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.authentication_port", "1812"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.accounting_port", "1813"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.timeout", "5"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.retransmit", "3"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.secret_key", "1234567"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.encryption_type", "7"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_server_groups.0.servers.0.key_type", "pac"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_clients.0.client_ip", "2.2.2.2"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_clients.0.von_configurations.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_clients.0.von_configurations.0.server_key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_server_key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_domain_stripping", "yes"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_authentication_type", "all"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_port", "1700"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_dynamic_author_cts_authorization_list", "ALIST1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "radius_trustsec_group", "GROUP1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.group_name", "GROUP1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.vpn_id", "1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.source_interface", "e1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.port", "49"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.timeout", "5"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.secret_key", "1234567"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "tacacs_server_groups.0.servers.0.encryption_type", "7"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.name", "RULE1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.method", "exec"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.privilege_level", "15"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "accounting_rules.0.start_stop", "true"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_console", "true"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_config_commands", "true"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.name", "RULE1"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.method", "commands"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.privilege_level", "15"),
					resource.TestCheckResourceAttr("sdwan_cedge_aaa_feature_template.test", "authorization_rules.0.authenticated", "true"),
				),
			},
		},
	})
}

func testAccSdwanCEdgeAAAFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cedge_aaa_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		server_groups_priority_order = "100"
	}
	`
}

func testAccSdwanCEdgeAAAFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cedge_aaa_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		dot1x_authentication = true
		dot1x_accounting = true
		server_groups_priority_order = "100"
		users = [{
			name = "user1"
			password = "password123"
			secret = "secret123"
			privilege_level = "15"
			ssh_pubkeys = [{
				key_string = "abc123"
				key_type = "rsa"
			}]
		}]
		radius_server_groups = [{
			group_name = "GROUP1"
			vpn_id = 1
			source_interface = "e1"
			servers = [{
				address = "1.1.1.1"
				authentication_port = 1812
				accounting_port = 1813
				timeout = 5
				retransmit = 3
				key = "key123"
				secret_key = "1234567"
				encryption_type = "7"
				key_type = "pac"
			}]
		}]
		radius_clients = [{
			client_ip = "2.2.2.2"
			von_configurations = [{
				vpn_id = "1"
				server_key = "key123"
			}]
		}]
		radius_dynamic_author_server_key = "key123"
		radius_dynamic_author_domain_stripping = "yes"
		radius_dynamic_author_authentication_type = "all"
		radius_dynamic_author_port = 1700
		radius_dynamic_author_cts_authorization_list = "ALIST1"
		radius_trustsec_group = "GROUP1"
		tacacs_server_groups = [{
			group_name = "GROUP1"
			vpn_id = 1
			source_interface = "e1"
			servers = [{
				address = "1.1.1.1"
				port = 49
				timeout = 5
				key = "key123"
				secret_key = "1234567"
				encryption_type = "7"
			}]
		}]
		accounting_rules = [{
			name = "RULE1"
			method = "exec"
			privilege_level = "15"
			start_stop = true
			group = ["GROUP1"]
		}]
		authorization_console = true
		authorization_config_commands = true
		authorization_rules = [{
			name = "RULE1"
			method = "commands"
			privilege_level = "15"
			group = ["GROUP1"]
			authenticated = true
		}]
	}
	`
}
