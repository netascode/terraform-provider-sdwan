package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &CLIDeviceTemplateDataSource{}
	_ datasource.DataSourceWithConfigure = &CLIDeviceTemplateDataSource{}
)

func NewCLIDeviceTemplateDataSource() datasource.DataSource {
	return &CLIDeviceTemplateDataSource{}
}

type CLIDeviceTemplateDataSource struct {
	client *sdwan.Client
}

func (d *CLIDeviceTemplateDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cli_device_template"
}

func (d *CLIDeviceTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read a CLI device template.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the deivce template",
				Required:            true,
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
			"cli_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI type").String,
				Computed:            true,
			},
			"cli_configuration": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("CLI configuration").String,
				Computed:            true,
			},
		},
	}
}

func (d *CLIDeviceTemplateDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *CLIDeviceTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config CLIDeviceTemplate

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
