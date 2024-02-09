package interview_150

import "strconv"

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	existed := map[int]bool{}

	for {
		s := strconv.Itoa(n)
		n = getTotal(s)
		if n == 1 {
			return true
		}

		if existed[n] {
			return false
		}

		existed[n] = true
	}

	return false
}

func getTotal(s string) int {
	total := 0
	for _, r := range s {
		i, _ := strconv.Atoi(string(r))
		total += i * i
	}
	return total
}
