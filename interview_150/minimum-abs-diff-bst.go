package interview_150

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {

	node := root
	min := getMinFromNode(node)

	for node.Left != nil {
		getAbs(node.Val, node.Left.Val)
	}

	for node.Right != nil {

	}
	return min
}

func getMin(node *TreeNode, min int) int {
	return getMinFromNode(node)
}

func getMinFromNode(node *TreeNode) int {
	l, r := 0, 0
	if node.Left != nil {
		l = getAbs(node.Val, node.Left.Val)
	}
	if node.Right != nil {
		r = getAbs(node.Val, node.Right.Val)
	}

	if l == 0 {
		return r
	}

	if r == 0 {
		return l
	}

	if l < r {
		return l
	}

	return r
}

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
