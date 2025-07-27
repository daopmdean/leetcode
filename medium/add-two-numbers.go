package medium

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	l1Cur, l2Cur, resultCur := l1, l2, result
	hold := 0

	for l1Cur != nil && l2Cur != nil {
		val := l1Cur.Val + l2Cur.Val + hold
		if val <= 9 {
			resultCur.Val = val
			hold = 0
		} else {
			val = val % 10
			hold = 1
			resultCur.Val = val
		}
		l1Cur = l1Cur.Next
		l2Cur = l2Cur.Next
		if l1Cur != nil || l2Cur != nil {
			resultCur.Next = &ListNode{}
			resultCur = resultCur.Next
		}
	}

	for l1Cur != nil {
		val := l1Cur.Val + hold
		if val <= 9 {
			resultCur.Val = val
			hold = 0
		} else {
			val = val % 10
			resultCur.Val = val
			hold = 1
		}
		l1Cur = l1Cur.Next
		if l1Cur != nil {
			resultCur.Next = &ListNode{}
			resultCur = resultCur.Next
		}
	}

	for l2Cur != nil {
		val := l2Cur.Val + hold
		if val <= 9 {
			resultCur.Val = val
			hold = 0
		} else {
			val = val % 10
			resultCur.Val = val
			hold = 1
		}
		l2Cur = l2Cur.Next
		if l2Cur != nil {
			resultCur.Next = &ListNode{}
			resultCur = resultCur.Next
		}
	}

	if hold == 1 {
		resultCur.Next = &ListNode{}
		resultCur = resultCur.Next
		resultCur.Val = 1
	}

	return result
}
