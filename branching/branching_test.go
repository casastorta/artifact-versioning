package branching

import (
	"github.com/casastorta/artifact-versioning/capabilities"
	"github.com/casastorta/artifact-versioning/config"
	"github.com/casastorta/artifact-versioning/types/typeconfig"
	"github.com/casastorta/artifact-versioning/types/uniquestringlist"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Branching", func() {
	Describe("Branching constants and defaults", func() {
		It("should contain default variable to allow fallback detection", func() {
			Expect(AllowFallbackDetection).To(BeTrue())
		})

		It("should contain all the branching methods", func() {
			detectionMethods := capabilities.BranchingDetectionMethods()

			Expect(detectionMethods).To(ContainElement(capabilities.Git))
			Expect(detectionMethods).To(ContainElement(capabilities.Bamboo))
			Expect(detectionMethods).To(ContainElement(capabilities.Jenkins))
		})

		It("should contain default order of detection fallbacks", func() {
			detectionOrders := GetDetectionOrder()

			expectedOrder := uniquestringlist.UniqueStringList{
				capabilities.Git, capabilities.Bamboo, capabilities.Jenkins,
			}
			Expect(detectionOrders).To(Equal(expectedOrder))

			notExpectedOrder := uniquestringlist.UniqueStringList{
				capabilities.Jenkins, capabilities.Git, capabilities.Bamboo,
			}
			Expect(detectionOrders).To(Not(Equal(notExpectedOrder)))
		})
	})

	Describe("Settings/variables should be modifiable", func() {
		It("should allow to modify AllowFallbackDetection", func() {
			AllowFallbackDetection = false
			Expect(AllowFallbackDetection).To(BeFalse())
		})

		It("should allow to modify detection order", func() {
			defaultOrder := GetDetectionOrder()
			newOrder := uniquestringlist.UniqueStringList{capabilities.Jenkins, capabilities.Git}

			err := SetDetectionOrder(newOrder...)
			Expect(err).To(BeNil())
			Expect(GetDetectionOrder()).To(Equal(newOrder))

			err2 := SetDetectionOrder(defaultOrder...)
			Expect(err2).To(BeNil())
			Expect(GetDetectionOrder()).To(Equal(defaultOrder))
		})

		It("should not allow entries which are not detection mechanisms in detectionOrder", func() {
			notExpectedOrder := []string{capabilities.Git, "banana tree"}
			err := SetDetectionOrder(notExpectedOrder...)

			Expect(err).To(Not(BeNil()))
		})

		It("should not allow duplicate items in detectionOrder", func() {
			newOrder := []string{capabilities.Jenkins, capabilities.Git, capabilities.Bamboo, capabilities.Git}
			err := SetDetectionOrder(newOrder...)

			Expect(err).To(Not(BeNil()))
		})
	})

	Describe("GetYearWeek", func() {
		var (
			configInstance *typeconfig.TypeConfig
			YearWeek       string
		)

		BeforeEach(func() {
			configInstance, _ = config.DefaultConfig()
			YearWeek = GetYearWeek(configInstance)
		})

		Describe("With default config instance", func() {
			It("should return year and week", func() {
				Ω(YearWeek).Should(MatchRegexp("[1-9][0-9][0-9][0-9]\\.[0-5][0-9]"))
			})
		})
	})
})
