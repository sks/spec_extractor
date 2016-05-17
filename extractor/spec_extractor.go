package extractor

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Spec struct {
	Description string      `yaml:"description"`
	Default     interface{} `yaml:"default"`
}

type SpecExtractor struct {
}

type SpecFile struct {
	Properties map[string]Spec `yaml:"properties"`
}

func NewSpecExtractor() SpecExtractor {
	return SpecExtractor{}
}

func (s SpecExtractor) Extract(specFiles []string) (map[string]Spec, error) {
	var returnValue map[string]Spec
	returnValue = make(map[string]Spec)

	for _, specFile := range specFiles {
		specFileContent, err := ioutil.ReadFile(specFile)
		if err != nil {
			return returnValue, err
		}
		var spec SpecFile
		err = yaml.Unmarshal([]byte(specFileContent), &spec)
		if err != nil {
			return returnValue, fmt.Errorf("%q is not a valid spec file, %s", specFile, err)
		}

		for key, value := range spec.Properties {
			returnValue[key] = value
		}
	}
	return returnValue, nil
}
