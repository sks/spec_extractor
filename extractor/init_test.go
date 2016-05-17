package extractor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDelta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "extractor")
}
