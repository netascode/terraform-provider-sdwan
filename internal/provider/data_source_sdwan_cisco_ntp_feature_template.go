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
	_ datasource.DataSource              = &CiscoNTPFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoNTPFeatureTemplateDataSource{}
)

func NewCiscoNTPFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoNTPFeatureTemplateDataSource{}
}

type CiscoNTPFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoNTPFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_ntp_feature_template"
}

func (d *CiscoNTPFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco NTP feature template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Required:            true,
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
			"master": schema.BoolAttribute{
				MarkdownDescription: "Configure device as NTP master",
				Computed:            true,
			},
			"master_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"master_stratum": schema.Int64Attribute{
				MarkdownDescription: "Master Stratum <1..15>",
				Computed:            true,
			},
			"master_stratum_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"master_source_interface": schema.StringAttribute{
				MarkdownDescription: "Set interface for NTP Master",
				Computed:            true,
			},
			"master_source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"trusted_keys": schema.ListAttribute{
				MarkdownDescription: "Designate authentication key as trustworthy",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"trusted_keys_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"authentication_keys": schema.ListNestedAttribute{
				MarkdownDescription: "Set MD5 authentication key",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "MD5 authentication key ID",
							Computed:            true,
						},
						"id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: "Enter cleartext or AES-encrypted MD5 authentication key",
							Computed:            true,
						},
						"value_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"servers": schema.ListNestedAttribute{
				MarkdownDescription: "Configure NTP servers",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip": schema.StringAttribute{
							MarkdownDescription: "Set hostname or IP address of server",
							Computed:            true,
						},
						"hostname_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"authentication_key_id": schema.Int64Attribute{
							MarkdownDescription: "Set authentication key for the server",
							Computed:            true,
						},
						"authentication_key_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "Set VPN in which NTP server is located",
							Computed:            true,
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "Set NTP version",
							Computed:            true,
						},
						"version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: "Set interface to use to reach NTP server",
							Computed:            true,
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"prefer": schema.BoolAttribute{
							MarkdownDescription: "Prefer this NTP server",
							Computed:            true,
						},
						"prefer_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *CiscoNTPFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*sdwan.Client)
}

func (d *CiscoNTPFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoNTP

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
