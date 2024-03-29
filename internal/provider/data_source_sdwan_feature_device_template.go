package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &FeatureDeviceTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &FeatureDeviceTemplateDataSource{}
)

func NewFeatureDeviceTemplateDataSource() datasource.DataSource {
	return &FeatureDeviceTemplateDataSource{}
}

type FeatureDeviceTemplateDataSource struct {
	client *sdwan.Client
}

func (d *FeatureDeviceTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feature_device_template"
}

func (d *FeatureDeviceTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read a feature device template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the device template",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the device template",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the device template",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Computed:            true,
			},
			"device_type": schema.StringAttribute{
				MarkdownDescription: "The device type (e.g., `vedge-ISR-4331`)",
				Computed:            true,
			},
			"device_role": schema.StringAttribute{
				MarkdownDescription: "The device role",
				Computed:            true,
			},
			"policy_id": schema.StringAttribute{
				MarkdownDescription: "The policy ID",
				Computed:            true,
			},
			"policy_version": schema.Int64Attribute{
				MarkdownDescription: "The policy version",
				Computed:            true,
			},
			"security_policy_id": schema.StringAttribute{
				MarkdownDescription: "The security policy ID",
				Computed:            true,
			},
			"general_templates": schema.SetNestedAttribute{
				MarkdownDescription: "General templates",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Feature template ID",
							Computed:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "The version of the feature template",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Feature template type",
							Computed:            true,
						},
						"sub_templates": schema.SetNestedAttribute{
							MarkdownDescription: "Sub templates",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Feature template ID",
										Computed:            true,
									},
									"version": schema.Int64Attribute{
										MarkdownDescription: "The version of the feature template",
										Computed:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Feature template type",
										Computed:            true,
									},
									"sub_templates": schema.SetNestedAttribute{
										MarkdownDescription: "Sub templates",
										Computed:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: "Feature template ID",
													Computed:            true,
												},
												"version": schema.Int64Attribute{
													MarkdownDescription: "The version of the feature template",
													Computed:            true,
												},
												"type": schema.StringAttribute{
													MarkdownDescription: "Feature template type",
													Computed:            true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *FeatureDeviceTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *FeatureDeviceTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config FeatureDeviceTemplate

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/device/object/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
