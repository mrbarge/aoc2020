package test

import (
	"aoc2020/helper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("File helpers", func() {

	Context("ReadLines", func() {
		It("Behaves correctly", func() {
			in := `abcd
efgh
ijkl`
			ir := strings.NewReader(in)
			out := []string{"abcd", "efgh", "ijkl"}
			res, err := helper.ReadLines(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
		It("Ignores empty lines", func() {
			in := `abcd

efgh`
			ir := strings.NewReader(in)
			out := []string{"abcd", "efgh"}
			res, err := helper.ReadLines(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})

	})

	Context("ReadLinesAsInt", func() {
		It("Behaves correctly", func() {
			in := `123
-456
0
3322351`
			ir := strings.NewReader(in)
			out := []int{123,-456,0,3322351}
			res, err := helper.ReadLinesAsInt(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
		It("Ignores empty lines", func() {
			in := `123

0`
			ir := strings.NewReader(in)
			out := []int{123,0}
			res, err := helper.ReadLinesAsInt(ir)
			Expect(err).To(BeNil())
			Expect(res).To(Equal(out))
		})
	})

})
