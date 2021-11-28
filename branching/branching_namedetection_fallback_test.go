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

/* Black box testing for the branching name fallback detection methods */
var _ = Describe("Branching Name Detection with fallbacks", func() {
	var (
		t            *typeconfig.TypeConfig
		gitCommand   branching.TCommandLine
		cmd          string
		par          []string
		mockCtrl     *gomock.Controller
		mockCmder    *mockCommander.MockICommand
		mockExecutor *branching.TCommandRunner

		errorCannotIdentifyBranch = errors.New("could not fetch branch name via specified methods")
		bambooBranchVariableName  = branching.EnvVariableBamboo
		jenkinsBranchVariableName = branching.EnvVariableJenkins
	)

	BeforeEach(func() {
		_ = os.Unsetenv(bambooBranchVariableName)  // Clear out Bamboo branch name variable
		_ = os.Unsetenv(jenkinsBranchVariableName) // Clear out Jenkins branch name variable

		gitCommand = branching.CommandGit()
		cmd = gitCommand.Command
		par = strings.Fields(gitCommand.Params)

		mockCtrl = gomock.NewController(GinkgoT())
		mockCmder = mockCommander.NewMockICommand(mockCtrl)
		mockExecutor = &branching.TCommandRunner{Command: mockCmder}

		t, _ = config.DefaultConfig()
		t.BranchingDetection = capabilities.Git
		t.BranchingFallBackDetection = true
		_ = t.SetBranchingFallbackDetectionMechanisms(capabilities.Jenkins, capabilities.Bamboo)

		branching.CommandRunner = *mockExecutor
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Methods hit when fallbacks defined", func() {
		It("should detect git branch", func() {
			v := "test_git_branch"
			vb := []byte(v)
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, nil)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(branchName).To(Equal(v))
		})

		It("should return secondary method branch name", func() {
			b := "test-jenkins-branch"
			_ = os.Setenv(jenkinsBranchVariableName, b)
			v := ""
			vb := []byte(v)
			me := errors.New("mock error")
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, me)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(branchName).To(Equal(b))
		})

		It("should return tertiary method branch name", func() {
			b := "test-bamboo-branch"
			_ = os.Setenv(bambooBranchVariableName, b)
			v := ""
			vb := []byte(v)
			me := errors.New("mock error")
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, me)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(BeNil())
			Expect(branchName).To(Equal(b))
		})
	})

	Describe("When no method returns anything useful", func() {
		It("should raise an error", func() {
			v := ""
			vb := []byte(v)
			me := errors.New("mock error")
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, me)
			branchName, err := branching.GetBranchName(t)

			Expect(err).To(Equal(errorCannotIdentifyBranch))
			Expect(branchName).To(BeEmpty())
		})
	})
})
