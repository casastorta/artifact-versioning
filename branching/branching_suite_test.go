package branching_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBranching(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Branching Suite")
}