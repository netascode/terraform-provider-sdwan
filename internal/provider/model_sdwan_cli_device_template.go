package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CLIDeviceTemplate struct {
	Id               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	DeviceType       types.String `tfsdk:"device_type"`
	CLIType          types.String `tfsdk:"cli_type"`
	CLIConfiguration types.String `tfsdk:"cli_configuration"`
}

func (data CLIDeviceTemplate) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	body, _ = sjson.Set(body, "templateConfiguration", data.CLIConfiguration.ValueString())
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "configType", "file")
	body, _ = sjson.Set(body, "cliType", data.CLIType.ValueString())
	body, _ = sjson.Set(body, "draftMode", false)
	return body
}

func (data *CLIDeviceTemplate) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceType = types.StringValue(value.String())
	} else {
		data.DeviceType = types.StringNull()
	}
	if value := res.Get("templateConfiguration"); value.Exists() {
		data.CLIConfiguration = types.StringValue(value.String())
	} else {
		data.CLIConfiguration = types.StringNull()
	}
	if value := res.Get("cliType"); value.Exists() {
		data.CLIType = types.StringValue(value.String())
	} else {
		data.CLIType = types.StringNull()
	}
}
