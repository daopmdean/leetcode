package interview_150

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, num := range nums {
		if v, ok := m[num]; ok {
			if i-v <= k {
				return true
			}
		}

		m[num] = i
	}

	return false
}
