package kata

func Pyramid(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			result[i][j] = 1
		}
	}
	return result
}
