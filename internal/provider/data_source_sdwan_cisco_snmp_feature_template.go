// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CiscoSNMPFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoSNMPFeatureTemplateDataSource{}
)

func NewCiscoSNMPFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoSNMPFeatureTemplateDataSource{}
}

type CiscoSNMPFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoSNMPFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_snmp_feature_template"
}

func (d *CiscoSNMPFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco SNMP feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_types": schema.ListAttribute{
				MarkdownDescription: "List of supported device types",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"shutdown": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable SNMP",
				Computed:            true,
			},
			"shutdown_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"contact": schema.StringAttribute{
				MarkdownDescription: "Set the contact for this managed node",
				Computed:            true,
			},
			"contact_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"location": schema.StringAttribute{
				MarkdownDescription: "Set the physical location of this managed node",
				Computed:            true,
			},
			"location_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"views": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a view record",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Set the name of the SNMP view",
							Computed:            true,
						},
						"object_identifiers": schema.ListNestedAttribute{
							MarkdownDescription: "Configure SNMP object identifier",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Configure identifier of subtree of MIB objects",
										Computed:            true,
									},
									"id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"exclude": schema.BoolAttribute{
										MarkdownDescription: "Exclude the OID",
										Computed:            true,
									},
									"exclude_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Computed:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"communities": schema.ListNestedAttribute{
				MarkdownDescription: "Configure SNMP community",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP community",
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP view",
							Computed:            true,
						},
						"view_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"authorization": schema.StringAttribute{
							MarkdownDescription: "Configure access permissions",
							Computed:            true,
						},
						"authorization_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"groups": schema.ListNestedAttribute{
				MarkdownDescription: "Configure an SNMP group",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP group",
							Computed:            true,
						},
						"security_level": schema.StringAttribute{
							MarkdownDescription: "Configure security level",
							Computed:            true,
						},
						"view": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP view",
							Computed:            true,
						},
						"view_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"users": schema.ListNestedAttribute{
				MarkdownDescription: "Configure an SNMP user",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP user",
							Computed:            true,
						},
						"authentication_protocol": schema.StringAttribute{
							MarkdownDescription: "Configure authentication protocol",
							Computed:            true,
						},
						"authentication_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"authentication_password": schema.StringAttribute{
							MarkdownDescription: "Specify authentication protocol password",
							Computed:            true,
						},
						"authentication_password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"privacy_protocol": schema.StringAttribute{
							MarkdownDescription: "Configure privacy protocol",
							Computed:            true,
						},
						"privacy_protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"privacy_password": schema.StringAttribute{
							MarkdownDescription: "Specify privacy protocol password",
							Computed:            true,
						},
						"privacy_password_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"group": schema.StringAttribute{
							MarkdownDescription: "Name of the SNMP group",
							Computed:            true,
						},
						"group_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"trap_targets": schema.ListNestedAttribute{
				MarkdownDescription: "Configure SNMP server to receive SNMP traps",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Set VPN in which SNMP server is located",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: "Set IPv4/IPv6 address of SNMP server",
							Computed:            true,
						},
						"ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"udp_port": schema.Int64Attribute{
							MarkdownDescription: "Set UDP port number to connect to SNMP server",
							Computed:            true,
						},
						"udp_port_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"community_name": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP community",
							Computed:            true,
						},
						"community_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"user": schema.StringAttribute{
							MarkdownDescription: "Set name of the SNMP user",
							Computed:            true,
						},
						"user_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Source interface for outgoing SNMP traps",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CiscoSNMPFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*sdwan.Client)
}

func (d *CiscoSNMPFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoSNMP

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/feature/object/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
