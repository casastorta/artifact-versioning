package uniquestringlist_test

import (
	"github.com/casastorta/artifact-versioning/types/uniquestringlist"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UniqueStringList", func() {
	var InitialUniqueStringList uniquestringlist.UniqueStringList

	BeforeEach(func() {
		InitialUniqueStringList = uniquestringlist.UniqueStringList{"a", "something here", "b", "contains", "d"}
	})

	Describe("Contains method", func() {
		It("should detect item 'b'", func() {
			Expect(InitialUniqueStringList.Contains("b")).To(BeTrue())
		})

		It("should detect random string is not in the list", func() {
			Expect(InitialUniqueStringList.Contains("something else")).To(BeFalse())
		})
	})

	Describe("Append method", func() {
		It("should add item 'anything else' successfully", func() {
			err := InitialUniqueStringList.Append("anything else")

			Expect(err).To(BeNil())
			Expect(InitialUniqueStringList.Contains("anything else")).To(BeTrue())
		})

		It("should fail on adding 'contains' to the list", func() {
			err := InitialUniqueStringList.Append("contains")

			Expect(err).To(Not(BeNil()))
		})
	})

	Describe("FromString method", func() {
		It("should convert string separated by comma into the type", func() {
			var (
				sp = "a, b, c"
				us uniquestringlist.UniqueStringList
			)
			err := us.FromString(sp, ",")

			Expect(err).To(BeNil())
			Expect(us).To(Equal(uniquestringlist.UniqueStringList{"a", "b", "c"}))
		})

		It("should convert string separated by pipe into the type", func() {
			var (
				sp = "a | b | c"
				us uniquestringlist.UniqueStringList
			)
			err := us.FromString(sp, "|")

			Expect(err).To(BeNil())
			Expect(us).To(Equal(uniquestringlist.UniqueStringList{"a", "b", "c"}))
		})

		It("should fail to create type from the string if elements repeat", func() {
			var (
				sp = "c, d, e, d, f"
				us uniquestringlist.UniqueStringList
			)
			err := us.FromString(sp, ",")

			Expect(err).To(Not(BeNil()))
		})

		It("should be able to create type with empty element", func() {
			var (
				sp = "c,,d, e, h"
				us uniquestringlist.UniqueStringList
			)
			err := us.FromString(sp, ",")

			Expect(err).To(BeNil())
			Expect(us).To(Equal(uniquestringlist.UniqueStringList{"c", "", "d", "e", "h"}))
		})

		It("should obey with separator definition completely", func() {
			var (
				sp = "c,,d, e, h"
				us uniquestringlist.UniqueStringList
			)
			err := us.FromString(sp, ", ")

			Expect(err).To(BeNil())
			Expect(us).To(Equal(uniquestringlist.UniqueStringList{"c,,d", "e", "h"}))
		})
	})
})
