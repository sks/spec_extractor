package main

import (
	"io/ioutil"
	"os"

	"github.com/cloudfoundry-incubator/candiedyaml"
)

func main() {

	var oldYAMLPath, newYAMLPath string
	oldYAMLPath = os.Args[1]
	newYAMLPath = os.Args[2]
	oldYAMLContent, err := ioutil.ReadFile(oldYAMLPath)
	if err != nil {
		panic(err)
	}

	newYAMLContent, err := ioutil.ReadFile(newYAMLPath)
	if err != nil {
		panic(err)
	}
	var oldYAML interface{}
	var newYAML interface{}

	candiedyaml.Unmarshal(oldYAMLContent, &oldYAML)
	candiedyaml.Unmarshal(newYAMLContent, &newYAML)

}
