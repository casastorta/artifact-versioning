package dateutil_test

import (
	"time"

	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/dateutil"
	"github.com/casastorta/artifact-versioning/types/typeconfig"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dateutil", func() {
	var (
		configInstance *typeconfig.TypeConfig
		location       time.Location
	)

	BeforeEach(func() {
		configInstance, _ = config.DefaultConfig()
		location = configInstance.TimeZoneLocation
	})

	Describe("With default config instance", func() {
		yearWeek := dateutil.GetCurrentYearWeek(&location)

		It("should return year and week", func() {
			Ω(yearWeek).Should(MatchRegexp("[1-9][0-9][0-9][0-9]\\.[0-5][0-9]"))
		})
	})

	Describe("With custom time passed on", func() {
		dateTime := time.Date(1980, 4, 28, 0o3, 23, 58, 0, time.UTC)
		yearWeek := dateutil.GetYearWeek(&location, &dateTime)

		It("should return appropriate year and week", func() {
			Expect(yearWeek).To(Equal("1980.18"))
		})
	})
})
