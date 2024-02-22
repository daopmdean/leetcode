package interview_150

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return 1
	}

	start, end := 0, len(nums)-1
	for {
		if start+1 == end {
			if nums[start] >= target {
				return start
			}
			if nums[end] >= target {
				return end
			}
			return end + 1
		}

		i := (start + end) / 2
		if nums[i] < target {
			start = i
		} else {
			end = i
		}
	}
}
