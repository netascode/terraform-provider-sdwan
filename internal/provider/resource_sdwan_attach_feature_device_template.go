package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &AttachFeatureDeviceTemplateResource{}

func NewAttachFeatureDeviceTemplateResource() resource.Resource {
	return &AttachFeatureDeviceTemplateResource{}
}

type AttachFeatureDeviceTemplateResource struct {
	client *sdwan.Client
}

func (r *AttachFeatureDeviceTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_attach_feature_device_template"
}

func (r *AttachFeatureDeviceTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can attach a feature device template.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The ID of the device template",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the device template",
				Optional:            true,
			},
			"devices": schema.ListNestedAttribute{
				MarkdownDescription: "Devices",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Device ID",
							Required:            true,
						},
						"variables": schema.MapAttribute{
							MarkdownDescription: "Device variables",
							ElementType:         types.StringType,
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *AttachFeatureDeviceTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*sdwan.Client)
}

func (r *AttachFeatureDeviceTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AttachFeatureDeviceTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body, err := plan.toBody(ctx, r.client, false)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to render body, got error: %s", err))
		return
	}
	res, err := r.client.Post("/template/device/config/attachfeature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	actionId := res.Get("id").String()
	err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to attach device template, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AttachFeatureDeviceTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AttachFeatureDeviceTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	err := state.readVariables(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *AttachFeatureDeviceTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AttachFeatureDeviceTemplate

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body, err := plan.toBody(ctx, r.client, true)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to render body, got error: %s", err))
		return
	}
	res, err := r.client.Post("/template/device/config/attachfeature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	actionId := res.Get("id").String()
	err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to attach device template, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AttachFeatureDeviceTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AttachFeatureDeviceTemplate

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))

	res, err := state.detachDevices(ctx, r.client)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve currently attached devices, got error: %s", err))
		return
	}
	actionId := res.Get("id").String()
	err = helpers.WaitForActionToComplete(ctx, r.client, actionId)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to detach device template, got error: %s", err))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}
