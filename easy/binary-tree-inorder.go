package easy
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := []int{}
    traverse(root, &result)
    return result
}

func traverse(root *TreeNode, sl *[]int) {
    if root != nil {
        traverse(root.Left, sl)
        *sl = append(*sl, root.Val)
        traverse(root.Right, sl)
    }
}
