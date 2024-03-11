package interview_150

func removeElement(nums []int, val int) int {
	result := make([]int, 0)

	for _, num := range nums {
		if num != val {
			result = append(result, num)
		}
	}

	copy(nums, result)

	return len(result)
}
