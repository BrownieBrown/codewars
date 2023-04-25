package kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Tests", func() {
	It("Testing for 0", func() { Expect(Pyramid(0)).To(Equal([][]int{})) })
	It("Testing for 1", func() { Expect(Pyramid(1)).To(Equal([][]int{[]int{1}})) })
	It("Testing for 2", func() { Expect(Pyramid(2)).To(Equal([][]int{[]int{1}, []int{1, 1}})) })
	It("Testing for 3", func() { Expect(Pyramid(3)).To(Equal([][]int{[]int{1}, []int{1, 1}, []int{1, 1, 1}})) })
})
