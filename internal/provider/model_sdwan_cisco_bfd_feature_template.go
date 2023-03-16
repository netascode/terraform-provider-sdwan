// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CiscoBFD struct {
	Id                   types.String     `tfsdk:"id"`
	Version              types.Int64      `tfsdk:"version"`
	TemplateType         types.String     `tfsdk:"template_type"`
	Name                 types.String     `tfsdk:"name"`
	Description          types.String     `tfsdk:"description"`
	DeviceTypes          types.List       `tfsdk:"device_types"`
	Multiplier           types.Int64      `tfsdk:"multiplier"`
	MultiplierVariable   types.String     `tfsdk:"multiplier_variable"`
	PollInterval         types.Int64      `tfsdk:"poll_interval"`
	PollIntervalVariable types.String     `tfsdk:"poll_interval_variable"`
	DefaultDscp          types.Int64      `tfsdk:"default_dscp"`
	DefaultDscpVariable  types.String     `tfsdk:"default_dscp_variable"`
	Colors               []CiscoBFDColors `tfsdk:"colors"`
}

type CiscoBFDColors struct {
	Optional              types.Bool   `tfsdk:"optional"`
	Color                 types.String `tfsdk:"color"`
	ColorVariable         types.String `tfsdk:"color_variable"`
	HelloInterval         types.Int64  `tfsdk:"hello_interval"`
	HelloIntervalVariable types.String `tfsdk:"hello_interval_variable"`
	Multiplier            types.Int64  `tfsdk:"multiplier"`
	MultiplierVariable    types.String `tfsdk:"multiplier_variable"`
	PmtuDiscovery         types.Bool   `tfsdk:"pmtu_discovery"`
	PmtuDiscoveryVariable types.String `tfsdk:"pmtu_discovery_variable"`
	Dscp                  types.Int64  `tfsdk:"dscp"`
	DscpVariable          types.String `tfsdk:"dscp_variable"`
}

func (data CiscoBFD) getModel() string {
	return "cisco_bfd"
}

func (data CiscoBFD) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_bfd")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.MultiplierVariable.IsNull() {
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipVariableName", data.MultiplierVariable.ValueString())
	} else if data.Multiplier.IsNull() {
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"app-route.multiplier."+"vipValue", data.Multiplier.ValueInt64())
	}

	if !data.PollIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipVariableName", data.PollIntervalVariable.ValueString())
	} else if data.PollInterval.IsNull() {
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"app-route.poll-interval."+"vipValue", data.PollInterval.ValueInt64())
	}

	if !data.DefaultDscpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"default-dscp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"default-dscp."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"default-dscp."+"vipVariableName", data.DefaultDscpVariable.ValueString())
	} else if data.DefaultDscp.IsNull() {
		body, _ = sjson.Set(body, path+"default-dscp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"default-dscp."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"default-dscp."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"default-dscp."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"default-dscp."+"vipValue", data.DefaultDscp.ValueInt64())
	}
	if len(data.Colors) > 0 {
		body, _ = sjson.Set(body, path+"color."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"color."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"color."+"vipPrimaryKey", []string{"color"})
		body, _ = sjson.Set(body, path+"color."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"color."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"color."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"color."+"vipPrimaryKey", []string{"color"})
		body, _ = sjson.Set(body, path+"color."+"vipValue", []interface{}{})
	}
	for _, item := range data.Colors {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "color")

		if !item.ColorVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "color."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "color."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "color."+"vipVariableName", item.ColorVariable.ValueString())
		} else if item.Color.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "color."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "color."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "color."+"vipValue", item.Color.ValueString())
		}
		itemAttributes = append(itemAttributes, "hello-interval")

		if !item.HelloIntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipVariableName", item.HelloIntervalVariable.ValueString())
		} else if item.HelloInterval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "hello-interval."+"vipValue", item.HelloInterval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "multiplier")

		if !item.MultiplierVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipVariableName", item.MultiplierVariable.ValueString())
		} else if item.Multiplier.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipValue", item.Multiplier.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "pmtu-discovery")

		if !item.PmtuDiscoveryVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipVariableName", item.PmtuDiscoveryVariable.ValueString())
		} else if item.PmtuDiscovery.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "pmtu-discovery."+"vipValue", strconv.FormatBool(item.PmtuDiscovery.ValueBool()))
		}
		itemAttributes = append(itemAttributes, "dscp")

		if !item.DscpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipVariableName", item.DscpVariable.ValueString())
		} else if item.Dscp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "dscp."+"vipValue", item.Dscp.ValueInt64())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"color."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *CiscoBFD) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "app-route.multiplier.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Multiplier = types.Int64Null()

			v := res.Get(path + "app-route.multiplier.vipVariableName")
			data.MultiplierVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Multiplier = types.Int64Null()
			data.MultiplierVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "app-route.multiplier.vipValue")
			data.Multiplier = types.Int64Value(v.Int())
			data.MultiplierVariable = types.StringNull()
		}
	} else {
		data.Multiplier = types.Int64Null()
		data.MultiplierVariable = types.StringNull()
	}
	if value := res.Get(path + "app-route.poll-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PollInterval = types.Int64Null()

			v := res.Get(path + "app-route.poll-interval.vipVariableName")
			data.PollIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PollInterval = types.Int64Null()
			data.PollIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "app-route.poll-interval.vipValue")
			data.PollInterval = types.Int64Value(v.Int())
			data.PollIntervalVariable = types.StringNull()
		}
	} else {
		data.PollInterval = types.Int64Null()
		data.PollIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "default-dscp.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.DefaultDscp = types.Int64Null()

			v := res.Get(path + "default-dscp.vipVariableName")
			data.DefaultDscpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DefaultDscp = types.Int64Null()
			data.DefaultDscpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "default-dscp.vipValue")
			data.DefaultDscp = types.Int64Value(v.Int())
			data.DefaultDscpVariable = types.StringNull()
		}
	} else {
		data.DefaultDscp = types.Int64Null()
		data.DefaultDscpVariable = types.StringNull()
	}
	if value := res.Get(path + "color.vipValue"); len(value.Array()) > 0 {
		data.Colors = make([]CiscoBFDColors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoBFDColors{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("color.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Color = types.StringNull()

					cv := v.Get("color.vipVariableName")
					item.ColorVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Color = types.StringNull()
					item.ColorVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("color.vipValue")
					item.Color = types.StringValue(cv.String())
					item.ColorVariable = types.StringNull()
				}
			} else {
				item.Color = types.StringNull()
				item.ColorVariable = types.StringNull()
			}
			if cValue := v.Get("hello-interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.HelloInterval = types.Int64Null()

					cv := v.Get("hello-interval.vipVariableName")
					item.HelloIntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.HelloInterval = types.Int64Null()
					item.HelloIntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("hello-interval.vipValue")
					item.HelloInterval = types.Int64Value(cv.Int())
					item.HelloIntervalVariable = types.StringNull()
				}
			} else {
				item.HelloInterval = types.Int64Null()
				item.HelloIntervalVariable = types.StringNull()
			}
			if cValue := v.Get("multiplier.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Multiplier = types.Int64Null()

					cv := v.Get("multiplier.vipVariableName")
					item.MultiplierVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Multiplier = types.Int64Null()
					item.MultiplierVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("multiplier.vipValue")
					item.Multiplier = types.Int64Value(cv.Int())
					item.MultiplierVariable = types.StringNull()
				}
			} else {
				item.Multiplier = types.Int64Null()
				item.MultiplierVariable = types.StringNull()
			}
			if cValue := v.Get("pmtu-discovery.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.PmtuDiscovery = types.BoolNull()

					cv := v.Get("pmtu-discovery.vipVariableName")
					item.PmtuDiscoveryVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.PmtuDiscovery = types.BoolNull()
					item.PmtuDiscoveryVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("pmtu-discovery.vipValue")
					item.PmtuDiscovery = types.BoolValue(cv.Bool())
					item.PmtuDiscoveryVariable = types.StringNull()
				}
			} else {
				item.PmtuDiscovery = types.BoolNull()
				item.PmtuDiscoveryVariable = types.StringNull()
			}
			if cValue := v.Get("dscp.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Dscp = types.Int64Null()

					cv := v.Get("dscp.vipVariableName")
					item.DscpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Dscp = types.Int64Null()
					item.DscpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("dscp.vipValue")
					item.Dscp = types.Int64Value(cv.Int())
					item.DscpVariable = types.StringNull()
				}
			} else {
				item.Dscp = types.Int64Null()
				item.DscpVariable = types.StringNull()
			}
			data.Colors = append(data.Colors, item)
			return true
		})
	}
}
