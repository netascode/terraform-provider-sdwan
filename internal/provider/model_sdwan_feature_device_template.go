package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type FeatureDeviceTemplate struct {
	Id               types.String                           `tfsdk:"id"`
	Name             types.String                           `tfsdk:"name"`
	Description      types.String                           `tfsdk:"description"`
	DeviceType       types.String                           `tfsdk:"device_type"`
	DeviceRole       types.String                           `tfsdk:"device_role"`
	PolicyId         types.String                           `tfsdk:"policy_id"`
	SecurityPolicyId types.String                           `tfsdk:"security_policy_id"`
	GeneralTemplates []FeatureDeviceTemplateGeneralTemplate `tfsdk:"general_templates"`
}

type FeatureDeviceTemplateGeneralTemplate struct {
	Id           types.String                       `tfsdk:"id"`
	Type         types.String                       `tfsdk:"type"`
	SubTemplates []FeatureDeviceTemplateSubTemplate `tfsdk:"sub_templates"`
}

type FeatureDeviceTemplateSubTemplate struct {
	Id           types.String                          `tfsdk:"id"`
	Type         types.String                          `tfsdk:"type"`
	SubTemplates []FeatureDeviceTemplateSubSubTemplate `tfsdk:"sub_templates"`
}

type FeatureDeviceTemplateSubSubTemplate struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (data FeatureDeviceTemplate) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "deviceType", data.DeviceType.ValueString())
	body, _ = sjson.Set(body, "deviceRole", data.DeviceRole.ValueString())
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "configType", "template")
	body, _ = sjson.Set(body, "draftMode", false)
	body, _ = sjson.Set(body, "featureTemplateUidRange", []interface{}{})
	body, _ = sjson.Set(body, "policyId", data.PolicyId.ValueString())
	body, _ = sjson.Set(body, "securityPolicyId", data.SecurityPolicyId.ValueString())
	for _, item := range data.GeneralTemplates {
		itemBody := ""
		itemBody, _ = sjson.Set(itemBody, "templateId", item.Id.ValueString())
		itemBody, _ = sjson.Set(itemBody, "templateType", item.Type.ValueString())
		for _, sub := range item.SubTemplates {
			subBody := ""
			subBody, _ = sjson.Set(subBody, "templateId", sub.Id.ValueString())
			subBody, _ = sjson.Set(subBody, "templateType", sub.Type.ValueString())
			for _, subSub := range sub.SubTemplates {
				subSubBody := ""
				subSubBody, _ = sjson.Set(subSubBody, "templateId", subSub.Id.ValueString())
				subSubBody, _ = sjson.Set(subSubBody, "templateType", subSub.Type.ValueString())
				subBody, _ = sjson.SetRaw(subBody, "subTemplates.-1", subSubBody)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "subTemplates.-1", subBody)
		}
		body, _ = sjson.SetRaw(body, "generalTemplates.-1", itemBody)
	}
	return body
}

func (data *FeatureDeviceTemplate) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateDescription"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceType = types.StringValue(value.String())
	} else {
		data.DeviceType = types.StringNull()
	}
	if value := res.Get("deviceRole"); value.Exists() && value.String() != "" {
		data.DeviceRole = types.StringValue(value.String())
	} else {
		data.DeviceRole = types.StringNull()
	}
	if value := res.Get("policyId"); value.Exists() && value.String() != "" {
		data.PolicyId = types.StringValue(value.String())
	} else {
		data.PolicyId = types.StringNull()
	}
	if value := res.Get("securityPolicyId"); value.Exists() && value.String() != "" {
		data.SecurityPolicyId = types.StringValue(value.String())
	} else {
		data.SecurityPolicyId = types.StringNull()
	}
	if value := res.Get("generalTemplates"); value.Exists() {
		data.GeneralTemplates = make([]FeatureDeviceTemplateGeneralTemplate, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := FeatureDeviceTemplateGeneralTemplate{}
			if cValue := v.Get("templateId"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("templateType"); cValue.Exists() {
				item.Type = types.StringValue(cValue.String())
			} else {
				item.Type = types.StringNull()
			}
			if sValue := v.Get("subTemplates"); sValue.Exists() {
				item.SubTemplates = make([]FeatureDeviceTemplateSubTemplate, 0)
				sValue.ForEach(func(k, v gjson.Result) bool {
					sItem := FeatureDeviceTemplateSubTemplate{}
					if csValue := v.Get("templateId"); csValue.Exists() {
						sItem.Id = types.StringValue(csValue.String())
					} else {
						sItem.Id = types.StringNull()
					}
					if csValue := v.Get("templateType"); csValue.Exists() {
						sItem.Type = types.StringValue(csValue.String())
					} else {
						sItem.Type = types.StringNull()
					}
					if ssValue := v.Get("subTemplates"); ssValue.Exists() {
						sItem.SubTemplates = make([]FeatureDeviceTemplateSubSubTemplate, 0)
						ssValue.ForEach(func(k, v gjson.Result) bool {
							ssItem := FeatureDeviceTemplateSubSubTemplate{}
							if cssValue := v.Get("templateId"); cssValue.Exists() {
								ssItem.Id = types.StringValue(cssValue.String())
							} else {
								ssItem.Id = types.StringNull()
							}
							if cssValue := v.Get("templateType"); cssValue.Exists() {
								ssItem.Type = types.StringValue(cssValue.String())
							} else {
								ssItem.Type = types.StringNull()
							}
							sItem.SubTemplates = append(sItem.SubTemplates, ssItem)
							return true
						})
					}
					item.SubTemplates = append(item.SubTemplates, sItem)
					return true
				})
			}
			data.GeneralTemplates = append(data.GeneralTemplates, item)
			return true
		})
	}
}
