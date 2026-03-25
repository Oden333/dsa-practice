package easy

import (
	. "dsa/structs"
	"slices"
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	var trav func(node *TreeNode, path []int)
	trav = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}

		curPath := append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, pathConvert(curPath))
			return
		}

		trav(node.Left, (slices.Clone(curPath)))
		trav(node.Right, (slices.Clone(curPath)))
	}

	trav(root, []int{})

	return res
}

func pathConvert(res []int) string {
	b := strings.Builder{}
	for i, v := range res {
		if i != 0 {
			b.WriteString("->")
		}
		b.WriteString(strconv.Itoa(v))
	}

	return b.String()
}
