package medium

/**
 * Definition for a binary tree node.
 *
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
	}
	addValue(root, val)
	return root
}

func addValue(node *TreeNode, val int) {
	if node.Val > val {
		if node.Left != nil {
			addValue(node.Left, val)
		} else {
			node.Left = &TreeNode{
				Val: val,
			}
		}
	}

	if node.Val < val {
		if node.Right != nil {
			addValue(node.Right, val)
		} else {
			node.Right = &TreeNode{
				Val: val,
			}
		}
	}
}
