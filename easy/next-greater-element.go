package easy

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		n := nextGreater(nums2, num)
		result = append(result, n)
	}
	return result
}

func nextGreater(nums []int, x int) int {
	for i, num := range nums {
		if num == x {
			for j := i; j < len(nums); j++ {
				if nums[j] > x {
					return nums[j]
				}
			}
			break
		}
	}
	return -1
}
