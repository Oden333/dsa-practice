package structs

import (
	"fmt"
	"math"
	"strconv"
)

var (
	Nan = math.MinInt
)

type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

func Fill(values []int) *TreeNode {
	var traverseAndFill func(root *TreeNode, value int) *TreeNode

	traverseAndFill = func(root *TreeNode, value int) *TreeNode {
		if value == Nan {
			return root
		}

		if root == nil {
			root = &TreeNode{Val: value}
			return root
		}

		if value >= root.Val {
			root.Right = traverseAndFill(root.Right, value)
		}

		if value <= root.Val {
			root.Left = traverseAndFill(root.Left, value)
		}

		return root
	}

	var root *TreeNode
	for _, val := range values {
		root = traverseAndFill(root, val)
	}

	return root
}

func MapValueToCount(root *TreeNode) map[int]int {
	counts := make(map[int]int)

	var mapping func(root *TreeNode)
	mapping = func(root *TreeNode) {
		if root == nil {
			return
		}
		counts[root.Val]++
		mapping(root.Left)
		mapping(root.Right)
	}
	mapping(root)
	return counts
}

func NodeMapping(root *TreeNode) {
	var depth int
	levelsMap := make(map[int][]string)

	var addNodes func(root *TreeNode)
	addNodes = func(root *TreeNode) {

		_, doesLevelExist := levelsMap[depth]
		if !doesLevelExist {
			level := make([]string, 0, int(math.Pow(2, float64(depth))))
			levelsMap[depth] = level
		}
		levelsMap[depth] = append(levelsMap[depth], nodeStringConversion(root))
		if root == nil || root.Val == Nan {
			return
		}

		depth++
		addNodes(root.Left)
		depth--
		depth++
		addNodes(root.Right)
		depth--
	}

	addNodes(root)

	initialDepth := len(levelsMap)
	for i := 0; i < len(levelsMap); i++ {
		fmt.Printf("%*s", initialDepth*2, "")
		for _, value := range levelsMap[i] {
			fmt.Printf("%s ", value)
		}
		fmt.Println()
		initialDepth--
	}
}

func nodeStringConversion(node *TreeNode) string {
	if node == nil {
		return "-"
	}

	return strconv.Itoa(node.Val)
}
