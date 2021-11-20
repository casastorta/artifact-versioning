package util

import (
	"regexp"
	"strings"

	"github.com/casastorta/artifact-versioning/branching"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
)

// GetBranchName Pull out branch name, however appropriate, return it as lowered string
func GetBranchName(c *typeconfig.TypeConfig) (string, error) {
	rawBranchName, err := branching.GetBranchName(c)
	if err != nil {
		return "", err
	}
	rawBranchName = strings.ToLower(rawBranchName)
	if rawBranchName == "master" || rawBranchName == "main" {
		rawBranchName = ""
	}
	return rawBranchName, nil
}

// SplitBranchName Splits branch name by any non-alphanumerics in there
func SplitBranchName(branchName string) []string {
	re := regexp.MustCompile("([^a-zA-Z0-9])")
	sl := re.Split(branchName, -1)
	return sl
}

// OutputBranchName Combines splti branch name elements into the string appendable to the YYYY.MM.* part
func OutputBranchName(splitParts []string) string {
	var outputString string
	for idx, part := range splitParts {
		if idx > 1 {
			break
		}
		if outputString != "" {
			outputString += "."
		}
		switch l := len(part) >= 6; l {
		case true:
			outputString += part[:6]
		default:
			outputString += part
		}
	}
	return outputString
}
