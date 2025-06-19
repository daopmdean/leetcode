package easy

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := range numRows {
		result[i] = make([]int, i+1)
		for j := range result[i] {
			if j == 0 {
				result[i][j] = 1
				continue
			} else if j == len(result[i])-1 {
				result[i][j] = 1
				continue
			}

			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
