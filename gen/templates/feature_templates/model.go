//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
	TemplateType types.String `tfsdk:"template_type"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	DeviceTypes types.List `tfsdk:"device_types"`
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "ListString"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if eq .Variable true}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if eq .Variable true}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "ListString"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if eq .Variable true}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if eq .Variable true}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
{{ range .Attributes}}
{{- if eq .Type "List"}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if eq .Type "ListString"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- if eq .Variable true}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if eq .Variable true}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

func (data {{camelCase .Name}}) getModel() string {
	return "{{.Model}}"
}

func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "{{.Model}}")

	path := "templateDefinition."
	{{- range .Attributes}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
	{{if eq .Variable true}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.ValueString())
	} else
	{{- end}} if data.{{toGoName .TfName}}.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", data.{{toGoName .TfName}}.Value{{.Type}}())
	}
	{{- else if eq .Type "Bool"}}
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
	{{if eq .Variable true}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.ValueString())
	} else
	{{- end}} if data.{{toGoName .TfName}}.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(data.{{toGoName .TfName}}.ValueBool()))
	}
	{{- else if eq .Type "ListString"}}
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
	{{if eq .Variable true}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", data.{{toGoName .TfName}}Variable.ValueString())
	} else
	{{- end}} if data.{{toGoName .TfName}}.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		var values []string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
	}
	{{- else if eq .Type "List"}}
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
	} else {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
	}
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
	for _, item := range data.{{toGoName .TfName}} {
		itemBody := ""
		{{- range .Attributes}}
		{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
		itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		{{if eq .Variable true}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.ValueString())
		} else
		{{- end}} if item.{{toGoName .TfName}}.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", item.{{toGoName .TfName}}.Value{{.Type}}())
		}
		{{- else if eq .Type "Bool"}}
		itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		{{if eq .Variable true}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.ValueString())
		} else
		{{- end}} if item.{{toGoName .TfName}}.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(item.{{toGoName .TfName}}.ValueBool()))
		}
		{{- else if eq .Type "ListString"}}
		itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		{{if eq .Variable true}}
		if !item.{{toGoName .TfName}}Variable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", item.{{toGoName .TfName}}Variable.ValueString())
		} else
		{{- end}} if item.{{toGoName .TfName}}.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
			var values []string
			item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
		}
		{{- else if eq .Type "List"}}
		itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
		if len(item.{{toGoName .TfName}}) > 0 {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
		} else {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
		}
		itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipPrimaryKey", []string{ {{range .Keys}}"{{.}}",{{end}} })
		itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", []interface{}{})
		for _, childItem := range item.{{toGoName .TfName}} {
			itemChildBody := ""
			{{- range .Attributes}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
			itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			{{if eq .Variable true}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.ValueString())
			} else
			{{- end}} if childItem.{{toGoName .TfName}}.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", childItem.{{toGoName .TfName}}.Value{{.Type}}())
			}
			{{- else if eq .Type "Bool"}}
			itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			{{if eq .Variable true}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.ValueString())
			} else
			{{- end}} if childItem.{{toGoName .TfName}}.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", strconv.FormatBool(childItem.{{toGoName .TfName}}.ValueBool()))
			}
			{{- else if eq .Type "ListString"}}
			itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipObjectType", "{{.ObjectType}}")
			{{if eq .Variable true}}
			if !childItem.{{toGoName .TfName}}Variable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipVariableName", childItem.{{toGoName .TfName}}Variable.ValueString())
			} else
			{{- end}} if childItem.{{toGoName .TfName}}.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "ignore")
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipType", "constant")
				var values []string
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue", values)
			}
			{{- end}}
			{{- end}}
			itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue.-1", itemChildBody)
		}
		{{- end}}
		{{- end}}
		body, _ = sjson.SetRaw(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}."+"vipValue.-1", itemBody)
	}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
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
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if eq .Type "String"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.StringNull()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.StringNull()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.StringValue(v.String())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "Int64"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.Int64Null()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.Int64Null()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.Int64Value(v.Int())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "Float64"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.Float64Null()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.Float64Null()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.Float64Value(v.Float())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.Float64Null()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.BoolNull()
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.BoolNull()
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = types.BoolValue(v.Bool())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.BoolNull()
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "ListString"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.{{toGoName .TfName}} = types.ListNull(types.StringType)
			{{if .Variable}}
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
			data.{{toGoName .TfName}}Variable = types.StringValue(v.String())
			{{end}}
		} else if value.String() == "ignore" {
			data.{{toGoName .TfName}} = types.ListNull(types.StringType)
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		} else if value.String() == "constant" {
			v := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
			data.{{toGoName .TfName}} = helpers.GetStringList(v.Array())
			{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		}
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
		{{if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	}
	{{- else if eq .Type "List"}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue"); len(value.Array()) > 0 {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- if eq .Type "String"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.StringNull()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.StringNull()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.StringValue(cv.String())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.StringNull()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "Int64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.Int64Null()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.Int64Null()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.Int64Value(cv.Int())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.Int64Null()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "Float64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.Float64Null()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.Float64Null()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.Float64Value(cv.Float())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.Float64Null()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.BoolNull()
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.BoolNull()
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = types.BoolValue(cv.Bool())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.BoolNull()
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "ListString"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.{{toGoName .TfName}} = types.ListNull(types.StringType)
					{{if .Variable}}
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
					item.{{toGoName .TfName}}Variable = types.StringValue(cv.String())
					{{end}}
				} else if cValue.String() == "ignore" {
					item.{{toGoName .TfName}} = types.ListNull(types.StringType)
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				} else if cValue.String() == "constant" {
					cv := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
					item.{{toGoName .TfName}} = helpers.GetStringList(cv.Array())
					{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				}
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.StringType)
				{{if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			}
			{{- else if eq .Type "List"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue"); len(cValue.Array()) > 0 {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if eq .Type "String"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.StringNull()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.StringNull()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.StringValue(ccv.String())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.StringNull()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "Int64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.Int64Null()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.Int64Null()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.Int64Value(ccv.Int())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.Int64Null()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "Float64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.Float64Null()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.Float64Null()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.Float64Value(ccv.Float())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.Float64Null()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.BoolNull()
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.BoolNull()
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = types.BoolValue(ccv.Bool())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.BoolNull()
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- else if eq .Type "ListString"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipType"); len(ccValue.Array()) > 0 {
						if ccValue.String() == "variableName" {
							cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
							{{if .Variable}}
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipVariableName")
							cItem.{{toGoName .TfName}}Variable = types.StringValue(ccv.String())
							{{end}}
						} else if ccValue.String() == "ignore" {
							cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.vipValue")
							cItem.{{toGoName .TfName}} = helpers.GetStringList(ccv.Array())
							{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
						}
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
						{{if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
}
