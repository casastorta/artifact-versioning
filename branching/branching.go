/*
Package branching implements logic to detect SCM branch, either via launching SCM tool or via environment variables
(mostly for CI/CD systems)
*/
package branching

import (
	"fmt"
	"time"

	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
	"github.com/casastorta/artifact-versioning/types/uniquestringlist"
)

// detectionOrder Order in which we will try to detect branches, first one wins!
var detectionOrder = capabilities.BranchingDetectionMethods()

// AllowFallbackDetection Is fallback on detection methods allowed, or we do only one?
var AllowFallbackDetection = true

// GetDetectionOrder Retrieve detection order as a list of strings
func GetDetectionOrder() uniquestringlist.UniqueStringList {
	return detectionOrder
}

// SetDetectionOrder Set detection order from the list of strings
func SetDetectionOrder(methods ...string) error {
	var NewOrder uniquestringlist.UniqueStringList
	bm := capabilities.BranchingDetectionMethods()
	for _, method := range methods {
		if !bm.Contains(method) {
			return fmt.Errorf("method '%s' not in the list of allowed branch detection methods", method)
		}
		err := NewOrder.Append(method)
		if err != nil {
			return err
		}
	}
	detectionOrder = NewOrder
	return nil
}

func GetYearWeek(c *typeconfig.TypeConfig) string {
	t := time.Now()
	tl := t.In(&c.TimeZoneLocation)
	year, week := tl.ISOWeek()
	return fmt.Sprintf("%d.%d", year, week)
}
