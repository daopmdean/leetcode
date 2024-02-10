package interview_150

import "fmt"

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}

	start, end := 0, 0
	for i, num := range nums {
		if i == 0 {
			start = num
			end = num
			continue
		}

		if end == num-1 {
			end = num
			continue
		}

		if start == end {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, end))
		}
		start = num
		end = num

	}

	if start == end {
		result = append(result, fmt.Sprintf("%d", start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, end))
	}

	return result
}
