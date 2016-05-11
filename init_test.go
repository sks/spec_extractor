package main_test

import (
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	pathToMain string
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, ".")
}

var _ = BeforeSuite(func() {
	var (
		err error
	)
	pathToMain, err = gexec.Build("github.com/sks/yamlmatcher")
	Expect(err).NotTo(HaveOccurred())

})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
