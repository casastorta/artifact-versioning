package typeconfig

import (
	"fmt"
	"strings"
	"time"

	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/types/uniquestringlist"
)

// TypeConfig data type to hold all variable config directives
type TypeConfig struct {
	TimeValue                            time.Time
	TimeZoneLocation                     time.Location
	BranchingDetection                   string
	BranchingFallBackDetection           bool
	branchingFallBackDetectionMechanisms uniquestringlist.UniqueStringList
}

// Config Instance of the configuration directives with sanity checking in place
func Config(sTime string, location string, branchingMethod string, branchingFallback bool,
	branchingFallbackMethods string) (*TypeConfig, error) {
	// Parse location from string
	loc, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	// Parse time at location from string
	tp, errTimeParse := time.ParseInLocation("2006-01-02", sTime, loc)
	if errTimeParse != nil {
		return nil, errTimeParse
	}

	config := new(TypeConfig)

	config.TimeValue = tp
	config.TimeZoneLocation = *loc
	config.BranchingDetection = branchingMethod
	config.BranchingFallBackDetection = branchingFallback
	errFbm := config.SetBranchingFallbackDetectionMechanismsFromString(branchingFallbackMethods)
	if errFbm != nil {
		return nil, errFbm
	}

	return config, nil
}

// BranchingFallbackDetectionMechanisms returns the private struct member branchingFallBackDetectionMechanisms
func (tc *TypeConfig) BranchingFallbackDetectionMechanisms() uniquestringlist.UniqueStringList {
	return tc.branchingFallBackDetectionMechanisms
}

// SetBranchingFallbackDetectionMechanismsFromString sets the private struct member branchingFallBackDetectionMechanisms
// from the string input of comma-separated values
func (tc *TypeConfig) SetBranchingFallbackDetectionMechanismsFromString(s string) error {
	ss := strings.Split(s, ",")
	return tc.SetBranchingFallbackDetectionMechanisms(ss...)
}

// SetBranchingFallbackDetectionMechanisms sets the private struct member branchingFallBackDetectionMechanisms
// from the multiple strings parameters
func (tc *TypeConfig) SetBranchingFallbackDetectionMechanisms(s ...string) error {
	var so uniquestringlist.UniqueStringList
	bm := capabilities.BranchingDetectionMethods()
	for _, st := range s {
		st = strings.TrimSpace(st)
		// Sanity checks
		// ...is it empty? Can't be, skip it
		if st == "" {
			continue
		}
		// ...is it the same as default mechanism? Can't be
		if st == tc.BranchingDetection {
			return fmt.Errorf("cannot add default mechanism '%s' to the fallback mechanisms", st)
		}
		// ...is it present in the detection mechanisms overall? Must be.
		if !bm.Contains(st) {
			return fmt.Errorf("cannot add non-existing '%s' to fallback detection methods", st)
		}

		err := so.Append(st)
		if err != nil {
			return err
		}
	}
	tc.branchingFallBackDetectionMechanisms = so
	return nil
}
