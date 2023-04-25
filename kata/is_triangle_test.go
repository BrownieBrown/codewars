package kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsTriangle", func() {
	Context("with valid sides", func() {
		It("returns true", func() {
			Expect(IsTriangle(3, 4, 5)).To(BeTrue())
		})
	})

	Context("with valid sides", func() {
		It("returns false", func() {
			Expect(IsTriangle(5, 1, 2)).To(BeFalse())
		})
	})

	Context("with invalid sides", func() {
		It("returns false", func() {
			Expect(IsTriangle(0, 4, 5)).To(BeFalse())
		})
	})

	Context("with invalid sides", func() {
		It("returns false", func() {
			Expect(IsTriangle(-1, 4, 5)).To(BeFalse())
		})
	})
})
