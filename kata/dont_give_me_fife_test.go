package kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DontGiveMeFive", func() {
	It("returns count 8", func() {
		Expect(DontGiveMeFive(1, 9)).To(Equal(8))
	})

	It("returns count 12", func() {
		Expect(DontGiveMeFive(4, 17)).To(Equal(12))
	})
})
