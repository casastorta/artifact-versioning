package config

import (
	"time"

	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
)

// DefaultTime Default time to be used, if none provided
var DefaultTime = time.Now().Format("2006-01-02")

// DefaultTimeZoneLocation Default (safe to use) time-zone location
const DefaultTimeZoneLocation string = "UTC"

// DefaultBranchingDetection Default branching detection mechanism
const DefaultBranchingDetection = capabilities.Jenkins

// DefaultBranchingFallBackDetection Default value for branching detection fallback ability
const DefaultBranchingFallBackDetection bool = false

// DefaultFallbackBranchingDetectionMechanisms Default value for branching detection mechanisms fallbacks
const DefaultFallbackBranchingDetectionMechanisms string = ""

// DefaultConfig Default (safe) config parameters
func DefaultConfig() (*typeconfig.TypeConfig, error) {
	return typeconfig.Config(
		DefaultTime,
		DefaultTimeZoneLocation,
		DefaultBranchingDetection,
		DefaultBranchingFallBackDetection,
		DefaultFallbackBranchingDetectionMechanisms)
}
