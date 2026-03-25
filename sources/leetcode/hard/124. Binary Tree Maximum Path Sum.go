package hard

import (
	. "dsa/structs"
)

func maxPathSum(root *TreeNode) int {
	var res int = -((1 << 63) - 1)

	var trav func(node *TreeNode) int
	trav = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lpath := trav(node.Left)
		rpath := trav(node.Right)

		var maxPath int
		if lpath > rpath {
			maxPath = lpath + node.Val
		} else {
			maxPath = rpath + node.Val
		}
		if maxPath < 0 {
			maxPath = 0
		}

		if tmp := lpath + node.Val + rpath; tmp > res {
			res = tmp
		}

		return maxPath
	}

	trav(root)
	return res
}
