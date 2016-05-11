package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("yamlmatcher", func() {

	Context("YAML_matcher", func() {
		It("diffs the 2 given YAMLs", func() {
			command := exec.Command(pathToMain,
				"fixtures/old_yaml.yml",
				"fixtures/new_yaml.yml")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			Expect(session.Out.Contents()).To(ContainSubstring("old-value1"))
			Expect(session.Out.Contents()).To(ContainSubstring("new-value1"))
		})
	})
})
