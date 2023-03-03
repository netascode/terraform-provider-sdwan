//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
)

const (
	featureTemplateDefinitionsPath = "./gen/definitions/feature_templates/"
	featureTemplateModelsPath      = "./gen/models/feature_templates/"
	policyObjectDefinitionsPath    = "./gen/definitions/policy_objects/"
	providerTemplate               = "./gen/templates/provider.go"
	providerLocation               = "./internal/provider/provider.go"
	changelogTemplate              = "./gen/templates/changelog.md.tmpl"
	changelogLocation              = "./templates/guides/changelog.md.tmpl"
	changelogOriginal              = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var featureTemplateTemplates = []t{
	{
		path:   "./gen/templates/feature_templates/model.go",
		prefix: "./internal/provider/model_sdwan_",
		suffix: "_feature_template.go",
	},
	{
		path:   "./gen/templates/feature_templates/data_source.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_feature_template.go",
	},
	{
		path:   "./gen/templates/feature_templates/data_source_test.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_feature_template_test.go",
	},
	{
		path:   "./gen/templates/feature_templates/resource.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_feature_template.go",
	},
	{
		path:   "./gen/templates/feature_templates/resource_test.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_feature_template_test.go",
	},
	{
		path:   "./gen/templates/feature_templates/data-source.tf",
		prefix: "./examples/data-sources/sdwan_",
		suffix: "_feature_template/data-source.tf",
	},
	{
		path:   "./gen/templates/feature_templates/resource.tf",
		prefix: "./examples/resources/sdwan_",
		suffix: "_feature_template/resource.tf",
	},
	{
		path:   "./gen/templates/feature_templates/import.sh",
		prefix: "./examples/resources/sdwan_",
		suffix: "_feature_template/import.sh",
	},
}

var policyObjectTemplates = []t{
	{
		path:   "./gen/templates/policy_objects/model.go",
		prefix: "./internal/provider/model_sdwan_",
		suffix: "_policy_object.go",
	},
	{
		path:   "./gen/templates/policy_objects/data_source.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_policy_object.go",
	},
	{
		path:   "./gen/templates/policy_objects/data_source_test.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_policy_object_test.go",
	},
	{
		path:   "./gen/templates/policy_objects/resource.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_policy_object.go",
	},
	{
		path:   "./gen/templates/policy_objects/resource_test.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_policy_object_test.go",
	},
	{
		path:   "./gen/templates/policy_objects/data-source.tf",
		prefix: "./examples/data-sources/sdwan_",
		suffix: "_policy_object/data-source.tf",
	},
	{
		path:   "./gen/templates/policy_objects/resource.tf",
		prefix: "./examples/resources/sdwan_",
		suffix: "_policy_object/resource.tf",
	},
	{
		path:   "./gen/templates/policy_objects/import.sh",
		prefix: "./examples/resources/sdwan_",
		suffix: "_policy_object/import.sh",
	},
}

type YamlConfig struct {
	Name           string                `yaml:"name"`
	Model          string                `yaml:"model"`
	Type           string                `yaml:"type"`
	MinimumVersion string                `yaml:"minimum_version"`
	DsDescription  string                `yaml:"ds_description"`
	ResDescription string                `yaml:"res_description"`
	DocCategory    string                `yaml:"doc_category"`
	ExcludeTest    bool                  `yaml:"exclude_test"`
	Attributes     []YamlConfigAttribute `yaml:"attributes"`
}

type YamlConfigAttribute struct {
	ModelName       string                `yaml:"model_name"`
	TfName          string                `yaml:"tf_name"`
	Type            string                `yaml:"type"`
	ObjectType      string                `yaml:"object_type"`
	ModelTypeString bool                  `yaml:"model_type_string"`
	DataPath        []string              `yaml:"data_path"`
	Keys            []string              `yaml:"keys"`
	Id              bool                  `yaml:"id"`
	Variable        bool                  `yaml:"variable"`
	Mandatory       bool                  `yaml:"mandatory"`
	WriteOnly       bool                  `yaml:"write_only"`
	ExcludeTest     bool                  `yaml:"exclude_test"`
	ExcludeExample  bool                  `yaml:"exclude_example"`
	Description     string                `yaml:"description"`
	Example         string                `yaml:"example"`
	EnumValues      []string              `yaml:"enum_values"`
	MinInt          int64                 `yaml:"min_int"`
	MaxInt          int64                 `yaml:"max_int"`
	MinFloat        float64               `yaml:"min_float"`
	MaxFloat        float64               `yaml:"max_float"`
	StringPatterns  []string              `yaml:"string_patterns"`
	StringMinLength int64                 `yaml:"string_min_length"`
	StringMaxLength int64                 `yaml:"string_max_length"`
	DefaultValue    string                `yaml:"default_value"`
	Attributes      []YamlConfigAttribute `yaml:"attributes"`
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":  ToGoName,
	"camelCase": CamelCase,
	"snakeCase": SnakeCase,
	"sprintf":   fmt.Sprintf,
	"toLower":   strings.ToLower,
}

func parseFeatureTemplateAttribute(attr *YamlConfigAttribute, model gjson.Result) {
	var r gjson.Result
	if model.Get("fields").Exists() {
		r = model.Get("fields.#(key==\"" + attr.ModelName + "\")#")
	} else {
		r = model.Get("children.#(key==\"" + attr.ModelName + "\")#")
	}
	if len(r.Array()) > 1 {
		found := false
		for _, item := range r.Array() {
			if len(item.Get("dataPath").Array()) == 0 {
				fmt.Printf("WARNING: Non-unique key without data_path: %s\n", attr.ModelName)
			}
			match := true
			for _, dp := range item.Get("dataPath").Array() {
				if !contains(attr.DataPath, dp.String()) {
					match = false
					break
				}
			}
			if match {
				r = item
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("Could not find element in schema: %v\n%v\n\n", attr.ModelName, model))
		}
	} else if len(r.Array()) == 1 {
		r = r.Array()[0]
	} else {
		panic(fmt.Sprintf("Could not find element in schema: %v\n%v\n\n", attr.ModelName, model))
	}

	if attr.Description == "" && r.Get("details").String() != "" {
		attr.Description = r.Get("details").String()
	}
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}
	if len(attr.DataPath) == 0 && len(r.Get("dataPath").Array()) > 0 {
		paths := make([]string, len(r.Get("dataPath").Array()))
		for i, v := range r.Get("dataPath").Array() {
			paths[i] = v.String()
		}
		attr.DataPath = paths
	}
	if len(attr.Keys) == 0 && len(r.Get("primaryKeys").Array()) > 0 {
		keys := make([]string, len(r.Get("primaryKeys").Array()))
		for i, v := range r.Get("primaryKeys").Array() {
			keys[i] = v.String()
		}
		attr.Keys = keys
	}
	attr.ObjectType = r.Get("objectType").String()
	if r.Get("objectType").String() == "object" || r.Get("objectType").String() == "node-only" {
		t := r.Get("dataType.type").String()
		if r.Get("dataType").String() == "string" {
			t = "string"
		}
		if contains([]string{"string", "passphrase", "restrictedPassphrase", "datetimelocal", "ip", "ipv4", "ipv6", "ipv4-prefix", "ipv6-prefix", "dnsHostName", "interfaceList", "tlocExtension", "xConnect", "mac"}, t) {
			attr.Type = "String"
			if r.Get("dataType.minLength").Exists() {
				attr.StringMinLength = r.Get("dataType.minLength").Int()
			}
			if r.Get("dataType.maxLength").Exists() {
				attr.StringMaxLength = r.Get("dataType.maxLength").Int()
			}
		} else if t == "stringList" {
			attr.Type = "ListString"
		} else if t == "enum" || t == "enumList" {
			attr.Type = "String"
			for _, v := range r.Get("dataType.values").Array() {
				attr.EnumValues = append(attr.EnumValues, v.Get("key").String())
			}
		} else if t == "radioButtonList" {
			attr.Type = "String"
			for _, v := range r.Get("dataType.values").Array() {
				attr.EnumValues = append(attr.EnumValues, v.Get("value").String())
			}
		} else if contains([]string{"number", "numberFixedInterval"}, t) {
			attr.Type = "Int64"
			if r.Get("dataType.min").Exists() {
				attr.MinInt = r.Get("dataType.min").Int()
			}
			if r.Get("dataType.max").Exists() {
				attr.MaxInt = r.Get("dataType.max").Int()
			}
		} else if t == "float" {
			attr.Type = "Float64"
			if r.Get("dataType.min").Exists() {
				attr.MinFloat = r.Get("dataType.min").Float()
			}
			if r.Get("dataType.max").Exists() {
				attr.MaxFloat = r.Get("dataType.max").Float()
			}
		} else if t == "boolean" {
			attr.Type = "Bool"
		}
	} else if r.Get("objectType").String() == "list" {
		attr.Type = "ListString"
	}
	if r.Get("dataType.default").Exists() {
		attr.DefaultValue = r.Get("dataType.default").String()
	}
	types := r.Get("optionType").Array()
	ignore := false
	variable := false
	for t := range types {
		if types[t].String() == "ignore" {
			ignore = true
		} else if types[t].String() == "variable" {
			variable = true
		}
	}
	if !ignore && r.Get("objectType").String() != "tree" {
		attr.Mandatory = true
	}
	if variable && r.Get("objectType").String() != "tree" {
		attr.Variable = true
	}
	if r.Get("objectType").String() == "tree" && len(attr.Attributes) > 0 {
		attr.Type = "List"
		for a := range attr.Attributes {
			parseFeatureTemplateAttribute(&attr.Attributes[a], r)
		}
	}
}

func augmentFeatureTemplateConfig(config *YamlConfig) {
	if config.Model == "" {
		config.Model = SnakeCase(config.Name)
	}
	modelPath := featureTemplateModelsPath + config.Model + ".json"

	modelBytes, err := os.ReadFile(modelPath)
	if err != nil {
		log.Fatalf("Error reading model '%s': %v", modelPath, err)
	}

	model := gjson.ParseBytes(modelBytes)

	for ia := range config.Attributes {
		parseFeatureTemplateAttribute(&config.Attributes[ia], model)
	}

	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s feature template.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s feature template.", config.Name)
	}
}

func augmentPolicyObjectAttribute(attr *YamlConfigAttribute) {
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}
	if attr.Type == "List" {
		for a := range attr.Attributes {
			augmentPolicyObjectAttribute(&attr.Attributes[a])
		}
	}
}

func augmentPolicyObjectConfig(config *YamlConfig) {
	for ia := range config.Attributes {
		augmentPolicyObjectAttribute(&config.Attributes[ia])
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s policy object.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s policy object.", config.Name)
	}
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// create output file
	outputFile := filepath.Join(outputPath)
	os.MkdirAll(filepath.Dir(outputFile), 0755)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	f.Write(output.Bytes())
}

func main() {
	featureTemplateFiles, _ := ioutil.ReadDir(featureTemplateDefinitionsPath)
	featureTemplateConfigs := make([]YamlConfig, len(featureTemplateFiles))
	providerConfig := make(map[string][]string)
	providerConfig["FeatureTemplates"] = make([]string, 0)
	providerConfig["PolicyObjects"] = make([]string, 0)

	// Load feature template configs
	for i, filename := range featureTemplateFiles {
		yamlFile, err := os.ReadFile(filepath.Join(featureTemplateDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		featureTemplateConfigs[i] = config
	}

	for i := range featureTemplateConfigs {
		// Augment feature template config by model data
		augmentFeatureTemplateConfig(&featureTemplateConfigs[i])

		// Iterate over templates and render files
		for _, t := range featureTemplateTemplates {
			renderTemplate(t.path, t.prefix+SnakeCase(featureTemplateConfigs[i].Name)+t.suffix, featureTemplateConfigs[i])
		}
		providerConfig["FeatureTemplates"] = append(providerConfig["FeatureTemplates"], featureTemplateConfigs[i].Name)
	}

	policyObjectFiles, _ := ioutil.ReadDir(policyObjectDefinitionsPath)
	policyObjectConfigs := make([]YamlConfig, len(policyObjectFiles))

	// Load policy object configs
	for i, filename := range policyObjectFiles {
		yamlFile, err := os.ReadFile(filepath.Join(policyObjectDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		policyObjectConfigs[i] = config
	}

	for i := range policyObjectConfigs {
		// Augment policy object config
		augmentPolicyObjectConfig(&policyObjectConfigs[i])

		// Iterate over templates and render files
		for _, t := range policyObjectTemplates {
			renderTemplate(t.path, t.prefix+SnakeCase(policyObjectConfigs[i].Name)+t.suffix, policyObjectConfigs[i])
		}
		providerConfig["PolicyObjects"] = append(providerConfig["PolicyObjects"], policyObjectConfigs[i].Name)
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, providerConfig)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
