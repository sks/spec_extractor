package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Spec Extractor", func() {

	Context("given a directory get all specs out of the child directories", func() {

		It("executes and prints out the specs", func() {
			command := exec.Command(pathToMain,
				"-d", "fixtures/release",
			)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			Eventually(session.Out).Should(gbytes.Say("job_1.databases.db_scheme"))
		})
	})

	Context("failure cases", func() {
		It("errors in case the directory to traverse is not passed in", func() {
			command := exec.Command(pathToMain)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))

			Expect(session.Err.Contents()).To(ContainSubstring("Please pass the directory to traverse with \"-d\" flag"))
		})

		It("Errors when the given folder does not exists", func() {
			command := exec.Command(pathToMain,
				"-d", "non-existing-folder",
			)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))

			Expect(session.Err).Should(gbytes.Say("no such file or directory"))
		})

		It("Errors when there is a invalid spec file in invalid release", func() {
			command := exec.Command(pathToMain,
				"-d", "fixtures/invalid_release",
			)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(1))

			Expect(session.Err).Should(gbytes.Say("is not a valid spec file, yaml: unmarshal errors"))
		})
	})
})
