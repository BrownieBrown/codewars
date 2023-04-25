package kata

func isEven(str string) bool {
	if len(str)%2 == 0 {
		return true
	}

	return false
}

func Solution(str string) []string {
	var stringArray []string

	if !isEven(str) {
		str = str + "_"
	}

	for i := 0; i < len(str); i += 2 {
		stringArray = append(stringArray, str[i:i+2])
	}

	return stringArray
}
