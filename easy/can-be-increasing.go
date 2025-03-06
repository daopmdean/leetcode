package easy

func canBeIncreasing(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	for i := 0; i < len(nums); i++ {
		newNums := make([]int, 0, len(nums)-1)
		newNums = append(newNums, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		if strictlyIncreased(newNums) {
			return true
		}
	}

	return false
}

func strictlyIncreased(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}
