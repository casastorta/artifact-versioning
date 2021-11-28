package branching

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/* White box testing of the environment variables dependency for pulling branch name */
var _ = Describe("BranchingNameDetection via environment variables", func() {
	var (
		bambooBranchVariableName  = EnvVariableBamboo
		jenkinsBranchVariableName = EnvVariableJenkins
		testBranchName            = "test_branch_name"
	)

	BeforeEach(func() {
		_ = os.Unsetenv(bambooBranchVariableName)  // Clear out Bamboo branch name variable
		_ = os.Unsetenv(jenkinsBranchVariableName) // Clear out Jenkins branch name variable
	})

	Describe("Jenkins variable name should", func() {
		It("should return empty when variable not set", func() {
			branchName := getJenkinsBranchName()
			Expect(branchName).To(Equal(""))
		})

		It("should return branch name when variable is set", func() {
			err := os.Setenv(jenkinsBranchVariableName, testBranchName)
			Expect(err).To(BeNil())

			branchName := getJenkinsBranchName()
			Expect(branchName).To(Equal(testBranchName))
		})
	})

	Describe("Bamboo variable name should", func() {
		It("should return empty when variable not set", func() {
			branchName := getBambooBranchName()
			Expect(branchName).To(Equal(""))
		})

		It("should return branch name when variable is set", func() {
			err := os.Setenv(bambooBranchVariableName, testBranchName)
			Expect(err).To(BeNil())

			branchName := getBambooBranchName()
			Expect(branchName).To(Equal(testBranchName))
		})
	})
})
