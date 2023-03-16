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

type CiscoNTP struct {
	Id                            types.String                 `tfsdk:"id"`
	Version                       types.Int64                  `tfsdk:"version"`
	TemplateType                  types.String                 `tfsdk:"template_type"`
	Name                          types.String                 `tfsdk:"name"`
	Description                   types.String                 `tfsdk:"description"`
	DeviceTypes                   types.List                   `tfsdk:"device_types"`
	Master                        types.Bool                   `tfsdk:"master"`
	MasterVariable                types.String                 `tfsdk:"master_variable"`
	MasterStratum                 types.Int64                  `tfsdk:"master_stratum"`
	MasterStratumVariable         types.String                 `tfsdk:"master_stratum_variable"`
	MasterSourceInterface         types.String                 `tfsdk:"master_source_interface"`
	MasterSourceInterfaceVariable types.String                 `tfsdk:"master_source_interface_variable"`
	TrustedKeys                   types.List                   `tfsdk:"trusted_keys"`
	TrustedKeysVariable           types.String                 `tfsdk:"trusted_keys_variable"`
	AuthenticationKeys            []CiscoNTPAuthenticationKeys `tfsdk:"authentication_keys"`
	Servers                       []CiscoNTPServers            `tfsdk:"servers"`
}

type CiscoNTPAuthenticationKeys struct {
	Optional      types.Bool   `tfsdk:"optional"`
	Id            types.Int64  `tfsdk:"id"`
	IdVariable    types.String `tfsdk:"id_variable"`
	Value         types.String `tfsdk:"value"`
	ValueVariable types.String `tfsdk:"value_variable"`
}

type CiscoNTPServers struct {
	Optional                    types.Bool   `tfsdk:"optional"`
	HostnameIp                  types.String `tfsdk:"hostname_ip"`
	HostnameIpVariable          types.String `tfsdk:"hostname_ip_variable"`
	AuthenticationKeyId         types.Int64  `tfsdk:"authentication_key_id"`
	AuthenticationKeyIdVariable types.String `tfsdk:"authentication_key_id_variable"`
	VpnId                       types.Int64  `tfsdk:"vpn_id"`
	VpnIdVariable               types.String `tfsdk:"vpn_id_variable"`
	Version                     types.Int64  `tfsdk:"version"`
	VersionVariable             types.String `tfsdk:"version_variable"`
	SourceInterface             types.String `tfsdk:"source_interface"`
	SourceInterfaceVariable     types.String `tfsdk:"source_interface_variable"`
	Prefer                      types.Bool   `tfsdk:"prefer"`
	PreferVariable              types.String `tfsdk:"prefer_variable"`
}

func (data CiscoNTP) getModel() string {
	return "cisco_ntp"
}

func (data CiscoNTP) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_ntp")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.MasterVariable.IsNull() {
		body, _ = sjson.Set(body, path+"master.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.enable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"master.enable."+"vipVariableName", data.MasterVariable.ValueString())
	} else if data.Master.IsNull() {
		body, _ = sjson.Set(body, path+"master.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"master.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"master.enable."+"vipValue", strconv.FormatBool(data.Master.ValueBool()))
	}

	if !data.MasterStratumVariable.IsNull() {
		body, _ = sjson.Set(body, path+"master.stratum."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.stratum."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"master.stratum."+"vipVariableName", data.MasterStratumVariable.ValueString())
	} else if data.MasterStratum.IsNull() {
		body, _ = sjson.Set(body, path+"master.stratum."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.stratum."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"master.stratum."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.stratum."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"master.stratum."+"vipValue", data.MasterStratum.ValueInt64())
	}

	if !data.MasterSourceInterfaceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"master.source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.source."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"master.source."+"vipVariableName", data.MasterSourceInterfaceVariable.ValueString())
	} else if data.MasterSourceInterface.IsNull() {
		body, _ = sjson.Set(body, path+"master.source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.source."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"master.source."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"master.source."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"master.source."+"vipValue", data.MasterSourceInterface.ValueString())
	}

	if !data.TrustedKeysVariable.IsNull() {
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipVariableName", data.TrustedKeysVariable.ValueString())
	} else if data.TrustedKeys.IsNull() {
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipType", "constant")
		var values []string
		data.TrustedKeys.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"keys.trusted."+"vipValue", values)
	}
	if len(data.AuthenticationKeys) > 0 {
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipPrimaryKey", []string{"number"})
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipPrimaryKey", []string{"number"})
		body, _ = sjson.Set(body, path+"keys.authentication."+"vipValue", []interface{}{})
	}
	for _, item := range data.AuthenticationKeys {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "number")

		if !item.IdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipVariableName", item.IdVariable.ValueString())
		} else if item.Id.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipValue", item.Id.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "md5")

		if !item.ValueVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "md5."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "md5."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "md5."+"vipVariableName", item.ValueVariable.ValueString())
		} else if item.Value.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "md5."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "md5."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "md5."+"vipValue", item.Value.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"keys.authentication."+"vipValue.-1", itemBody)
	}
	if len(data.Servers) > 0 {
		body, _ = sjson.Set(body, path+"server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"server."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"server."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"server."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"server."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"server."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"server."+"vipValue", []interface{}{})
	}
	for _, item := range data.Servers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.HostnameIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.HostnameIpVariable.ValueString())
		} else if item.HostnameIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.HostnameIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "key")

		if !item.AuthenticationKeyIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "key."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "key."+"vipVariableName", item.AuthenticationKeyIdVariable.ValueString())
		} else if item.AuthenticationKeyId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "key."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "key."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "key."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "key."+"vipValue", item.AuthenticationKeyId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.VpnIdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipVariableName", item.VpnIdVariable.ValueString())
		} else if item.VpnId.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "version")

		if !item.VersionVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "version."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "version."+"vipVariableName", item.VersionVariable.ValueString())
		} else if item.Version.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "version."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "version."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "version."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "version."+"vipValue", item.Version.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "source-interface")

		if !item.SourceInterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipVariableName", item.SourceInterfaceVariable.ValueString())
		} else if item.SourceInterface.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "source-interface."+"vipValue", item.SourceInterface.ValueString())
		}
		itemAttributes = append(itemAttributes, "prefer")

		if !item.PreferVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipVariableName", item.PreferVariable.ValueString())
		} else if item.Prefer.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "prefer."+"vipValue", strconv.FormatBool(item.Prefer.ValueBool()))
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"server."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *CiscoNTP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "master.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Master = types.BoolNull()

			v := res.Get(path + "master.enable.vipVariableName")
			data.MasterVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Master = types.BoolNull()
			data.MasterVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "master.enable.vipValue")
			data.Master = types.BoolValue(v.Bool())
			data.MasterVariable = types.StringNull()
		}
	} else {
		data.Master = types.BoolNull()
		data.MasterVariable = types.StringNull()
	}
	if value := res.Get(path + "master.stratum.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MasterStratum = types.Int64Null()

			v := res.Get(path + "master.stratum.vipVariableName")
			data.MasterStratumVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MasterStratum = types.Int64Null()
			data.MasterStratumVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "master.stratum.vipValue")
			data.MasterStratum = types.Int64Value(v.Int())
			data.MasterStratumVariable = types.StringNull()
		}
	} else {
		data.MasterStratum = types.Int64Null()
		data.MasterStratumVariable = types.StringNull()
	}
	if value := res.Get(path + "master.source.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MasterSourceInterface = types.StringNull()

			v := res.Get(path + "master.source.vipVariableName")
			data.MasterSourceInterfaceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MasterSourceInterface = types.StringNull()
			data.MasterSourceInterfaceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "master.source.vipValue")
			data.MasterSourceInterface = types.StringValue(v.String())
			data.MasterSourceInterfaceVariable = types.StringNull()
		}
	} else {
		data.MasterSourceInterface = types.StringNull()
		data.MasterSourceInterfaceVariable = types.StringNull()
	}
	if value := res.Get(path + "keys.trusted.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.TrustedKeys = types.ListNull(types.StringType)

			v := res.Get(path + "keys.trusted.vipVariableName")
			data.TrustedKeysVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TrustedKeys = types.ListNull(types.StringType)
			data.TrustedKeysVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "keys.trusted.vipValue")
			data.TrustedKeys = helpers.GetStringList(v.Array())
			data.TrustedKeysVariable = types.StringNull()
		}
	} else {
		data.TrustedKeys = types.ListNull(types.StringType)
		data.TrustedKeysVariable = types.StringNull()
	}
	if value := res.Get(path + "keys.authentication.vipValue"); len(value.Array()) > 0 {
		data.AuthenticationKeys = make([]CiscoNTPAuthenticationKeys, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoNTPAuthenticationKeys{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("number.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Id = types.Int64Null()

					cv := v.Get("number.vipVariableName")
					item.IdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Id = types.Int64Null()
					item.IdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("number.vipValue")
					item.Id = types.Int64Value(cv.Int())
					item.IdVariable = types.StringNull()
				}
			} else {
				item.Id = types.Int64Null()
				item.IdVariable = types.StringNull()
			}
			if cValue := v.Get("md5.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Value = types.StringNull()

					cv := v.Get("md5.vipVariableName")
					item.ValueVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Value = types.StringNull()
					item.ValueVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("md5.vipValue")
					item.Value = types.StringValue(cv.String())
					item.ValueVariable = types.StringNull()
				}
			} else {
				item.Value = types.StringNull()
				item.ValueVariable = types.StringNull()
			}
			data.AuthenticationKeys = append(data.AuthenticationKeys, item)
			return true
		})
	}
	if value := res.Get(path + "server.vipValue"); len(value.Array()) > 0 {
		data.Servers = make([]CiscoNTPServers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoNTPServers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.HostnameIp = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.HostnameIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.HostnameIp = types.StringNull()
					item.HostnameIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.HostnameIp = types.StringValue(cv.String())
					item.HostnameIpVariable = types.StringNull()
				}
			} else {
				item.HostnameIp = types.StringNull()
				item.HostnameIpVariable = types.StringNull()
			}
			if cValue := v.Get("key.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AuthenticationKeyId = types.Int64Null()

					cv := v.Get("key.vipVariableName")
					item.AuthenticationKeyIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AuthenticationKeyId = types.Int64Null()
					item.AuthenticationKeyIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("key.vipValue")
					item.AuthenticationKeyId = types.Int64Value(cv.Int())
					item.AuthenticationKeyIdVariable = types.StringNull()
				}
			} else {
				item.AuthenticationKeyId = types.Int64Null()
				item.AuthenticationKeyIdVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

					cv := v.Get("vpn.vipVariableName")
					item.VpnIdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()
					item.VpnIdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())
					item.VpnIdVariable = types.StringNull()
				}
			} else {
				item.VpnId = types.Int64Null()
				item.VpnIdVariable = types.StringNull()
			}
			if cValue := v.Get("version.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Version = types.Int64Null()

					cv := v.Get("version.vipVariableName")
					item.VersionVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Version = types.Int64Null()
					item.VersionVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("version.vipValue")
					item.Version = types.Int64Value(cv.Int())
					item.VersionVariable = types.StringNull()
				}
			} else {
				item.Version = types.Int64Null()
				item.VersionVariable = types.StringNull()
			}
			if cValue := v.Get("source-interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.SourceInterface = types.StringNull()

					cv := v.Get("source-interface.vipVariableName")
					item.SourceInterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.SourceInterface = types.StringNull()
					item.SourceInterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("source-interface.vipValue")
					item.SourceInterface = types.StringValue(cv.String())
					item.SourceInterfaceVariable = types.StringNull()
				}
			} else {
				item.SourceInterface = types.StringNull()
				item.SourceInterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("prefer.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Prefer = types.BoolNull()

					cv := v.Get("prefer.vipVariableName")
					item.PreferVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Prefer = types.BoolNull()
					item.PreferVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("prefer.vipValue")
					item.Prefer = types.BoolValue(cv.Bool())
					item.PreferVariable = types.StringNull()
				}
			} else {
				item.Prefer = types.BoolNull()
				item.PreferVariable = types.StringNull()
			}
			data.Servers = append(data.Servers, item)
			return true
		})
	}
}
