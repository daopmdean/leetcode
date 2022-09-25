package easy

func twoSum(nums []int, target int) []int {
	haveMap := map[int]int{}
	needMap := map[int]bool{}
	for i, v := range nums {
		if needMap[v] {
			return []int{haveMap[target-v], i}
		}
		haveMap[v] = i
		needMap[target-v] = true
	}
	return []int{}
}
