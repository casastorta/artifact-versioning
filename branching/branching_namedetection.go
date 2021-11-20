package branching

import (
	"errors"
	"os"
	"strings"

	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/external"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
	"github.com/casastorta/artifact-versioning/types/uniquestringlist"
)

type (
	TCommandRunner struct {
		Command external.ICommand
	}

	TCommandLine struct {
		Command string
		Params  string
	}
)

var (
	CommandRunner = TCommandRunner{&external.TCommand{}}
	commandGit    = TCommandLine{"git", "branch --show-current"}
)

const (
	EnvVariableBamboo  = "planRepository_branchDisplayName"
	EnvVariableJenkins = "GIT_BRANCH"
)

func CommandGit() TCommandLine {
	return commandGit
}

// GetBranchName Returns the branch name, based on the mechanisms defined in the configuration
func GetBranchName(c *typeconfig.TypeConfig) (string, error) {
	// Assemble the list of checking mechanisms
	mechanisms := uniquestringlist.UniqueStringList{}
	_ = mechanisms.Append(c.BranchingDetection)
	if c.BranchingFallBackDetection {
		for _, m := range c.BranchingFallbackDetectionMechanisms() {
			err := mechanisms.Append(m)
			if err != nil {
				return "", err
			}
		}
	}

	// Try all of specified the detection mechanisms as long as one doesn't work
	for _, me := range mechanisms {
		var v string
		switch m := me; m {
		case capabilities.Bamboo:
			v = getBambooBranchName()
		case capabilities.Jenkins:
			v = getJenkinsBranchName()
		default:
			v = CommandRunner.getGitBranchName()
		}
		if v != "" {
			v = strings.TrimSpace(v)
			return v, nil
		}
	}

	// If nothing worked, return the error
	return "", errors.New("could not fetch branch name via specified methods")
}

// getBambooBranchName Return branch name from the environment Bamboo variable
func getBambooBranchName() string {
	branchName := os.Getenv(EnvVariableBamboo)
	return branchName
}

// getJenkinsBranchName Return branch name from the Jenkins environment variable
func getJenkinsBranchName() string {
	branchName := os.Getenv(EnvVariableJenkins)
	return branchName
}

func (cx *TCommandRunner) getGitBranchName() string {
	cmdLine := CommandGit()
	command := cmdLine.Command
	arguments := strings.Fields(cmdLine.Params)
	o, err := cx.Command.Execute(command, arguments...)
	if err != nil {
		return ""
	}
	return string(o)
}
