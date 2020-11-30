package test

import (
	"aoc2020/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array helpers", func() {

	Context("When converting an array of numbers", func() {
		It("Returns a correct representation", func() {
			in := []string{"12","5","-6","0","81293912"}
			out := []int{12,5,-6,0,81293912}

			r, err := helper.StrArrayToInt(in)
			Expect(err).To(BeNil())
			Expect(r).To(Equal(out))
		})

		It("Returns an error if an element is not an integer", func() {
			in := []string{"12","five"}
			_, err := helper.StrArrayToInt(in)
			Expect(err).NotTo(BeNil())

		})
		It("Handles an empty array", func() {
			in := []string{}
			r, err := helper.StrArrayToInt(in)
			Expect(err).To(BeNil())
			Expect(r).To(BeEmpty())
		})

	})


})