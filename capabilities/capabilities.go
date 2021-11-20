package capabilities

import (
	. "github.com/casastorta/artifact-versioning/types/uniquestringlist"
)

// Git We will do branch detection based on the `git` command output
// TODO: which command exactly?
const Git string = "git"

// Bamboo We will do branch detection based on the environment variable set by Bamboo
// (`planRepository.branchDisplayName`)
const Bamboo string = "bamboo"

// Jenkins We will do branch detection based on the environment variable set by Jenkins (`GIT_BRANCH`)
const Jenkins string = "jenkins"

// BranchingDetectionMethods should contain all the constants from the above section
var branchingDetectionMethods = UniqueStringList{
	Git,
	Bamboo,
	Jenkins,
}

// BranchingDetectionMethods Retrieve possible branching detection methods
func BranchingDetectionMethods() UniqueStringList {
	return branchingDetectionMethods
}
