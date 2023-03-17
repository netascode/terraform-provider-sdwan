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
var _ resource.Resource = &QoSMapPolicyDefinitionResource{}
var _ resource.ResourceWithImportState = &QoSMapPolicyDefinitionResource{}

func NewQoSMapPolicyDefinitionResource() resource.Resource {
	return &QoSMapPolicyDefinitionResource{}
}

type QoSMapPolicyDefinitionResource struct {
	client *sdwan.Client
}

func (r *QoSMapPolicyDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_qos_map_policy_definition"
}

func (r *QoSMapPolicyDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a QoS Map policy object.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the policy definition",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the policy definition",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Required:            true,
			},
			"qos_schedulers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of QoS schedulers").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"queue": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Queue number").AddIntegerRangeDescription(0, 7).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 7),
							},
						},
						"class_map_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Class map ID").String,
							Required:            true,
						},
						"class_map_version": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Class map version").String,
							Optional:            true,
						},
						"bandwidth_percent": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Bandwidth percent").AddIntegerRangeDescription(0, 100).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 100),
							},
						},
						"buffer_percent": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Buffer percent").AddIntegerRangeDescription(0, 100).String,
							Required:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 100),
							},
						},
						"burst": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Burst size").AddIntegerRangeDescription(5000, 10000000).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(5000, 10000000),
							},
						},
						"drop_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Drop type").AddStringEnumDescription("tail-drop", "red-drop").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("tail-drop", "red-drop"),
							},
						},
						"scheduling_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Scheduling type").AddStringEnumDescription("llq", "wrr").String,
							Required:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("llq", "wrr"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *QoSMapPolicyDefinitionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*sdwan.Client)
}

func (r *QoSMapPolicyDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan QoSMap

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/policy/definition/qosmap", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("definitionId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *QoSMapPolicyDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, oldState QoSMap

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

	res, err := r.client.Get("/template/policy/definition/qosmap/" + state.Id.ValueString())
	if res.Get("error.message").String() == "Failed to find specified resource" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	state.updateVersions(ctx, oldState)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *QoSMapPolicyDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state QoSMap

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

	body := plan.toBody(ctx)
	res, err := r.client.Put("/template/policy/definition/qosmap/"+plan.Id.ValueString(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *QoSMapPolicyDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state QoSMap

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/policy/definition/qosmap/" + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *QoSMapPolicyDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}