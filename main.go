package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/casastorta/artifact-versioning/cmdline/flags"
	"github.com/casastorta/artifact-versioning/cmdline/util"
	"github.com/casastorta/artifact-versioning/dateutil"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
)

func main() {
	c, err := typeconfig.Config(
		flags.InputTime, flags.InputLocation, flags.InputBranching, flags.InputFallbackEnabled,
		flags.InputFallBackMechanisms)
	if err != nil {
		panic(fmt.Errorf("error on parsing input parameters into config directives: %w", err))
	}
	if c == nil {
		message := fmt.Sprintf(
			"Could not instance Config with the following parameters: %s, %s, %s, %t, %s",
			flags.InputTime, flags.InputLocation, flags.InputBranching, flags.InputFallbackEnabled,
			flags.InputFallBackMechanisms)
		_, _ = fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}

	// Figure out YearWeek part of the version string
	inputTime := c.TimeValue
	inputLocation := c.TimeZoneLocation
	yearWeekValue := dateutil.GetYearWeek(&inputLocation, &inputTime)

	branchName, err2 := util.GetBranchName(c)
	if err2 != nil {
		message := fmt.Sprintf("error on retrieving branch name: %s", err2)
		_, _ = fmt.Fprintln(os.Stderr, message)
		os.Exit(2)
	}
	branchNameItems := util.SplitBranchName(branchName)
	finalBranchName := util.OutputBranchName(branchNameItems)

	var outputValue string
	switch b := finalBranchName == ""; b {
	case true:
		outputValue = yearWeekValue
	default:
		outputValue = fmt.Sprintf("%s.%s", yearWeekValue, finalBranchName)
	}

	fmt.Printf("%s.%s\n", outputValue, strconv.FormatInt(flags.InputMicroVersion, 10))
	os.Exit(0)
}
