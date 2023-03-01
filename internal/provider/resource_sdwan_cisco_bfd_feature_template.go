// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &CiscoBFDFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoBFDFeatureTemplateResource{}

func NewCiscoBFDFeatureTemplateResource() resource.Resource {
	return &CiscoBFDFeatureTemplateResource{}
}

type CiscoBFDFeatureTemplateResource struct {
	client *sdwan.Client
}

func (r *CiscoBFDFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_bfd_feature_template"
}

func (r *CiscoBFDFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco BFD feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
				Required:            true,
			},
			"multiplier": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set the number of polling intervals used to determine tunnel SLA class").AddIntegerRangeDescription(1, 6).AddDefaultValueDescription("6").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 6),
				},
			},
			"multiplier_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"poll_interval": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set how often BFD polls tunnels to collect packet latency, loss, and statistics").AddIntegerRangeDescription(1, 4294967295).AddDefaultValueDescription("600000").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967295),
				},
			},
			"poll_interval_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_dscp": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set default DSCP value for BFD packets").AddIntegerRangeDescription(0, 63).AddDefaultValueDescription("48").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 63),
				},
			},
			"default_dscp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"colors": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set color that identifies the WAN transport tunnel").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"color": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set color that identifies the WAN transport tunnel").AddStringEnumDescription("default", "mpls", "metro-ethernet", "biz-internet", "public-internet", "lte", "3g", "red", "green", "blue", "gold", "silver", "bronze", "custom1", "custom2", "custom3", "private1", "private2", "private3", "private4", "private5", "private6").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("default", "mpls", "metro-ethernet", "biz-internet", "public-internet", "lte", "3g", "red", "green", "blue", "gold", "silver", "bronze", "custom1", "custom2", "custom3", "private1", "private2", "private3", "private4", "private5", "private6"),
							},
						},
						"color_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"hello_interval": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how often BFD sends Hello packets").AddIntegerRangeDescription(100, 300000).AddDefaultValueDescription("1000").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(100, 300000),
							},
						},
						"hello_interval_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"multiplier": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how many Hello packet intervals to wait before declaring that a tunnel has failed").AddIntegerRangeDescription(1, 60).AddDefaultValueDescription("7").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(1, 60),
							},
						},
						"multiplier_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"pmtu_discovery": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Control automatic path MTU discovery").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"pmtu_discovery_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"dscp": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set BFD Default DSCP value for tloc color").AddIntegerRangeDescription(0, 63).AddDefaultValueDescription("48").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 63),
							},
						},
						"dscp_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *CiscoBFDFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*sdwan.Client)
}

func (r *CiscoBFDFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoBFD

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CiscoBFDFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoBFD

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + state.Id.ValueString())
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *CiscoBFDFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CiscoBFD

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put("/template/feature/"+plan.Id.ValueString(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CiscoBFDFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoBFD

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + state.Id.ValueString())
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *CiscoBFDFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
