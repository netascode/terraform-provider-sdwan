//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwan{{camelCase .Name}}FeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwan{{camelCase .Name}}FeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (.Mandatory) (ne .Type "ListString")}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_feature_template.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
				),
			},
			{
				Config: testAccSdwan{{camelCase .Name}}FeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest)}}
					{{- if eq .Type "List"}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest)}}
					{{- if eq .Type "List"}}
					{{- $clist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest)}}
					{{- if eq .Type "List"}}
					{{- $cclist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (ne .Type "ListString")}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_feature_template.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else if ne .Type "ListString"}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_feature_template.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else if ne .Type "ListString"}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_feature_template.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else if ne .Type "ListString"}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_feature_template.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

func testAccSdwan{{camelCase .Name}}FeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_{{snakeCase $name}}_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		{{- range .Attributes}}
		{{- if and (.Mandatory) (ne .Type "List")}}
		{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
		{{- end}}
		{{- end}}
	}
	`
}

func testAccSdwan{{camelCase .Name}}FeatureTemplateConfig_all() string {
	return `
	resource "sdwan_{{snakeCase $name}}_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		{{- range  .Attributes}}
		{{- if not .ExcludeTest}}
		{{- if eq .Type "List"}}
		{{.TfName}} = [{
			{{- range  .Attributes}}
			{{- if not .ExcludeTest}}
			{{- if eq .Type "List"}}
			{{.TfName}} = [{
				{{- range  .Attributes}}
				{{- if not .ExcludeTest}}
				{{- if eq .Type "List"}}
				{{.TfName}} = [{
					{{- range  .Attributes}}
					{{- if not .ExcludeTest}}
					{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
					{{- end}}
					{{- end}}
				}]
				{{- else}}
				{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
				{{- end}}
				{{- end}}
				{{- end}}
			}]
			{{- else}}
			{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
			{{- end}}
			{{- end}}
			{{- end}}
		}]
		{{- else}}
		{{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	`
}
