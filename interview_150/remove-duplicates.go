package interview_150

func removeDuplicates(nums []int) int {
	existed := map[int]bool{}
	result := []int{}
	for _, num := range nums {
		if existed[num] {
			continue
		}

		existed[num] = true
		result = append(result, num)
	}

	copy(nums, result)

	return len(existed)
}
