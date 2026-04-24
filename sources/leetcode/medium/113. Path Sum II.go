package medium

import (
	. "dsa/structs"
	"slices"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int = make([][]int, 0)

	var trav func(node *TreeNode, path []int, sum int)
	trav = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		sum -= node.Val
		if node.Left == nil && node.Right == nil &&
			sum == 0 {
			res = append(res, slices.Clone(path))
		}

		if node.Left != nil {
			trav(node.Left, path, sum)
		}

		if node.Right != nil {
			trav(node.Right, path, sum)
		}

	}

	trav(root, make([]int, 0), targetSum)

	return res
}
