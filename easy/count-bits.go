package easy

func countBits(n int) []int {
	result := make([]int, n+1)
	// result := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		result[i] = solve(i, result)
		// result = append(result, solve(i))
	}
	return result
}

// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101
// 6 --> 110
// 7 --> 111
// 8 --> 1000
// 9 --> 1001
// 10 -> 1010
// 11 -> 1011
func solve(num int, result []int) int {
	count := 0
	for num > 0 {
		if num%2 == 1 {
			count++
		}
		num = num / 2
		if result[num] != 0 {
			count += result[num]
			break
		}
	}
	return count
}
