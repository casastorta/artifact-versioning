package config_test

import (
	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		configInstance *typeconfig.TypeConfig
		err            error
	)

	Describe("Default config instancing", func() {
		It("should have default location and branching", func() {
			configInstance, err = config.DefaultConfig()

			Expect(err).To(BeNil())
			Expect(configInstance.TimeZoneLocation.String()).To(Equal(config.DefaultTimeZoneLocation))
			Expect(configInstance.BranchingDetection).To(Equal(config.DefaultBranchingDetection))
		})
	})
})
