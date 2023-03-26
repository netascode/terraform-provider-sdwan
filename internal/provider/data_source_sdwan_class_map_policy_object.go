// Code generated by "gen/generator.go"; DO NOT EDIT.

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
	_ datasource.DataSource              = &ClassMapPolicyObjectDataSource{}
	_ datasource.DataSourceWithConfigure = &ClassMapPolicyObjectDataSource{}
)

func NewClassMapPolicyObjectDataSource() datasource.DataSource {
	return &ClassMapPolicyObjectDataSource{}
}

type ClassMapPolicyObjectDataSource struct {
	client *sdwan.Client
}

func (d *ClassMapPolicyObjectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_class_map_policy_object"
}

func (d *ClassMapPolicyObjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Class Map policy object.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the policy object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the policy object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy object",
				Computed:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: "List of entries, only 1 entry supported",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"queue": schema.Int64Attribute{
							MarkdownDescription: "Queue",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ClassMapPolicyObjectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *ClassMapPolicyObjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ClassMap

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/policy/list/class/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
