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
	Version types.Int64 `tfsdk:"version"`
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "ListString"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if eq .Type "List"}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "ListString"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if eq .Type "List"}}
{{ range .Attributes}}
{{- if not .Value}}
{{- if eq .Type "List"}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if eq .Type "ListString"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

func (data {{camelCase .Name}}) getType() string {
	return "{{.Type}}"
}

func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body, _ := sjson.Set("", "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	body, _ = sjson.Set(body, "type", "{{.Type}}")
	path := "{{ if .RootElement}}{{.RootElement}}.{{end}}"
	{{- range .Attributes}}
	{{- if .Value}}
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "{{.Value}}")
	{{- else if not .TfOnly}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
	if !data.{{toGoName .TfName}}.IsNull() {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
	}
	{{- else if eq .Type "ListString"}}
	if !data.{{toGoName .TfName}}.IsNull() {
		var values []string
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if eq .Type "List"}}
	if len(data.{{toGoName .TfName}}) > 0 {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "{{.Value}}")
			{{- else if not .TfOnly}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
			if !item.{{toGoName .TfName}}.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
			}
			{{- else if eq .Type "ListString"}}
			if !item.{{toGoName .TfName}}.IsNull() {
				var values []string
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if eq .Type "List"}}
			if len(item.{{toGoName .TfName}}) > 0 {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "{{.Value}}")
					{{- else if not .TfOnly}}
					{{- if eq .Type "ListString"}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						var values []string
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
					}
					{{- else}}
					if !childItem.{{toGoName .TfName}}.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return body
}

func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	path := "{{ if .RootElement}}{{.RootElement}}.{{end}}"
	{{- range .Attributes}}
	{{- if and (not .TfOnly) (not .Value)}}
	{{- $cname := toGoName .TfName}}
	{{- if eq .Type "String"}}
	if value := res.Get(path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.StringValue(value.String())
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
	}
	{{- else if eq .Type "Int64"}}
	if value := res.Get(path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
	}
	{{- else if eq .Type "Float64"}}
	if value := res.Get(path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = types.Float64Value(value.Float())
	} else {
		data.{{toGoName .TfName}} = types.Float64Null()
	}
	{{- else if eq .Type "ListString"}}
	if value := res.Get(path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = helpers.GetStringList(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.ListNull(types.StringType)
	}
	{{- else if eq .Type "List"}}
	if value := res.Get(path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- if and (not .TfOnly) (not .Value)}}
			{{- if eq .Type "String"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.StringValue(cValue.String())
			} else {
				item.{{toGoName .TfName}} = types.StringNull()
			}
			{{- else if eq .Type "Int64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.Int64Value(cValue.Int())
			} else {
				item.{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Float64"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = types.Float64Value(cValue.Float())
			} else {
				item.{{toGoName .TfName}} = types.Float64Null()
			}
			{{- else if eq .Type "ListString"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = helpers.GetStringList(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if eq .Type "List"}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (not .TfOnly) (not .Value)}}
					{{- if eq .Type "String"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.StringValue(ccValue.String())
					} else {
						cItem.{{toGoName .TfName}} = types.StringNull()
					}
					{{- else if eq .Type "Int64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.Int64Value(ccValue.Int())
					} else {
						cItem.{{toGoName .TfName}} = types.Int64Null()
					}
					{{- else if eq .Type "Float64"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = types.Float64Value(ccValue.Float())
					} else {
						cItem.{{toGoName .TfName}} = types.Float64Null()
					}
					{{- else if eq .Type "ListString"}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = helpers.GetStringList(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.ListNull(types.StringType)
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
	{{- end}}
}
