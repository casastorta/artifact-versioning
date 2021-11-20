package branching

import (
	"errors"
	"strings"

	mockCommander "github.com/casastorta/artifact-versioning/external/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/* White box testing of the git functionality dependency for pulling branch name */
var _ = Describe("BranchingNameDetection via Git commandline", func() {
	var (
		mockCtrl     *gomock.Controller
		mockCmder    *mockCommander.MockICommand
		mockExecutor *TCommandRunner
		cmd          string
		par          []string
	)

	BeforeEach(func() {
		gitCmdLine := CommandGit()
		cmd = gitCmdLine.Command
		par = strings.Fields(gitCmdLine.Params)

		mockCtrl = gomock.NewController(GinkgoT())
		mockCmder = mockCommander.NewMockICommand(mockCtrl)
		mockExecutor = &TCommandRunner{Command: mockCmder}

		CommandRunner = *mockExecutor
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("getGitBranchName", func() {
		It("should return 'test_branch_name' when git returns value", func() {
			v := "test_branch_name"
			vb := []byte(v)
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, nil)
			testBranch := CommandRunner.getGitBranchName()
			Expect(testBranch).To(Equal(v))
		})

		It("should return empty string when something went wrong", func() {
			v := ""
			vb := []byte(v)
			me := errors.New("mock error")
			mockCmder.EXPECT().Execute(cmd, par).Return(vb, me)
			testBranch := CommandRunner.getGitBranchName()
			Expect(testBranch).To(BeEmpty())
		})
	})
})
