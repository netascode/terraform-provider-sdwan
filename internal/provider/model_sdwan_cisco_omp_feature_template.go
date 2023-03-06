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

type CiscoOMP struct {
	Id                             types.String                  `tfsdk:"id"`
	TemplateType                   types.String                  `tfsdk:"template_type"`
	Name                           types.String                  `tfsdk:"name"`
	Description                    types.String                  `tfsdk:"description"`
	DeviceTypes                    types.List                    `tfsdk:"device_types"`
	GracefulRestart                types.Bool                    `tfsdk:"graceful_restart"`
	GracefulRestartVariable        types.String                  `tfsdk:"graceful_restart_variable"`
	OverlayAs                      types.Int64                   `tfsdk:"overlay_as"`
	OverlayAsVariable              types.String                  `tfsdk:"overlay_as_variable"`
	SendPathLimit                  types.Int64                   `tfsdk:"send_path_limit"`
	SendPathLimitVariable          types.String                  `tfsdk:"send_path_limit_variable"`
	EcmpLimit                      types.Int64                   `tfsdk:"ecmp_limit"`
	EcmpLimitVariable              types.String                  `tfsdk:"ecmp_limit_variable"`
	Shutdown                       types.Bool                    `tfsdk:"shutdown"`
	ShutdownVariable               types.String                  `tfsdk:"shutdown_variable"`
	OmpAdminDistanceIpv4           types.Int64                   `tfsdk:"omp_admin_distance_ipv4"`
	OmpAdminDistanceIpv4Variable   types.String                  `tfsdk:"omp_admin_distance_ipv4_variable"`
	OmpAdminDistanceIpv6           types.Int64                   `tfsdk:"omp_admin_distance_ipv6"`
	OmpAdminDistanceIpv6Variable   types.String                  `tfsdk:"omp_admin_distance_ipv6_variable"`
	AdvertisementInterval          types.Int64                   `tfsdk:"advertisement_interval"`
	AdvertisementIntervalVariable  types.String                  `tfsdk:"advertisement_interval_variable"`
	GracefulRestartTimer           types.Int64                   `tfsdk:"graceful_restart_timer"`
	GracefulRestartTimerVariable   types.String                  `tfsdk:"graceful_restart_timer_variable"`
	EorTimer                       types.Int64                   `tfsdk:"eor_timer"`
	EorTimerVariable               types.String                  `tfsdk:"eor_timer_variable"`
	Holdtime                       types.Int64                   `tfsdk:"holdtime"`
	HoldtimeVariable               types.String                  `tfsdk:"holdtime_variable"`
	IgnoreRegionPathLength         types.Bool                    `tfsdk:"ignore_region_path_length"`
	IgnoreRegionPathLengthVariable types.String                  `tfsdk:"ignore_region_path_length_variable"`
	TransportGateway               types.String                  `tfsdk:"transport_gateway"`
	TransportGatewayVariable       types.String                  `tfsdk:"transport_gateway_variable"`
	AdvertiseIpv4Routes            []CiscoOMPAdvertiseIpv4Routes `tfsdk:"advertise_ipv4_routes"`
	AdvertiseIpv6Routes            []CiscoOMPAdvertiseIpv6Routes `tfsdk:"advertise_ipv6_routes"`
}

type CiscoOMPAdvertiseIpv4Routes struct {
	Protocol                      types.String `tfsdk:"protocol"`
	AdvertiseExternalOspf         types.String `tfsdk:"advertise_external_ospf"`
	AdvertiseExternalOspfVariable types.String `tfsdk:"advertise_external_ospf_variable"`
}

type CiscoOMPAdvertiseIpv6Routes struct {
	Protocol types.String `tfsdk:"protocol"`
}

func (data CiscoOMP) getModel() string {
	return "cisco_omp"
}

func (data CiscoOMP) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_omp")

	path := "templateDefinition."
	body, _ = sjson.Set(body, path+"graceful-restart."+"vipObjectType", "object")

	if !data.GracefulRestartVariable.IsNull() {
		body, _ = sjson.Set(body, path+"graceful-restart."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"graceful-restart."+"vipVariableName", data.GracefulRestartVariable.ValueString())
	} else if data.GracefulRestart.IsNull() {
		body, _ = sjson.Set(body, path+"graceful-restart."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"graceful-restart."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"graceful-restart."+"vipValue", strconv.FormatBool(data.GracefulRestart.ValueBool()))
	}
	body, _ = sjson.Set(body, path+"overlay-as."+"vipObjectType", "object")

	if !data.OverlayAsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"overlay-as."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"overlay-as."+"vipVariableName", data.OverlayAsVariable.ValueString())
	} else if data.OverlayAs.IsNull() {
		body, _ = sjson.Set(body, path+"overlay-as."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"overlay-as."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"overlay-as."+"vipValue", data.OverlayAs.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"send-path-limit."+"vipObjectType", "object")

	if !data.SendPathLimitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"send-path-limit."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"send-path-limit."+"vipVariableName", data.SendPathLimitVariable.ValueString())
	} else if data.SendPathLimit.IsNull() {
		body, _ = sjson.Set(body, path+"send-path-limit."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"send-path-limit."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"send-path-limit."+"vipValue", data.SendPathLimit.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"ecmp-limit."+"vipObjectType", "object")

	if !data.EcmpLimitVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ecmp-limit."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ecmp-limit."+"vipVariableName", data.EcmpLimitVariable.ValueString())
	} else if data.EcmpLimit.IsNull() {
		body, _ = sjson.Set(body, path+"ecmp-limit."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ecmp-limit."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ecmp-limit."+"vipValue", data.EcmpLimit.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"shutdown."+"vipObjectType", "object")

	if !data.ShutdownVariable.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"shutdown."+"vipVariableName", data.ShutdownVariable.ValueString())
	} else if data.Shutdown.IsNull() {
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"shutdown."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"shutdown."+"vipValue", strconv.FormatBool(data.Shutdown.ValueBool()))
	}
	body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipObjectType", "object")

	if !data.OmpAdminDistanceIpv4Variable.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipVariableName", data.OmpAdminDistanceIpv4Variable.ValueString())
	} else if data.OmpAdminDistanceIpv4.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv4."+"vipValue", data.OmpAdminDistanceIpv4.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipObjectType", "object")

	if !data.OmpAdminDistanceIpv6Variable.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipVariableName", data.OmpAdminDistanceIpv6Variable.ValueString())
	} else if data.OmpAdminDistanceIpv6.IsNull() {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"omp-admin-distance-ipv6."+"vipValue", data.OmpAdminDistanceIpv6.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"timers.advertisement-interval."+"vipObjectType", "object")

	if !data.AdvertisementIntervalVariable.IsNull() {
		body, _ = sjson.Set(body, path+"timers.advertisement-interval."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"timers.advertisement-interval."+"vipVariableName", data.AdvertisementIntervalVariable.ValueString())
	} else if data.AdvertisementInterval.IsNull() {
		body, _ = sjson.Set(body, path+"timers.advertisement-interval."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"timers.advertisement-interval."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"timers.advertisement-interval."+"vipValue", data.AdvertisementInterval.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"timers.graceful-restart-timer."+"vipObjectType", "object")

	if !data.GracefulRestartTimerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"timers.graceful-restart-timer."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"timers.graceful-restart-timer."+"vipVariableName", data.GracefulRestartTimerVariable.ValueString())
	} else if data.GracefulRestartTimer.IsNull() {
		body, _ = sjson.Set(body, path+"timers.graceful-restart-timer."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"timers.graceful-restart-timer."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"timers.graceful-restart-timer."+"vipValue", data.GracefulRestartTimer.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"timers.eor-timer."+"vipObjectType", "object")

	if !data.EorTimerVariable.IsNull() {
		body, _ = sjson.Set(body, path+"timers.eor-timer."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"timers.eor-timer."+"vipVariableName", data.EorTimerVariable.ValueString())
	} else if data.EorTimer.IsNull() {
		body, _ = sjson.Set(body, path+"timers.eor-timer."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"timers.eor-timer."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"timers.eor-timer."+"vipValue", data.EorTimer.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"timers.holdtime."+"vipObjectType", "object")

	if !data.HoldtimeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"timers.holdtime."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"timers.holdtime."+"vipVariableName", data.HoldtimeVariable.ValueString())
	} else if data.Holdtime.IsNull() {
		body, _ = sjson.Set(body, path+"timers.holdtime."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"timers.holdtime."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"timers.holdtime."+"vipValue", data.Holdtime.ValueInt64())
	}
	body, _ = sjson.Set(body, path+"ignore-region-path-length."+"vipObjectType", "object")

	if !data.IgnoreRegionPathLengthVariable.IsNull() {
		body, _ = sjson.Set(body, path+"ignore-region-path-length."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"ignore-region-path-length."+"vipVariableName", data.IgnoreRegionPathLengthVariable.ValueString())
	} else if data.IgnoreRegionPathLength.IsNull() {
		body, _ = sjson.Set(body, path+"ignore-region-path-length."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"ignore-region-path-length."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"ignore-region-path-length."+"vipValue", strconv.FormatBool(data.IgnoreRegionPathLength.ValueBool()))
	}
	body, _ = sjson.Set(body, path+"transport-gateway."+"vipObjectType", "object")

	if !data.TransportGatewayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipVariableName", data.TransportGatewayVariable.ValueString())
	} else if data.TransportGateway.IsNull() {
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipValue", data.TransportGateway.ValueString())
	}
	body, _ = sjson.Set(body, path+"advertise."+"vipObjectType", "tree")
	if len(data.AdvertiseIpv4Routes) > 0 {
		body, _ = sjson.Set(body, path+"advertise."+"vipType", "constant")
	} else {
		body, _ = sjson.Set(body, path+"advertise."+"vipType", "ignore")
	}
	body, _ = sjson.Set(body, path+"advertise."+"vipPrimaryKey", []string{"protocol"})
	body, _ = sjson.Set(body, path+"advertise."+"vipValue", []interface{}{})
	for _, item := range data.AdvertiseIpv4Routes {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
		if item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		itemBody, _ = sjson.Set(itemBody, "route."+"vipObjectType", "object")

		if !item.AdvertiseExternalOspfVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "route."+"vipVariableName", item.AdvertiseExternalOspfVariable.ValueString())
		} else if item.AdvertiseExternalOspf.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "route."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "route."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "route."+"vipValue", item.AdvertiseExternalOspf.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"advertise."+"vipValue.-1", itemBody)
	}
	body, _ = sjson.Set(body, path+"ipv6-advertise."+"vipObjectType", "tree")
	if len(data.AdvertiseIpv6Routes) > 0 {
		body, _ = sjson.Set(body, path+"ipv6-advertise."+"vipType", "constant")
	} else {
		body, _ = sjson.Set(body, path+"ipv6-advertise."+"vipType", "ignore")
	}
	body, _ = sjson.Set(body, path+"ipv6-advertise."+"vipPrimaryKey", []string{"protocol"})
	body, _ = sjson.Set(body, path+"ipv6-advertise."+"vipValue", []interface{}{})
	for _, item := range data.AdvertiseIpv6Routes {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "protocol."+"vipObjectType", "object")
		if item.Protocol.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "protocol."+"vipValue", item.Protocol.ValueString())
		}
		body, _ = sjson.SetRaw(body, path+"ipv6-advertise."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *CiscoOMP) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get(path + "graceful-restart.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GracefulRestart = types.BoolNull()

			v := res.Get(path + "graceful-restart.vipVariableName")
			data.GracefulRestartVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.GracefulRestart = types.BoolNull()
			data.GracefulRestartVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "graceful-restart.vipValue")
			data.GracefulRestart = types.BoolValue(v.Bool())
			data.GracefulRestartVariable = types.StringNull()
		}
	} else {
		data.GracefulRestart = types.BoolNull()
		data.GracefulRestartVariable = types.StringNull()
	}
	if value := res.Get(path + "overlay-as.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OverlayAs = types.Int64Null()

			v := res.Get(path + "overlay-as.vipVariableName")
			data.OverlayAsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OverlayAs = types.Int64Null()
			data.OverlayAsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "overlay-as.vipValue")
			data.OverlayAs = types.Int64Value(v.Int())
			data.OverlayAsVariable = types.StringNull()
		}
	} else {
		data.OverlayAs = types.Int64Null()
		data.OverlayAsVariable = types.StringNull()
	}
	if value := res.Get(path + "send-path-limit.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SendPathLimit = types.Int64Null()

			v := res.Get(path + "send-path-limit.vipVariableName")
			data.SendPathLimitVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SendPathLimit = types.Int64Null()
			data.SendPathLimitVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "send-path-limit.vipValue")
			data.SendPathLimit = types.Int64Value(v.Int())
			data.SendPathLimitVariable = types.StringNull()
		}
	} else {
		data.SendPathLimit = types.Int64Null()
		data.SendPathLimitVariable = types.StringNull()
	}
	if value := res.Get(path + "ecmp-limit.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EcmpLimit = types.Int64Null()

			v := res.Get(path + "ecmp-limit.vipVariableName")
			data.EcmpLimitVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EcmpLimit = types.Int64Null()
			data.EcmpLimitVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ecmp-limit.vipValue")
			data.EcmpLimit = types.Int64Value(v.Int())
			data.EcmpLimitVariable = types.StringNull()
		}
	} else {
		data.EcmpLimit = types.Int64Null()
		data.EcmpLimitVariable = types.StringNull()
	}
	if value := res.Get(path + "shutdown.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Shutdown = types.BoolNull()

			v := res.Get(path + "shutdown.vipVariableName")
			data.ShutdownVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Shutdown = types.BoolNull()
			data.ShutdownVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "shutdown.vipValue")
			data.Shutdown = types.BoolValue(v.Bool())
			data.ShutdownVariable = types.StringNull()
		}
	} else {
		data.Shutdown = types.BoolNull()
		data.ShutdownVariable = types.StringNull()
	}
	if value := res.Get(path + "omp-admin-distance-ipv4.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OmpAdminDistanceIpv4 = types.Int64Null()

			v := res.Get(path + "omp-admin-distance-ipv4.vipVariableName")
			data.OmpAdminDistanceIpv4Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OmpAdminDistanceIpv4 = types.Int64Null()
			data.OmpAdminDistanceIpv4Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "omp-admin-distance-ipv4.vipValue")
			data.OmpAdminDistanceIpv4 = types.Int64Value(v.Int())
			data.OmpAdminDistanceIpv4Variable = types.StringNull()
		}
	} else {
		data.OmpAdminDistanceIpv4 = types.Int64Null()
		data.OmpAdminDistanceIpv4Variable = types.StringNull()
	}
	if value := res.Get(path + "omp-admin-distance-ipv6.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OmpAdminDistanceIpv6 = types.Int64Null()

			v := res.Get(path + "omp-admin-distance-ipv6.vipVariableName")
			data.OmpAdminDistanceIpv6Variable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OmpAdminDistanceIpv6 = types.Int64Null()
			data.OmpAdminDistanceIpv6Variable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "omp-admin-distance-ipv6.vipValue")
			data.OmpAdminDistanceIpv6 = types.Int64Value(v.Int())
			data.OmpAdminDistanceIpv6Variable = types.StringNull()
		}
	} else {
		data.OmpAdminDistanceIpv6 = types.Int64Null()
		data.OmpAdminDistanceIpv6Variable = types.StringNull()
	}
	if value := res.Get(path + "timers.advertisement-interval.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AdvertisementInterval = types.Int64Null()

			v := res.Get(path + "timers.advertisement-interval.vipVariableName")
			data.AdvertisementIntervalVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AdvertisementInterval = types.Int64Null()
			data.AdvertisementIntervalVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "timers.advertisement-interval.vipValue")
			data.AdvertisementInterval = types.Int64Value(v.Int())
			data.AdvertisementIntervalVariable = types.StringNull()
		}
	} else {
		data.AdvertisementInterval = types.Int64Null()
		data.AdvertisementIntervalVariable = types.StringNull()
	}
	if value := res.Get(path + "timers.graceful-restart-timer.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GracefulRestartTimer = types.Int64Null()

			v := res.Get(path + "timers.graceful-restart-timer.vipVariableName")
			data.GracefulRestartTimerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.GracefulRestartTimer = types.Int64Null()
			data.GracefulRestartTimerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "timers.graceful-restart-timer.vipValue")
			data.GracefulRestartTimer = types.Int64Value(v.Int())
			data.GracefulRestartTimerVariable = types.StringNull()
		}
	} else {
		data.GracefulRestartTimer = types.Int64Null()
		data.GracefulRestartTimerVariable = types.StringNull()
	}
	if value := res.Get(path + "timers.eor-timer.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EorTimer = types.Int64Null()

			v := res.Get(path + "timers.eor-timer.vipVariableName")
			data.EorTimerVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.EorTimer = types.Int64Null()
			data.EorTimerVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "timers.eor-timer.vipValue")
			data.EorTimer = types.Int64Value(v.Int())
			data.EorTimerVariable = types.StringNull()
		}
	} else {
		data.EorTimer = types.Int64Null()
		data.EorTimerVariable = types.StringNull()
	}
	if value := res.Get(path + "timers.holdtime.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Holdtime = types.Int64Null()

			v := res.Get(path + "timers.holdtime.vipVariableName")
			data.HoldtimeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Holdtime = types.Int64Null()
			data.HoldtimeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "timers.holdtime.vipValue")
			data.Holdtime = types.Int64Value(v.Int())
			data.HoldtimeVariable = types.StringNull()
		}
	} else {
		data.Holdtime = types.Int64Null()
		data.HoldtimeVariable = types.StringNull()
	}
	if value := res.Get(path + "ignore-region-path-length.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IgnoreRegionPathLength = types.BoolNull()

			v := res.Get(path + "ignore-region-path-length.vipVariableName")
			data.IgnoreRegionPathLengthVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IgnoreRegionPathLength = types.BoolNull()
			data.IgnoreRegionPathLengthVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "ignore-region-path-length.vipValue")
			data.IgnoreRegionPathLength = types.BoolValue(v.Bool())
			data.IgnoreRegionPathLengthVariable = types.StringNull()
		}
	} else {
		data.IgnoreRegionPathLength = types.BoolNull()
		data.IgnoreRegionPathLengthVariable = types.StringNull()
	}
	if value := res.Get(path + "transport-gateway.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TransportGateway = types.StringNull()

			v := res.Get(path + "transport-gateway.vipVariableName")
			data.TransportGatewayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TransportGateway = types.StringNull()
			data.TransportGatewayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "transport-gateway.vipValue")
			data.TransportGateway = types.StringValue(v.String())
			data.TransportGatewayVariable = types.StringNull()
		}
	} else {
		data.TransportGateway = types.StringNull()
		data.TransportGatewayVariable = types.StringNull()
	}
	if value := res.Get(path + "advertise.vipValue"); len(value.Array()) > 0 {
		data.AdvertiseIpv4Routes = make([]CiscoOMPAdvertiseIpv4Routes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOMPAdvertiseIpv4Routes{}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())

				}
			} else {
				item.Protocol = types.StringNull()

			}
			if cValue := v.Get("route.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.AdvertiseExternalOspf = types.StringNull()

					cv := v.Get("route.vipVariableName")
					item.AdvertiseExternalOspfVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.AdvertiseExternalOspf = types.StringNull()
					item.AdvertiseExternalOspfVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("route.vipValue")
					item.AdvertiseExternalOspf = types.StringValue(cv.String())
					item.AdvertiseExternalOspfVariable = types.StringNull()
				}
			} else {
				item.AdvertiseExternalOspf = types.StringNull()
				item.AdvertiseExternalOspfVariable = types.StringNull()
			}
			data.AdvertiseIpv4Routes = append(data.AdvertiseIpv4Routes, item)
			return true
		})
	}
	if value := res.Get(path + "ipv6-advertise.vipValue"); len(value.Array()) > 0 {
		data.AdvertiseIpv6Routes = make([]CiscoOMPAdvertiseIpv6Routes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoOMPAdvertiseIpv6Routes{}
			if cValue := v.Get("protocol.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Protocol = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.Protocol = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("protocol.vipValue")
					item.Protocol = types.StringValue(cv.String())

				}
			} else {
				item.Protocol = types.StringNull()

			}
			data.AdvertiseIpv6Routes = append(data.AdvertiseIpv6Routes, item)
			return true
		})
	}
}
