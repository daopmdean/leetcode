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

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	node := root
	for node != nil {
		if node.Val > val && node.Left != nil {
			node = node.Left
		} else if node.Val > val && node.Left == nil {
			node.Left = &TreeNode{
				Val: val,
			}
			break
		} else if node.Val < val && node.Right != nil {
			node = node.Right
		} else if node.Val < val && node.Right == nil {
			node.Right = &TreeNode{
				Val: val,
			}
			break
		}
	}

	return root
}
