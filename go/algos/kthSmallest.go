package algos

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	var inorder func(node *TreeNode)
	result := []int{}

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	return result[k-1]
}

// Time Complexity: O(n)
// Space Complexity: O(n)

func KthSmallestNoSpace(root *TreeNode, k int) int {
	var inorder func(node *TreeNode, k *int) int

	inorder = func(node *TreeNode, k *int) int {
		if node == nil {
			return -1
		}

		left := inorder(node.Left, k)
		if left != -1 {
			return left
		}

		*k--
		if *k == 0 {
			return node.Val
		}

		return inorder(node.Right, k)
	}

	return inorder(root, &k)
}

// Time Complexity: O(n)
// Space Complexity: O(1)
