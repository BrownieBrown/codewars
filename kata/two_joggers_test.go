package kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func doTest(x, y uint16, exp [2]uint16) {
	var actual = NbrOfLaps(x, y)
	Expect(actual).To(Equal(exp))
}

var _ = Describe("Sample tests", func() {
	It("should handle basic cases", func() {
		doTest(5, 3, [2]uint16{3, 5})
		doTest(4, 6, [2]uint16{3, 2})
		doTest(5, 5, [2]uint16{1, 1})
	})
})
