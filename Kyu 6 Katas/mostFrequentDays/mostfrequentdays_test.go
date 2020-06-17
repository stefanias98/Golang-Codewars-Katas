package mostfrequentdays_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "codewarrior/mostfrequentdays"
)

func test(input int, expected ...string) {
	Expect(MostFrequentDays(input)).To(Equal(expected))
}

var _ = Describe("Sample Tests", func() {
	It("Sample Tests", func() {
		test(2427, "Friday")
		test(2185, "Saturday")
		test(1167, "Sunday")
		test(1770, "Monday")
		test(1785, "Saturday")
		test(1984, "Monday", "Sunday")
		test(3076, "Saturday", "Sunday")
	})
})
