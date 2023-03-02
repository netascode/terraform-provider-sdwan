// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ExpandedCommunityList struct {
	Id      types.String                   `tfsdk:"id"`
	Name    types.String                   `tfsdk:"name"`
	Entries []ExpandedCommunityListEntries `tfsdk:"entries"`
}

type ExpandedCommunityListEntries struct {
	Community types.String `tfsdk:"community"`
}

func (data ExpandedCommunityList) getType() string {
	return "expandedcommunity"
}

func (data ExpandedCommunityList) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "description", "Desc Not Required")
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "type", "expandedcommunity")
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			itemBody, _ = sjson.Set(itemBody, "community", fmt.Sprint(item.Community.ValueString()))
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *ExpandedCommunityList) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]ExpandedCommunityListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ExpandedCommunityListEntries{}
			if cValue := v.Get("community"); cValue.Exists() {
				item.Community = types.StringValue(cValue.String())
			} else {
				item.Community = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}