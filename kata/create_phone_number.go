package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var numbersString []string
	for _, number := range numbers {
		numbersString = append(numbersString, strconv.FormatUint(uint64(number), 10))
	}
	numberAsSingleString := strings.Join(numbersString, "")
	firstPart := numberAsSingleString[:3]
	secondPart := numberAsSingleString[3:6]
	thirdPart := numberAsSingleString[6:]
	return fmt.Sprintf("(%s) %s-%s", firstPart, secondPart, thirdPart)
}
