package typeconfig_test

import (
	"fmt"

	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/types/uniquestringlist"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
)

var _ = Describe("TypeConfig", func() {
	var (
		configInstance *typeconfig.TypeConfig
		err            error
	)

	Describe("BranchingFallbackDetectionMechanisms", func() {
		It("should return empty list on default instance", func() {
			c, err := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				config.DefaultFallbackBranchingDetectionMechanisms)
			Expect(err).To(BeNil())

			b := c.BranchingFallbackDetectionMechanisms()
			Expect(b).To(BeEmpty())
		})

		It("should return populated list when set", func() {
			var ca uniquestringlist.UniqueStringList
			err1 := ca.Append(capabilities.Bamboo, capabilities.Git)
			Expect(err1).To(BeNil())

			c, err2 := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				fmt.Sprintf("%s, %s", capabilities.Bamboo, capabilities.Git))
			Expect(err2).To(BeNil())

			b := c.BranchingFallbackDetectionMechanisms()
			Expect(b).To(Equal(ca))
		})

		It("should skip without error if one or more items in the input list are empty", func() {
			var ca uniquestringlist.UniqueStringList
			erra := ca.Append(capabilities.Bamboo, capabilities.Git)
			Expect(erra).To(BeNil())

			c, err := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				fmt.Sprintf("%s,%s, %s, %s",
					"", capabilities.Bamboo, "", capabilities.Git))
			Expect(err).To(BeNil())

			b := c.BranchingFallbackDetectionMechanisms()
			Expect(b).To(Equal(ca))
		})

		It("should return empty set if only empty element(s) are in the input", func() {
			c, err := typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
				config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
				fmt.Sprintf("%s,%s, %s,,,",
					"", "", ""))
			Expect(err).To(BeNil())

			b := c.BranchingFallbackDetectionMechanisms()
			Expect(b).To(BeEmpty())
		})
	})

	Describe("Config instancing with Europe/Berlin Timezne", func() {
		It("should instance successfully", func() {
			alternativeTimeZone := "Europe/Berlin"
			configInstance, err =
				typeconfig.Config(config.DefaultTime, alternativeTimeZone,
					config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
					config.DefaultFallbackBranchingDetectionMechanisms)

			Expect(err).To(BeNil())
			Expect(configInstance.TimeZoneLocation.String()).To(Equal(alternativeTimeZone))
		})
	})

	Describe("Config instancing with the safe Timezone", func() {
		It("should instance successfully", func() {
			configInstance, err =
				typeconfig.Config(config.DefaultTime, config.DefaultTimeZoneLocation,
					config.DefaultBranchingDetection, config.DefaultBranchingFallBackDetection,
					config.DefaultFallbackBranchingDetectionMechanisms)

			Expect(err).To(BeNil())
			Expect(configInstance.TimeZoneLocation.String()).To(Equal(config.DefaultTimeZoneLocation))
		})
	})
})
