package easy

func binarySearch(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start < end {
		checkIdx := (start + end) / 2
		if nums[checkIdx] == target {
			return checkIdx
		} else if nums[checkIdx] < target {
			start = checkIdx + 1
		} else {
			end = checkIdx - 1
		}
	}

	if start == end && nums[start] == target {
		return start
	}

	return -1
}
