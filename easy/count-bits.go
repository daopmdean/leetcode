package easy

func countBits(n int) []int {
	result := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		result = append(result, solve(i))
	}
	return result
}

func solve(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num = num / 2
	}
	return count
}
