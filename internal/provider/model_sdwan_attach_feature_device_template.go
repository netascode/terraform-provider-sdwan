package provider

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type AttachFeatureDeviceTemplate struct {
	Id      types.String                        `tfsdk:"id"`
	Devices []AttachFeatureDeviceTemplateDevice `tfsdk:"devices"`
}

type AttachFeatureDeviceTemplateDevice struct {
	Id        types.String `tfsdk:"id"`
	Variables types.Map    `tfsdk:"variables"`
}

func (data AttachFeatureDeviceTemplate) getVariables(ctx context.Context, client *sdwan.Client) (map[string]map[string]string, error) {
	deviceVariables := make(map[string]map[string]string)
	for _, item := range data.Devices {
		// Retrieve device variables
		body, _ := sjson.Set("", "templateId", data.Id.ValueString())
		body, _ = sjson.Set(body, "deviceIds", []string{item.Id.ValueString()})
		body, _ = sjson.Set(body, "isEdited", false)
		body, _ = sjson.Set(body, "isMasterEdited", false)
		res, err := client.Post("/template/device/config/input", body)
		if err != nil {
			return nil, err
		}
		// Get variable mappings
		mappings := make(map[string]string)
		if value := res.Get("header.columns"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				if v.Get("editable").Bool() {
					title := v.Get("title").String()
					if strings.Contains(title, "(") {
						matches := regexp.MustCompile(`\((.*?)\)`).FindAllString(title, -1)
						varName := matches[len(matches)-1]
						if len(varName) > 0 {
							varName = varName[1 : len(varName)-1]
						}
						mappings[v.Get("property").String()] = varName
					}
				}
				return true
			})
		}
		// Get current device variables
		variables := make(map[string]string)
		if value := res.Get("data.0"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				variables[k.String()] = v.String()
				return true
			})
		}
		// Resolve variable names and insert template variable values
		var templateVariables map[string]string
		item.Variables.ElementsAs(ctx, &templateVariables, false)
		for k, _ := range variables {
			templateVariableName, ok := mappings[k]
			if ok {
				variables[k] = templateVariables[templateVariableName]
			}
		}
		deviceVariables[item.Id.ValueString()] = variables
	}

	return deviceVariables, nil
}

func (data AttachFeatureDeviceTemplate) getAttachedDevices(ctx context.Context, client *sdwan.Client) (map[string]string, error) {
	res, err := client.Get("/template/device/config/attached/" + data.Id.ValueString())
	if err != nil {
		return nil, err
	}
	devices := make(map[string]string)
	for _, device := range res.Get("data").Array() {
		devices[device.Get("uuid").String()] = device.Get("deviceIP").String()

	}
	return devices, nil
}

func (data AttachFeatureDeviceTemplate) detachDevices(ctx context.Context, client *sdwan.Client) (gjson.Result, error) {
	devices, err := data.getAttachedDevices(ctx, client)
	if err != nil {
		return gjson.Result{}, err
	}

	body, _ := sjson.Set("", "deviceType", "vedge")
	index := 0
	for k, v := range devices {
		body, _ = sjson.Set(body, "devices."+strconv.Itoa(index)+".deviceId", k)
		body, _ = sjson.Set(body, "devices."+strconv.Itoa(index)+".deviceIP", v)
		index++
	}

	res, err := client.Post("/template/config/device/mode/cli", body)
	if err != nil {
		return gjson.Result{}, err
	}
	return res, nil
}

func (data AttachFeatureDeviceTemplate) toBody(ctx context.Context, client *sdwan.Client) (string, error) {
	deviceVariables, err := data.getVariables(ctx, client)
	if err != nil {
		return "", err
	}
	body, _ := sjson.Set("", "deviceTemplateList.0.templateId", data.Id.ValueString())
	body, _ = sjson.Set(body, "deviceTemplateList.0.isEdited", false)
	body, _ = sjson.Set(body, "deviceTemplateList.0.isMasterEdited", false)
	body, _ = sjson.Set(body, "deviceTemplateList.0.isDraftDisabled", false)

	for _, v := range deviceVariables {
		itemBody := ""
		for name, value := range v {
			itemBody, _ = sjson.Set(itemBody, name, value)
		}
		body, _ = sjson.SetRaw(body, "deviceTemplateList.0.device.-1", itemBody)
	}

	return body, nil
}

func (data *AttachFeatureDeviceTemplate) readVariables(ctx context.Context, client *sdwan.Client) error {
	for i := range data.Devices {
		// Retrieve device variables
		body, _ := sjson.Set("", "templateId", data.Id.ValueString())
		body, _ = sjson.Set(body, "deviceIds", []string{data.Devices[i].Id.ValueString()})
		body, _ = sjson.Set(body, "isEdited", false)
		body, _ = sjson.Set(body, "isMasterEdited", false)
		res, err := client.Post("/template/device/config/input", body)
		if err != nil {
			return err
		}
		// Get variable mappings
		mappings := make(map[string]string)
		if value := res.Get("header.columns"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				if v.Get("editable").Bool() {
					title := v.Get("title").String()
					if strings.Contains(title, "(") {
						matches := regexp.MustCompile(`\((.*?)\)`).FindAllString(title, -1)
						varName := matches[len(matches)-1]
						if len(varName) > 0 {
							varName = varName[1 : len(varName)-1]
						}
						mappings[v.Get("property").String()] = varName
					}
				}
				return true
			})
		}
		// Get current device variables
		variables := make(map[string]string)
		if value := res.Get("data.0"); value.Exists() {
			value.ForEach(func(k, v gjson.Result) bool {
				variables[k.String()] = v.String()
				return true
			})
		}
		// Resolve variable names and insert template variable values
		templateVariables := make(map[string]attr.Value)
		for k, _ := range variables {
			templateVariableName, ok := mappings[k]
			if ok {
				templateVariables[templateVariableName] = types.StringValue(variables[k])
			}
		}
		data.Devices[i].Variables = types.MapValueMust(types.StringType, templateVariables)
	}

	return nil
}
