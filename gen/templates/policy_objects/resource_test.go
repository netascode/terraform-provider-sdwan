//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwan{{camelCase .Name}}PolicyObject(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwan{{camelCase .Name}}PolicyObjectConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (ne .WriteOnly true) (ne .ExcludeTest true)}}
					{{- if eq .Type "List"}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (ne .WriteOnly true) (ne .ExcludeTest true)}}
					{{- if eq .Type "List"}}
					{{- $clist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (ne .WriteOnly true) (ne .ExcludeTest true)}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_policy_object.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_policy_object.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_policy_object.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

func testAccSdwan{{camelCase .Name}}PolicyObjectConfig_all() string {
	return `
	resource "sdwan_{{snakeCase $name}}_policy_object" "test" {
		name = "TF_TEST_ALL"
		{{- range  .Attributes}}
		{{- if not .ExcludeTest}}
		{{- if eq .Type "List"}}
		{{.TfName}} = [{
			{{- range  .Attributes}}
			{{- if ne .ExcludeTest true}}
			{{- if eq .Type "List"}}
			{{.TfName}} = [{
				{{- range  .Attributes}}
				{{- if ne .ExcludeTest true}}
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
	}
	`
}
