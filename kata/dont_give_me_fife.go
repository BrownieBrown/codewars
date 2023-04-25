package kata

import (
	"strconv"
	"strings"
)

func DontGiveMeFive(start int, end int) int {
	count := 0

	for i := start; i <= end; i++ {
		if !HasDigit5(i) {
			count++
		}
	}
	return count
}

func HasDigit5(n int) bool {
	s := strconv.Itoa(n)
	return strings.Contains(s, "5")
}
