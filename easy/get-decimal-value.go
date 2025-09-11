/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	result, l := 0, len(nums)
	for i, num := range nums {
		if num == 0 {
			continue
		}
		sum := int(math.Pow(2, float64(l-i-1)))
		result += sum
	}
	return result
}
