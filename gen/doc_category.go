//go:build ignore

package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath = "./gen/definitions/feature_templates/"
)

type YamlConfig struct {
	Name string `yaml:"name"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}

var extraDocs = map[string]string{
	"cli_device_template":            "Device Templates",
	"feature_device_template":        "Device Templates",
	"attach_feature_device_template": "Device Templates",
}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

func main() {
	items, _ := ioutil.ReadDir(definitionsPath)
	configs := make([]YamlConfig, len(items))

	// Load configs
	for i, filename := range items {
		yamlFile, err := ioutil.ReadFile(filepath.Join(definitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		configs[i] = config
	}

	// Update feature template doc category
	for i := range configs {
		for _, path := range docPaths {
			filename := path + SnakeCase(configs[i].Name) + "_feature_template.md"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "Feature Templates"`)

			ioutil.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Update extra (non-feature templates) doc categories
	for doc, cat := range extraDocs {
		for _, path := range docPaths {
			filename := path + doc + ".md"
			content, err := ioutil.ReadFile(filename)
			if err == nil {
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				ioutil.WriteFile(filename, []byte(s), 0644)
			}
		}
	}
}
