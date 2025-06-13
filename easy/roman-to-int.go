package easy

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	nums := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		nums = append(nums, m[s[i]])
	}

	total, tmp, pre := 0, 0, 0
	for _, num := range nums {
		if pre == 0 {
			pre = num
		}

		if num < pre {
			total += tmp
			tmp = num
			pre = num
		} else if num == pre {
			tmp += num
		} else if num > pre {
			total = total + num - tmp
			tmp, pre = 0, 0
		}
	}

	return total + tmp
}
