package capabilities_test

import (
	"github.com/casastorta/artifact-versioning/capabilities"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Capabilities", func() {
	It("should contain all capabilities", func() {
		c := capabilities.BranchingDetectionMethods()
		containsGit := c.Contains(capabilities.Git)
		containsBamboo := c.Contains(capabilities.Bamboo)
		containsJenkins := c.Contains(capabilities.Jenkins)

		Expect(containsGit).To(BeTrue())
		Expect(containsBamboo).To(BeTrue())
		Expect(containsJenkins).To(BeTrue())
	})
})
