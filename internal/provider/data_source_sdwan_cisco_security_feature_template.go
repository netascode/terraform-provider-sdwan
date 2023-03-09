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
	_ datasource.DataSource              = &CiscoSecurityFeatureTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CiscoSecurityFeatureTemplateDataSource{}
)

func NewCiscoSecurityFeatureTemplateDataSource() datasource.DataSource {
	return &CiscoSecurityFeatureTemplateDataSource{}
}

type CiscoSecurityFeatureTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CiscoSecurityFeatureTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_security_feature_template"
}

func (d *CiscoSecurityFeatureTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Cisco Security feature template.",

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
			"rekey_interval": schema.Int64Attribute{
				MarkdownDescription: "Set how often to change the AES key for DTLS connections",
				Computed:            true,
			},
			"rekey_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"replay_window": schema.StringAttribute{
				MarkdownDescription: "Set the sliding replay window size",
				Computed:            true,
			},
			"replay_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"extended_ar_window": schema.Int64Attribute{
				MarkdownDescription: "Extended Anti-Replay Window",
				Computed:            true,
			},
			"extended_ar_window_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"authentication_type": schema.ListAttribute{
				MarkdownDescription: "Set the authentication type for DTLS connections",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"authentication_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"integrity_type": schema.ListAttribute{
				MarkdownDescription: "Set the authentication type for DTLS connections",
				ElementType:         types.StringType,
				Computed:            true,
			},
			"integrity_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"pairwise_keying": schema.BoolAttribute{
				MarkdownDescription: "Enable or disable IPsec pairwise-keying",
				Computed:            true,
			},
			"pairwise_keying_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"keychains": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a Keychain",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "Specify the name of the Keychain",
							Computed:            true,
						},
						"key_id": schema.Int64Attribute{
							MarkdownDescription: "Specify the Key ID",
							Computed:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Computed:            true,
						},
					},
				},
			},
			"keys": schema.ListNestedAttribute{
				MarkdownDescription: "Configure a Key",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Select the Key ID",
							Computed:            true,
						},
						"chain_name": schema.StringAttribute{
							MarkdownDescription: "Select the chain name",
							Computed:            true,
						},
						"send_id": schema.Int64Attribute{
							MarkdownDescription: "Specify the Send ID",
							Computed:            true,
						},
						"send_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"receive_id": schema.Int64Attribute{
							MarkdownDescription: "Specify the Receiver ID",
							Computed:            true,
						},
						"receive_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"crypto_algorithm": schema.StringAttribute{
							MarkdownDescription: "Crypto Algorithm",
							Computed:            true,
						},
						"key_string": schema.StringAttribute{
							MarkdownDescription: "Specify the Key String",
							Computed:            true,
						},
						"key_string_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_lifetime": schema.BoolAttribute{
							MarkdownDescription: "Configure Send lifetime Local",
							Computed:            true,
						},
						"send_lifetime_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_lifetime_start_time": schema.StringAttribute{
							MarkdownDescription: "Configure Key lifetime start time",
							Computed:            true,
						},
						"send_lifetime_end_time_format": schema.StringAttribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"send_lifetime_duration": schema.Int64Attribute{
							MarkdownDescription: "Configure Send lifetime Duration",
							Computed:            true,
						},
						"send_lifetime_duration_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"send_lifetime_end_time": schema.StringAttribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"send_lifetime_infinite": schema.BoolAttribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"send_lifetime_infinite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_lifetime": schema.BoolAttribute{
							MarkdownDescription: "Configure Accept Lifetime Local",
							Computed:            true,
						},
						"accept_lifetime_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_lifetime_start_time": schema.StringAttribute{
							MarkdownDescription: "Configure Key lifetime start time",
							Computed:            true,
						},
						"accept_lifetime_end_time_format": schema.StringAttribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"accept_lifetime_duration": schema.Int64Attribute{
							MarkdownDescription: "Configure Accept lifetime Duration",
							Computed:            true,
						},
						"accept_lifetime_duration_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_lifetime_end_time": schema.StringAttribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"accept_lifetime_infinite": schema.BoolAttribute{
							MarkdownDescription: "Configure Key lifetime end time",
							Computed:            true,
						},
						"accept_lifetime_infinite_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"include_tcp_options": schema.BoolAttribute{
							MarkdownDescription: "Configure Include TCP Options",
							Computed:            true,
						},
						"include_tcp_options_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"accept_ao_mismatch": schema.BoolAttribute{
							MarkdownDescription: "Configure Accept AO Mismatch",
							Computed:            true,
						},
						"accept_ao_mismatch_variable": schema.StringAttribute{
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

func (d *CiscoSecurityFeatureTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*sdwan.Client)
}

func (d *CiscoSecurityFeatureTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CiscoSecurity

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
