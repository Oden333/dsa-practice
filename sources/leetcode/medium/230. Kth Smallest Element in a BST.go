package medium

import (
	. "dsa/structs"
)

func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var result int

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || result != 0 {
			return
		}

		inorder(node.Left)

		count++
		if count == k {
			result = node.Val
		}

		inorder(node.Right)
	}

	inorder(root)
	return result
}
