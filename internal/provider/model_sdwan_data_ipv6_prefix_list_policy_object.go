// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type DataIPv6PrefixList struct {
	Id      types.String                `tfsdk:"id"`
	Name    types.String                `tfsdk:"name"`
	Entries []DataIPv6PrefixListEntries `tfsdk:"entries"`
}

type DataIPv6PrefixListEntries struct {
	Prefix types.String `tfsdk:"prefix"`
}

func (data DataIPv6PrefixList) getType() string {
	return "dataipv6prefix"
}

func (data DataIPv6PrefixList) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "description", "Desc Not Required")
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "type", "dataipv6prefix")
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6Prefix", item.Prefix.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *DataIPv6PrefixList) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]DataIPv6PrefixListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := DataIPv6PrefixListEntries{}
			if cValue := v.Get("ipv6Prefix"); cValue.Exists() {
				item.Prefix = types.StringValue(cValue.String())
			} else {
				item.Prefix = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}
