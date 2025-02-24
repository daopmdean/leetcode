package interview_150

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1 -> 2 -> 4
// 1 -> 3 -> 4
// --> 1 -> 1 -> 2 -> 3 -> 4 -> 4
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var node *ListNode

	for list1 != nil || list2 != nil {
		if list1 == nil {
			return joinListToNode(node, list2)
		}
		if list2 == nil {
			return joinListToNode(node, list1)
		}

		if list1.Val <= list2.Val {
			node = appendToList(node, list1.Val)
			list1 = list1.Next
		} else {
			node = appendToList(node, list2.Val)
			list2 = list2.Next
		}
	}

	return node
}

func joinListToNode(node *ListNode, list *ListNode) *ListNode {
	for list != nil {
		node = appendToList(node, list.Val)
		list = list.Next
	}

	return node
}

func appendToList(node *ListNode, val int) *ListNode {
	if node == nil {
		return &ListNode{Val: val}
	}

	iter := node
	for iter.Next != nil {
		iter = iter.Next
	}

	iter.Next = &ListNode{Val: val}
	return node
}
