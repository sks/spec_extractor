package extractor_test

import (
	"github.com/sks/spec_extractor/extractor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetAllSpecs", func() {

	var specExtractor extractor.SpecExtractor

	BeforeEach(func() {
		specExtractor = extractor.NewSpecExtractor()
	})

	Context("joins all the specs from the given files", func() {
		var specFiles []string

		BeforeEach(func() {
			specFiles = []string{
				"../fixtures/release/jobs/job-1/spec",
				"../fixtures/release/jobs/job-2/spec",
				"../fixtures/release/jobs/job-3/spec",
			}
		})
		It("comes up with slice of all specs", func() {
			specs, err := specExtractor.Extract(specFiles)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(specs)).To(Equal(6))
		})

		Context("failure cases", func() {
			It("the file does not exists", func() {
				_, err := specExtractor.Extract([]string{"non-exsting-file"})
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("open non-exsting-file: no such file or directory"))
			})

			It("the file is not a valid yaml file", func() {
				_, err := specExtractor.Extract([]string{"../main.go"})
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("\"../main.go\" is not a valid spec file, ")))
			})
		})
	})
})
