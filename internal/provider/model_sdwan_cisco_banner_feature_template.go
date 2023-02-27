// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CiscoBanner struct {
	Id            types.String `tfsdk:"id"`
	TemplateType  types.String `tfsdk:"template_type"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	DeviceTypes   types.List   `tfsdk:"device_types"`
	Login         types.String `tfsdk:"login"`
	LoginVariable types.String `tfsdk:"login_variable"`
	Motd          types.String `tfsdk:"motd"`
	MotdVariable  types.String `tfsdk:"motd_variable"`
}

func (data CiscoBanner) getModel() string {
	return "cisco_banner"
}

func (data CiscoBanner) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_banner")

	path := "templateDefinition."
	body, _ = sjson.Set(body, path+"login."+"vipObjectType", "object")

	if !data.LoginVariable.IsNull() {
		body, _ = sjson.Set(body, path+"login."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"login."+"vipVariableName", data.LoginVariable.ValueString())
	} else if data.Login.IsNull() {
		body, _ = sjson.Set(body, path+"login."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"login."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"login."+"vipValue", data.Login.ValueString())
	}
	body, _ = sjson.Set(body, path+"motd."+"vipObjectType", "object")

	if !data.MotdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"motd."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"motd."+"vipVariableName", data.MotdVariable.ValueString())
	} else if data.Motd.IsNull() {
		body, _ = sjson.Set(body, path+"motd."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"motd."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"motd."+"vipValue", data.Motd.ValueString())
	}
	return body
}

func (data *CiscoBanner) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringList(value.Array())
	} else {
		data.DeviceTypes = types.ListNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "login.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Login = types.StringNull()

			v := res.Get(path + "login.vipVariableName")
			data.LoginVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Login = types.StringNull()
			data.LoginVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "login.vipValue")
			data.Login = types.StringValue(v.String())
			data.LoginVariable = types.StringNull()
		}
	} else {
		data.Login = types.StringNull()
		data.LoginVariable = types.StringNull()
	}
	if value := res.Get(path + "motd.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Motd = types.StringNull()

			v := res.Get(path + "motd.vipVariableName")
			data.MotdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Motd = types.StringNull()
			data.MotdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "motd.vipValue")
			data.Motd = types.StringValue(v.String())
			data.MotdVariable = types.StringNull()
		}
	} else {
		data.Motd = types.StringNull()
		data.MotdVariable = types.StringNull()
	}
}
