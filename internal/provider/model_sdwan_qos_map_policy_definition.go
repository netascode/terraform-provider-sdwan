// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type QoSMap struct {
	Id            types.String          `tfsdk:"id"`
	Name          types.String          `tfsdk:"name"`
	Description   types.String          `tfsdk:"description"`
	QosSchedulers []QoSMapQosSchedulers `tfsdk:"qos_schedulers"`
}

type QoSMapQosSchedulers struct {
	BandwidthPercent types.Int64  `tfsdk:"bandwidth_percent"`
	BufferPercent    types.Int64  `tfsdk:"buffer_percent"`
	Burst            types.Int64  `tfsdk:"burst"`
	ClassMapId       types.String `tfsdk:"class_map_id"`
	DropType         types.String `tfsdk:"drop_type"`
	Queue            types.Int64  `tfsdk:"queue"`
	SchedulingType   types.String `tfsdk:"scheduling_type"`
}

func (data QoSMap) getType() string {
	return "qosMap"
}

func (data QoSMap) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	body, _ = sjson.Set(body, "type", "qosMap")
	path := "definition."
	if len(data.QosSchedulers) > 0 {
		body, _ = sjson.Set(body, path+"qosSchedulers", []interface{}{})
		for _, item := range data.QosSchedulers {
			itemBody := ""
			if !item.BandwidthPercent.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bandwidthPercent", fmt.Sprint(item.BandwidthPercent.ValueInt64()))
			}
			if !item.BufferPercent.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bufferPercent", fmt.Sprint(item.BufferPercent.ValueInt64()))
			}
			if !item.Burst.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "burst", fmt.Sprint(item.Burst.ValueInt64()))
			}
			if !item.ClassMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "classMapRef", item.ClassMapId.ValueString())
			}
			if !item.DropType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "drops", item.DropType.ValueString())
			}
			if !item.Queue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "queue", fmt.Sprint(item.Queue.ValueInt64()))
			}
			if !item.SchedulingType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "scheduling", item.SchedulingType.ValueString())
			}
			body, _ = sjson.SetRaw(body, path+"qosSchedulers.-1", itemBody)
		}
	}
	return body
}

func (data *QoSMap) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "definition."
	if value := res.Get(path + "qosSchedulers"); value.Exists() {
		data.QosSchedulers = make([]QoSMapQosSchedulers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := QoSMapQosSchedulers{}
			if cValue := v.Get("bandwidthPercent"); cValue.Exists() {
				item.BandwidthPercent = types.Int64Value(cValue.Int())
			} else {
				item.BandwidthPercent = types.Int64Null()
			}
			if cValue := v.Get("bufferPercent"); cValue.Exists() {
				item.BufferPercent = types.Int64Value(cValue.Int())
			} else {
				item.BufferPercent = types.Int64Null()
			}
			if cValue := v.Get("burst"); cValue.Exists() {
				item.Burst = types.Int64Value(cValue.Int())
			} else {
				item.Burst = types.Int64Null()
			}
			if cValue := v.Get("classMapRef"); cValue.Exists() {
				item.ClassMapId = types.StringValue(cValue.String())
			} else {
				item.ClassMapId = types.StringNull()
			}
			if cValue := v.Get("drops"); cValue.Exists() {
				item.DropType = types.StringValue(cValue.String())
			} else {
				item.DropType = types.StringNull()
			}
			if cValue := v.Get("queue"); cValue.Exists() {
				item.Queue = types.Int64Value(cValue.Int())
			} else {
				item.Queue = types.Int64Null()
			}
			if cValue := v.Get("scheduling"); cValue.Exists() {
				item.SchedulingType = types.StringValue(cValue.String())
			} else {
				item.SchedulingType = types.StringNull()
			}
			data.QosSchedulers = append(data.QosSchedulers, item)
			return true
		})
	}
}
