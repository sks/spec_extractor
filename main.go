package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/sks/spec_extractor/extractor"
)

func main() {

	logger := log.New(os.Stderr, "", 0)

	var directoryToCheck string

	flag.StringVar(&directoryToCheck, "d", "", "Directory to traverse")
	flag.Parse()

	if len(directoryToCheck) == 0 {
		logger.Fatalf("Please pass the directory to traverse with %q flag ", "-d")
	}

	var specFiles []string

	err := filepath.Walk(directoryToCheck, extractor.FindSpecFiles(&specFiles))
	if nil != err {
		logger.Fatal(err)
	}
	logger.Printf("Obtained %d spec files from %q", len(specFiles), directoryToCheck)

	specExtractor := extractor.NewSpecExtractor()
	specs, err := specExtractor.Extract(specFiles)
	if nil != err {
		logger.Fatal(err)
	}
	logger.Printf("Obtained %d specs from %q", len(specs), directoryToCheck)
	d, _ := yaml.Marshal(&specs)
	fmt.Println(fmt.Sprintf("%s", d))

}
