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
var _ resource.Resource = &CiscoLoggingFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoLoggingFeatureTemplateResource{}

func NewCiscoLoggingFeatureTemplateResource() resource.Resource {
	return &CiscoLoggingFeatureTemplateResource{}
}

type CiscoLoggingFeatureTemplateResource struct {
	client *sdwan.Client
}

func (r *CiscoLoggingFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_logging_feature_template"
}

func (r *CiscoLoggingFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco Logging feature template.").AddMinimumVersionDescription("15.0.0").String,

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
			"disk_logging": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging to local disk").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"disk_logging_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"max_size": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum size of file before it is rotated").AddIntegerRangeDescription(1, 20).AddDefaultValueDescription("10").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
			},
			"max_size_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"log_rotations": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set number of syslog files to create before discarding oldest files").AddIntegerRangeDescription(1, 10).AddDefaultValueDescription("10").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
			},
			"log_rotations_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"tls_profiles": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure a TLS profile").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Specify the name of the TLS profile").String,
							Optional:            true,
						},
						"name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("TLS Version").AddStringEnumDescription("TLSv1.1", "TLSv1.2").AddDefaultValueDescription("TLSv1.1").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("TLSv1.1", "TLSv1.2"),
							},
						},
						"version_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"authentication_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Authentication Type").AddStringEnumDescription("Server", "Mutual").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("Server", "Mutual"),
							},
						},
						"ciphersuite_list": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Syslog secure server ciphersuites").AddStringEnumDescription("aes-128-cbc-sha", "aes-256-cbc-sha", "dhe-aes-cbc-sha2", "dhe-aes-gcm-sha2", "ecdhe-ecdsa-aes-gcm-sha2", "ecdhe-rsa-aes-cbc-sha2", "ecdhe-rsa-aes-gcm-sha2", "rsa-aes-cbc-sha2", "rsa-aes-gcm-sha2").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("aes-128-cbc-sha", "aes-256-cbc-sha", "dhe-aes-cbc-sha2", "dhe-aes-gcm-sha2", "ecdhe-ecdsa-aes-gcm-sha2", "ecdhe-rsa-aes-cbc-sha2", "ecdhe-rsa-aes-gcm-sha2", "rsa-aes-cbc-sha2", "rsa-aes-gcm-sha2"),
							},
						},
						"ciphersuite_list_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"ipv4_servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging to remote server").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set hostname or IPv4 address of server").String,
							Optional:            true,
						},
						"hostname_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VPN in which syslog server is located").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65530),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface to use to reach syslog server").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"logging_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set logging level for messages logged to server").AddStringEnumDescription("information", "debugging", "notice", "warn", "error", "critical", "alert", "emergency").AddDefaultValueDescription("information").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("information", "debugging", "notice", "warn", "error", "critical", "alert", "emergency"),
							},
						},
						"logging_level_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"enable_tls": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable TLS").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"enable_tls_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"custom_profile": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define custom profile").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"custom_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure a TLS profile").String,
							Optional:            true,
						},
						"profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"ipv6_servers": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Enable logging to remote IPv6 server").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hostname_ip": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set IPv6 hostname or IPv6 address of server").String,
							Optional:            true,
						},
						"hostname_ip_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set VPN in which syslog server is located").AddIntegerRangeDescription(0, 65530).AddDefaultValueDescription("0").String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 65530),
							},
						},
						"vpn_id_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set interface to use to reach syslog server").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(0, 32),
							},
						},
						"source_interface_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"logging_level": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set logging level for messages logged to server").AddStringEnumDescription("information", "debugging", "notification", "warn", "error", "critical", "alert", "emergency").AddDefaultValueDescription("information").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("information", "debugging", "notification", "warn", "error", "critical", "alert", "emergency"),
							},
						},
						"logging_level_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"enable_tls": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable TLS").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"enable_tls_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"custom_profile": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Define custom profile").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"custom_profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"profile": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Configure a TLS profile").String,
							Optional:            true,
						},
						"profile_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *CiscoLoggingFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*sdwan.Client)
}

func (r *CiscoLoggingFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoLogging

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

func (r *CiscoLoggingFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoLogging

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

func (r *CiscoLoggingFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CiscoLogging

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

func (r *CiscoLoggingFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoLogging

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

func (r *CiscoLoggingFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
