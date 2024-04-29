package interview_150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	step, jump := head, head

	for jump != nil && jump.Next != nil {

		step = step.Next
		jump = jump.Next.Next

		if step == jump {
			return true
		}
	}

	return false
}
