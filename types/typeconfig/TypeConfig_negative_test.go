package typeconfig_test

import (
	"fmt"

	"github.com/casastorta/artifact-versioning/capabilities"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
)

var _ = Describe("TypeConfig", func() {
	var err error

	Describe("negative tests", func() {
		It("should fail with error if default detection mechanism passed among fallbacks", func() {
			_, err := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				fmt.Sprintf("%s, %s, %s",
					capabilities.Bamboo, config.DefaultBranchingDetection, capabilities.Git))
			Expect(err).To(Not(BeNil()))
		})

		It("should fail if invalid entry is passed in", func() {
			_, err := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				fmt.Sprintf("%s,%s, %s",
					capabilities.Git, "banana tree", capabilities.Bamboo))
			Expect(err).To(Not(BeNil()))
		})

		It("should fail if same entry is passed in multiple times", func() {
			_, err := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				fmt.Sprintf("%s,%s, %s",
					capabilities.Bamboo, capabilities.Git, capabilities.Bamboo))
			Expect(err).To(Not(BeNil()))
		})

		It("should fail on parsing invalid date", func() {
			_, err =
				typeconfig.Config("unparseable date", config.DefaultTimeZoneLocation,
					config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
					config.DefaultFallbackBranchingDetectionMechanisms)

			Expect(err).To(Not(BeNil()))
		})

		It("should fail on parsing invalid timezone information", func() {
			_, err =
				typeconfig.Config(config.DefaultTime, "Something irrelevant to timezones",
					config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
					config.DefaultFallbackBranchingDetectionMechanisms)

			Expect(err).To(Not(BeNil()))
		})
	})
})
