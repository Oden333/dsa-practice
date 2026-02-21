package medium

import (
	. "leetcode/structs"
	"slices"
)

func KthLargestLevelSum(root *TreeNode, k int) int64 {

	level := 0
	sumSlice := make([]int, 0)
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		if len(sumSlice) <= level {
			sumSlice = append(sumSlice, 0)
		}
		sumSlice[level] += root.Val

		level++
		traverse(root.Left)
		traverse(root.Right)
		level--
	}

	traverse(root)
	if len(sumSlice) < k {
		return -1
	}

	slices.Sort(sumSlice)
	return int64(sumSlice[len(sumSlice)-k])
}
