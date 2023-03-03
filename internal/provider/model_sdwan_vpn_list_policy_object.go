// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type VPNList struct {
	Id      types.String     `tfsdk:"id"`
	Name    types.String     `tfsdk:"name"`
	Entries []VPNListEntries `tfsdk:"entries"`
}

type VPNListEntries struct {
	VpnId types.String `tfsdk:"vpn_id"`
}

func (data VPNList) getType() string {
	return "vpn"
}

func (data VPNList) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "description", "Desc Not Required")
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "type", "vpn")
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.VpnId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "vpn", item.VpnId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

func (data *VPNList) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]VPNListEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := VPNListEntries{}
			if cValue := v.Get("vpn"); cValue.Exists() {
				item.VpnId = types.StringValue(cValue.String())
			} else {
				item.VpnId = types.StringNull()
			}
			data.Entries = append(data.Entries, item)
			return true
		})
	}
}
