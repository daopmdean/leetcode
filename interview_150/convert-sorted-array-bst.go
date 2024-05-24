package interview_150

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	i := len(nums) / 2
	t := &TreeNode{
		Val: nums[i],
	}
	t.Left = sortedArrayToBST(nums[:i])
	t.Right = sortedArrayToBST(nums[i+1:])

	return t
}
