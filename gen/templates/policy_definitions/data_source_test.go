//go:build ignore
{{if .ExcludeTest}}//go:build testAll{{end}}
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwan{{camelCase .Name}}PolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwan{{camelCase .Name}}PolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value)}}
					{{- if eq .Type "List"}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value)}}
					{{- if eq .Type "List"}}
					{{- $clist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .TfOnly) (not .Value)}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}_policy_definition.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}_policy_definition.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else}}
					resource.TestCheckResourceAttr("data.sdwan_{{snakeCase $name}}_policy_definition.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

const testAccDataSourceSdwan{{camelCase .Name}}PolicyDefinitionConfig = `

resource "sdwan_{{snakeCase $name}}_policy_definition" "test" {
  name = "TF_TEST_MIN"
  description = "Terraform integration test"
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
{{- if eq .Type "List"}}
  {{.TfName}} = [{
    {{- range  .Attributes}}
    {{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
	{{- if eq .Type "List"}}
	{{.TfName}} = [{
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .TfOnly) (not .Value)}}
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

data "sdwan_{{snakeCase .Name}}_policy_definition" "test" {
  id = sdwan_{{snakeCase $name}}_policy_definition.test.id
}
`
