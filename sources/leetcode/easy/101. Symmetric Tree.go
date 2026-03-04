package easy

import . "dsa/structs"

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var symmetric func(Lroot *TreeNode, Rroot *TreeNode) bool
	symmetric = func(L *TreeNode, R *TreeNode) bool {

		if L == nil && R == nil {
			return true
		}
		if xor((L == nil), (R == nil)) || (L.Val != R.Val) {
			return false
		}

		return symmetric(L.Left, R.Right) && symmetric(L.Right, R.Left)

	}
	return symmetric(root.Left, root.Right)
}

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}
