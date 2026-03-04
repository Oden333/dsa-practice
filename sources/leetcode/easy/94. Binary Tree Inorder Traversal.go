package easy

import (
	. "dsa/structs"
)

func inorderTraversal(root *TreeNode) []int {
	nums := make([]int, 0)
	counter := 0

	var inorderTraversal func(root *TreeNode)
	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		nums = append(nums, root.Val)
		counter++
		inorderTraversal(root.Right)
	}

	inorderTraversal(root)

	return nums[:counter]
}
