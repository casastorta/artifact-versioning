// Package flags parses out the input parameters to the commandline utility
package flags

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/casastorta/artifact-versioning/config"
)

// InputTime time.Time object parsed from input string
var InputTime string

// InputLocation time.Location object parsed from input string
var InputLocation string

// InputFallbackEnabled bool object parsed from the input string
var InputFallbackEnabled bool

// InputBranching Branching detection method, parsed from input string
var InputBranching string

// InputFallBackMechanisms Branching detection methods used for fallback
var InputFallBackMechanisms string

var InputMicroVersion int64

func init() {
	currYear, currMonth, currDay := time.Now().Date()
	defaultDate := fmt.Sprintf("%04d-%02d-%02d", currYear, int(currMonth), currDay)

	flagTime := flag.String("t", defaultDate,
		"Time you want to use in YYYY-MM-DD format (default will be current time)")
	flagLocation := flag.String("l", config.DefaultTimeZoneLocation,
		"Timezone/location info")
	flagBranchingDetection := flag.String("b", config.DefaultBranchingDetection,
		"Branching detections mechanism to use")
	flagFallbackBranchingEnabled := flag.Bool("f", false,
		"Use fallback mechanisms to try and detect branch name (default is false [not set])")
	flagFallBackBranchingMechanisms := flag.String("fm", "",
		"Fallback mechanisms to be used for branching detection if the main one fails")
	flagMicroVersion := flag.Int64("m", 0, "Micro version number (i. e. build number)")

	flag.Parse()

	InputTime = strings.ToLower(*flagTime)
	InputLocation = strings.ToLower(*flagLocation)
	InputBranching = strings.ToLower(*flagBranchingDetection)
	InputFallBackMechanisms = strings.ToLower(*flagFallBackBranchingMechanisms)
	InputFallbackEnabled = *flagFallbackBranchingEnabled
	InputMicroVersion = *flagMicroVersion

	if InputFallbackEnabled && InputFallBackMechanisms == "" {
		InputFallbackEnabled = false
		_, _ = fmt.Fprintln(os.Stderr,
			"WARNING: Fallback will be disabled as no fallback mechanisms are defined")
	}
}
