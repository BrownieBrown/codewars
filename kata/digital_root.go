package kata

import "strconv"

func DigitalRoot(n int) int {
	stringNumber := strconv.Itoa(n)
	sum := 0
	sum = buildSum(stringNumber)

	for sum >= 10 {
		stringNumber = strconv.Itoa(sum)
		sum = 0
		sum = buildSum(stringNumber)
	}

	return sum
}

func buildSum(s string) int {
	sum := 0
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}

	return sum
}
