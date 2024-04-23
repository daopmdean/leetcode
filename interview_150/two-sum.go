package interview_150

func twoSum(nums []int, target int) []int {
	mIndex := map[int]int{}
	mNeed := map[int]bool{}

	for i, v := range nums {
		if mNeed[v] {
			return []int{mIndex[target-v], i}
		}
		mIndex[v] = i
		mNeed[target-v] = true
	}

	return []int{}
}
