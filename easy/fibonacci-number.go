package easy

// 0, 1, 1, 2, 3, 5, 8
func fib(n int) int {
	a, b := 0, 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			continue
		}
		if i == 1 {
			b = 1
			continue
		}
		b += a
		a = b - a
	}
	return b
}
