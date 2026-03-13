package easy

import (
	. "dsa/structs"
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var trav func(node *TreeNode, path []string)
	trav = func(node *TreeNode, path []string) {
		if node == nil {
			return
		}

		if node.Val != -101 { // nil
			path = append(path, strconv.Itoa(node.Val))
		}

		if node.Left == nil && node.Right == nil {
			res = append(res, strings.Join(path, "->"))
			return
		}

		if node.Left != nil {
			trav(node.Left, path)
		}

		if node.Right != nil {
			trav(node.Right, path)
		}
	}

	trav(root, make([]string, 0))

	return res
}
