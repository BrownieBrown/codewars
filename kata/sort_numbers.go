package kata

func SortNumbers(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)

	for i := 0; i < len(sortedNumbers)-1; i++ {
		for j := 0; j < len(sortedNumbers)-i-1; j++ {
			if sortedNumbers[j] > sortedNumbers[j+1] {
				lesserNumber := sortedNumbers[j+1]
				greaterNumber := sortedNumbers[j]
				sortedNumbers[j] = lesserNumber
				sortedNumbers[j+1] = greaterNumber
			}
		}
	}

	return sortedNumbers
}
