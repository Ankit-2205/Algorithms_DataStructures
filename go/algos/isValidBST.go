package algos

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsValidBST(root *TreeNode) bool {
	var inorder func(node *TreeNode) bool
	prev := -1 << 63
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !inorder(node.Left) {
			return false
		}
		if node.Val <= prev {
			return false
		}
		prev = node.Val
		return inorder(node.Right)
	}
	return inorder(root)
}
