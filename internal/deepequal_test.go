package deepequal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sks/yamlmatcher/internal"
)

var _ = Describe("DeepEqual", func() {
	Context("Equal", func() {
		It("return empty Comparison when comparing null", func() {

			comparison, err := deepequal.Compare(nil, nil)
			Expect(err).NotTo(HaveOccurred())
			Expect(len(comparison.Removed)).To(Equal(0))
			Expect(len(comparison.Added)).To(Equal(0))
			Expect(len(comparison.Modified)).To(Equal(0))

		})
	})
})
