package provider

import (
	"context"
	"fmt"

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
var _ resource.Resource = &FeatureDeviceTemplateResource{}
var _ resource.ResourceWithImportState = &FeatureDeviceTemplateResource{}

func NewFeatureDeviceTemplateResource() resource.Resource {
	return &FeatureDeviceTemplateResource{}
}

type FeatureDeviceTemplateResource struct {
	client *sdwan.Client
}

func (r *FeatureDeviceTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feature_device_template"
}

func (r *FeatureDeviceTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a feature device template.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the device template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the device template",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the device template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the device template",
				Required:            true,
			},
			"device_type": schema.StringAttribute{
				MarkdownDescription: "The device type (e.g., `vedge-ISR-4331`)",
				Required:            true,
			},
			"device_role": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("The device role").AddStringEnumDescription("sdwan-edge", "service-node").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("interface", "pool", "loopback"),
				},
			},
			"policy_id": schema.StringAttribute{
				MarkdownDescription: "The policy ID",
				Optional:            true,
			},
			"policy_version": schema.Int64Attribute{
				MarkdownDescription: "The policy version",
				Optional:            true,
			},
			"security_policy_id": schema.StringAttribute{
				MarkdownDescription: "The security policy ID",
				Optional:            true,
			},
			"general_templates": schema.SetNestedAttribute{
				MarkdownDescription: "General templates",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Feature template ID",
							Required:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "Feature template version",
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Feature template type",
							Required:            true,
						},
						"sub_templates": schema.SetNestedAttribute{
							MarkdownDescription: "Sub templates",
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										MarkdownDescription: "Feature template ID",
										Required:            true,
									},
									"version": schema.Int64Attribute{
										MarkdownDescription: "Feature template version",
										Optional:            true,
									},
									"type": schema.StringAttribute{
										MarkdownDescription: "Feature template type",
										Required:            true,
									},
									"sub_templates": schema.SetNestedAttribute{
										MarkdownDescription: "Sub templates",
										Optional:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													MarkdownDescription: "Feature template ID",
													Required:            true,
												},
												"version": schema.Int64Attribute{
													MarkdownDescription: "Feature template version",
													Optional:            true,
												},
												"type": schema.StringAttribute{
													MarkdownDescription: "Feature template type",
													Required:            true,
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

func (r *FeatureDeviceTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*sdwan.Client)
}

func (r *FeatureDeviceTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FeatureDeviceTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/device/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FeatureDeviceTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, oldState FeatureDeviceTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.State.Get(ctx, &oldState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/device/object/" + state.Id.ValueString())
	if res.Get("error.message").String() == "Template definition not found" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	state.updateTemplateVersions(ctx, oldState)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *FeatureDeviceTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FeatureDeviceTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		res, err := r.client.Put("/template/device/"+plan.Id.ValueString(), body)
		if err != nil && res.Get("error.message").String() != "Template locked in edit mode." {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", plan.Name.ValueString()))
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FeatureDeviceTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FeatureDeviceTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/device/" + state.Id.ValueString())
	if err != nil && res.Get("error.message").String() != "Template definition not found" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *FeatureDeviceTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
