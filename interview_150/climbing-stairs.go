package interview_150

/*
1 -> 1
2 -> 2
3 -> 3
4 -> 5
5 -> 8

1 1 1 1
1 2 1
1 1 2
2 1 1
2 2

1 1 1 1 1
1 2 1 1
1 1 2 1
1 1 1 2
2 1 1 1
2 2 1
2 1 2
1 2 2
*/

func climbStairs(n int) int {
	sl := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			sl[i] = 1
		}
		if i == 1 {
			sl[i] = 2
		}
		if i > 1 {
			sl[i] = sl[i-1] + sl[i-2]
		}
	}

	return sl[n-1]
}
