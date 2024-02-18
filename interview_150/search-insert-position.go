package interview_150

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if num < target {
			continue
		}

		return i
	}

	return len(nums)
}
