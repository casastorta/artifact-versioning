package branching_test

import (
	"errors"
	"os"
	"strings"

	"github.com/casastorta/artifact-versioning/branching"
	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/types/typeconfig"

	mockCommander "github.com/casastorta/artifact-versioning/external/mocks"
	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/* Black box testing for the branching name detection methods */
var _ = Describe("Branching Name Detection without fallbacks", func() {
	var (
		t          *typeconfig.TypeConfig
		gitCommand branching.TCommandLine
		cmd        string
		par        []string

		errorCannotIdentifyBranch = errors.New("could not fetch branch name via specified methods")
		bambooBranchVariableName  = branching.EnvVariableBamboo
		jenkinsBranchVariableName = branching.EnvVariableJenkins
	)

	BeforeEach(func() {
		_ = os.Unsetenv(bambooBranchVariableName)  // Clear out Bamboo branch name variable
		_ = os.Unsetenv(jenkinsBranchVariableName) // Clear out Jenkins branch name variable

		t, _ = config.DefaultConfig()
	})

	Describe("Defining duplicate fallback mechanisms", func() {
		It("should throw an error", func() {
			t.BranchingDetection = capabilities.Jenkins
			t.BranchingFallBackDetection = true
			_ = t.SetBranchingFallbackDetectionMechanisms(capabilities.Bamboo)
			t.BranchingDetection = capabilities.Bamboo
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(Not(BeNil()))
			Expect(branchName).To(BeEmpty())
		})
	})

	Describe("Detecting Bamboo branch", func() {
		BeforeEach(func() {
			t.BranchingDetection = capabilities.Bamboo
			t.BranchingFallBackDetection = false
			_ = t.SetBranchingFallbackDetectionMechanisms("")
		})

		It("should return branch name from Bamboo variable when bamboo branch is set", func() {
			v := "test_bamboo_branch"
			_ = os.Setenv(bambooBranchVariableName, v)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(branchName).To(Equal(v))
		})

		It("should return empty string when Bamboo variable is not set and raise an error", func() {
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(Equal(errorCannotIdentifyBranch))
			Expect(branchName).To(BeEmpty())
		})
	})

	Describe("Detecting Jenkins branch", func() {
		BeforeEach(func() {
			t.BranchingDetection = capabilities.Jenkins
			t.BranchingFallBackDetection = false
			_ = t.SetBranchingFallbackDetectionMechanisms("")
		})

		It("should return branch name from Jenkins variable when bamboo branch is set", func() {
			v := "test_jenkins_branch"
			_ = os.Setenv(jenkinsBranchVariableName, v)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(branchName).To(Equal(v))
		})

		It("should return empty string when Jenkins variable is not set and raise an error", func() {
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(Equal(errorCannotIdentifyBranch))
			Expect(branchName).To(BeEmpty())
		})
	})

	Describe("Detecting Git branch", func() {
		var (
			mockCtrl     *gomock.Controller
			mockCmder    *mockCommander.MockICommand
			mockExecutor *branching.TCommandRunner
		)
		BeforeEach(func() {
			gitCommand = branching.CommandGit()
			cmd = gitCommand.Command
			par = strings.Fields(gitCommand.Params)

			mockCtrl = gomock.NewController(GinkgoT())
			mockCmder = mockCommander.NewMockICommand(mockCtrl)
			mockExecutor = &branching.TCommandRunner{Command: mockCmder}

			t.BranchingDetection = capabilities.Git
			t.BranchingFallBackDetection = false
			_ = t.SetBranchingFallbackDetectionMechanisms("")

			branching.CommandRunner = *mockExecutor
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		It("should return branch when git functionality is working", func() {
			v := "test_git_branch"
			vb := []byte(v)
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, nil)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(branchName).To(Equal(v))
		})

		It("should return empty string when git functionality is not working and raise an error", func() {
			v := ""
			vb := []byte(v)
			me := errors.New("mock error")
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, me)
			branchName, err := branching.GetBranchName(t)

			Expect(branchName).To(BeEmpty())
			Expect(err).To(Equal(errorCannotIdentifyBranch))
		})
	})
})
