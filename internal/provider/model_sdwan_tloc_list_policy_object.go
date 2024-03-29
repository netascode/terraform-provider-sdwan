// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type TLOCList struct {
	Id      types.String      `tfsdk:"id"`
	Version types.Int64       `tfsdk:"version"`
	Name    types.String      `tfsdk:"name"`
	Entries []TLOCListEntries `tfsdk:"entries"`
}

type TLOCListEntries struct {
	TlocIp        types.String `tfsdk:"tloc_ip"`
	Color         types.String `tfsdk:"color"`
	Encapsulation types.String `tfsdk:"encapsulation"`
	Preference    types.Int64  `tfsdk:"preference"`
}

func (data TLOCList) getType() string {
	return "tloc"
}

func (data TLOCList) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "description", "Desc Not Required")
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "type", "tloc")
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.TlocIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tloc", item.TlocIp.ValueString())
			}
			if !item.Color.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "color", item.Color.ValueString())
			}
			if !item.Encapsulation.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "encap", item.Encapsulation.ValueString())
			}
			if !item.Preference.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "preference", fmt.Sprint(item.Preference.ValueInt64()))
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *TLOCList) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]TLOCListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TLOCListEntries{}
			if cValue := v.Get("tloc"); cValue.Exists() {
				item.TlocIp = types.StringValue(cValue.String())
			} else {
				item.TlocIp = types.StringNull()
			}
			if cValue := v.Get("color"); cValue.Exists() {
				item.Color = types.StringValue(cValue.String())
			} else {
				item.Color = types.StringNull()
			}
			if cValue := v.Get("encap"); cValue.Exists() {
				item.Encapsulation = types.StringValue(cValue.String())
			} else {
				item.Encapsulation = types.StringNull()
			}
			if cValue := v.Get("preference"); cValue.Exists() {
				item.Preference = types.Int64Value(cValue.Int())
			} else {
				item.Preference = types.Int64Null()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}
