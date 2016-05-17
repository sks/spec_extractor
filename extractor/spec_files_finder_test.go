package extractor_test

import (
	"os"

	"github.com/sks/spec_extractor/extractor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindSpecFiles", func() {

	Context("given an existing directory", func() {

		It("returns the file if a valid spec", func() {
			locationToCheck, err := os.Open("../fixtures/release/jobs/job-1/spec")
			Expect(err).NotTo(HaveOccurred())

			folderInfo, err := locationToCheck.Stat()
			Expect(err).NotTo(HaveOccurred())

			var specFiles = []string{}
			err = extractor.FindSpecFiles(&specFiles)(locationToCheck.Name(), folderInfo, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(1).To(Equal(len(specFiles)))
		})
	})

	Context("failure cases", func() {
		It("it ignores directory", func() {
			locationToCheck, err := os.Open("../fixtures/release/config/spec")
			Expect(err).NotTo(HaveOccurred())

			folderInfo, err := locationToCheck.Stat()
			Expect(err).NotTo(HaveOccurred())

			var specFiles = []string{}
			err = extractor.FindSpecFiles(&specFiles)(locationToCheck.Name(), folderInfo, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(0).To(Equal(len(specFiles)))
		})

		It("it ignores any files with name other than spec", func() {
			locationToCheck, err := os.Open("../fixtures/release/jobs/job-1/template")
			Expect(err).NotTo(HaveOccurred())

			folderInfo, err := locationToCheck.Stat()
			Expect(err).NotTo(HaveOccurred())

			var specFiles = []string{}
			err = extractor.FindSpecFiles(&specFiles)(locationToCheck.Name(), folderInfo, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(0).To(Equal(len(specFiles)))
		})

		It("it ignores any specs out of jobs folder", func() {
			locationToCheck, err := os.Open("../fixtures/release/config/something/spec")
			Expect(err).NotTo(HaveOccurred())

			folderInfo, err := locationToCheck.Stat()
			Expect(err).NotTo(HaveOccurred())

			var specFiles = []string{}
			err = extractor.FindSpecFiles(&specFiles)(locationToCheck.Name(), folderInfo, nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(0).To(Equal(len(specFiles)))
		})
	})
})
