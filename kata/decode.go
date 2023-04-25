package kata

var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Decode(roman string) int {
	var sum = 0
	var prev rune

	for _, char := range []rune(roman) {
		value := romanNumerals[string(char)]
		if prev != 0 && romanNumerals[string(prev)] < value {
			// subtract the previous value if it's smaller than the current value
			sum -= romanNumerals[string(prev)]
			// then add the combined value of the previous and current values
			sum += value - romanNumerals[string(prev)]
		} else {
			// add the current value to the result
			sum += value
		}
		prev = char
	}
	return sum
}
