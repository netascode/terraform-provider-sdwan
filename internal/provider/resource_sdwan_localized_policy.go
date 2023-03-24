package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &LocalizedPolicyResource{}
var _ resource.ResourceWithImportState = &LocalizedPolicyResource{}

func NewLocalizedPolicyResource() resource.Resource {
	return &LocalizedPolicyResource{}
}

type LocalizedPolicyResource struct {
	client *sdwan.Client
}

func (r *LocalizedPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_localized_policy"
}

func (r *LocalizedPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a localized policy.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the localized policy",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the localized policy",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the localized policy",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the localized policy",
				Required:            true,
			},
			"flow_visibility_ipv4": schema.BoolAttribute{
				MarkdownDescription: "IPv4 flow visibilty",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"flow_visibility_ipv6": schema.BoolAttribute{
				MarkdownDescription: "IPv6 flow visibilty",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"application_visibility_ipv4": schema.BoolAttribute{
				MarkdownDescription: "IPv4 application visibilty",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"application_visibility_ipv6": schema.BoolAttribute{
				MarkdownDescription: "IPv6 application visibilty",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"cloud_qos": schema.BoolAttribute{
				MarkdownDescription: "Cloud QoS",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"cloud_qos_service_side": schema.BoolAttribute{
				MarkdownDescription: "Cloud QoS service side",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"implicit_acl_logging": schema.BoolAttribute{
				MarkdownDescription: "Implicit ACL logging",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
			"log_frequency": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Log frequency").AddIntegerRangeDescription(1, 2147483647).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 2147483647),
				},
			},
			"ipv4_visibility_cache_entries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv4 visibility cache entries").AddIntegerRangeDescription(16, 2000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(16, 2000000),
				},
			},
			"ipv6_visibility_cache_entries": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("IPv6 visibility cache entries").AddIntegerRangeDescription(16, 2000000).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(16, 2000000),
				},
			},
			"definitions": schema.SetNestedAttribute{
				MarkdownDescription: "List of policy definitions",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Policy definition ID",
							Required:            true,
						},
						"version": schema.Int64Attribute{
							MarkdownDescription: "Policy definition version",
							Optional:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "Policy definiton type",
							Required:            true,
						},
					},
				},
			},
		},
	}
}

func (r *LocalizedPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*sdwan.Client)
}

func (r *LocalizedPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LocalizedPolicy

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/policy/vedge", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("policyId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *LocalizedPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, oldState LocalizedPolicy

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

	res, err := r.client.Get("/template/policy/vedge/definition/" + state.Id.ValueString())
	if res.Get("error.message").String() == "Policy definition not found" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	state.updateDefinitionVersions(ctx, oldState)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *LocalizedPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state LocalizedPolicy

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
	res, err := r.client.Put("/template/policy/vedge/"+plan.Id.ValueString(), body)
	if err != nil && res.Get("error.message").String() != "Failed to update variables" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *LocalizedPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LocalizedPolicy

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/policy/vedge/" + state.Id.ValueString())
	if err != nil && res.Get("error.message").String() != "Policy definition not found" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *LocalizedPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
