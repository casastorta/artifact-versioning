package util_test

import (
	"os"
	"strings"

	"github.com/casastorta/artifact-versioning/branching"
	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/cmdline/util"
	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/types/typeconfig"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	Describe("SplitBranchName function", func() {
		It("should split the branch name into string list", func() {
			demo_branch_name := "JIRA-12345_New_feature.we're_adding"
			split_should_be := []string{"JIRA", "12345", "New", "feature", "we", "re", "adding"}
			split_branch_name := util.SplitBranchName(demo_branch_name)
			count_parts := len(split_branch_name)

			Expect(count_parts).To(Equal(7))
			Expect(split_branch_name).To(Equal(split_should_be))
		})
	})

	Describe("OutputBranchName function", func() {
		It("should return empty string if nothing sent", func() {
			inputStringList := []string{}
			v := util.OutputBranchName(inputStringList)
			Expect(v).To(BeEmpty())
		})

		It("should return formatted string if something is sent", func() {
			inputStringList := []string{"jira", "245645645221", "some", "feature", "name"}
			expectedReturn := "jira.245645"
			v := util.OutputBranchName(inputStringList)
			Expect(v).To(Equal(expectedReturn))
		})

		It("should return string trimmed to 6 characters if one item is sent", func() {
			inputStringList := []string{"jira245645645221some feature name"}
			expectedReturn := "jira24"
			v := util.OutputBranchName(inputStringList)
			Expect(v).To(Equal(expectedReturn))
		})
	})

	Describe("GetBranchName function", func() {
		var t *typeconfig.TypeConfig

		BeforeEach(func() {
			_ = os.Unsetenv(branching.EnvVariableBamboo)
			t, _ = config.DefaultConfig()

			t.BranchingDetection = capabilities.Bamboo
			t.BranchingFallBackDetection = false
			_ = t.SetBranchingFallbackDetectionMechanisms("")
		})

		AfterEach(func() {
			t = nil
		})

		It("should return branch name converted to lower case", func() {
			demo_branch_name := "Some Branch Name goes here"
			demo_branch_name_lowered := strings.ToLower(demo_branch_name)

			_ = os.Setenv(branching.EnvVariableBamboo, demo_branch_name)
			v, err := util.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(v).To(Equal(demo_branch_name_lowered))
		})

		It("should return empty string of branch is master", func() {
			demoBranchName := "master"

			_ = os.Setenv(branching.EnvVariableBamboo, demoBranchName)
			v, err := util.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(v).To(BeEmpty())
		})

		It("should return empty string of branch is main", func() {
			demoBranchName := "main"

			_ = os.Setenv(branching.EnvVariableBamboo, demoBranchName)
			v, err := util.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(v).To(BeEmpty())
		})

		It("should fail with error if branch not detectable", func() {
			v, err := util.GetBranchName(t)
			Expect(err).To(Not(BeNil()))
			Expect(v).To(BeEmpty())
		})
	})
})
