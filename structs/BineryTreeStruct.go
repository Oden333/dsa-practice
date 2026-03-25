// Package structs предоставляет структуры данных для алгоритмических задач.
package structs

import (
	"fmt"
	"math"
	"strconv"
)

var (
	// Null — маркер отсутствующего узла в слайсе для инициализации дерева.
	// Используем MinInt, т.к. в задачах LeetCode значения обычно в диапазоне [-100, 100].
	Null = math.MinInt

	// EmptyVal - маркер отсутствующего узла в дереве для принта
	EmptyVal = "-"
)

// TreeNode — узел бинарного дерева.
type TreeNode struct {
	Left, Right *TreeNode
	Val         int
}

// Fill строит BST из слайса значений путём последовательной вставки.
// Fill использует логику BST: value >= root.Val → вправо, иначе влево.
func Fill(values []int) *TreeNode {
	var traverseAndFill func(node *TreeNode, value int) *TreeNode
	traverseAndFill = func(node *TreeNode, value int) *TreeNode {
		// Пропускаем маркеры nil-узлов
		if value == Null {
			return node
		}

		// Создаём новый узел, если дошли до nil
		if node == nil {
			return &TreeNode{Val: value}
		}

		if value >= node.Val {
			node.Right = traverseAndFill(node.Right, value)
		} else if value <= node.Val {
			node.Left = traverseAndFill(node.Left, value)
		}

		return node
	}

	var root *TreeNode
	for _, val := range values {
		root = traverseAndFill(root, val)
	}
	return root
}

// BuildTree создаёт дерево из level-order слайса
/*
Пример: [1,2,3,nil,5] →
	  1
	 / \
	2   3
	 \
	  5
*/
func BuildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == Null {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	idx := 1

	for len(queue) > 0 && idx < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if idx < len(vals) && vals[idx] != Null {
			node.Left = &TreeNode{Val: vals[idx]}
			queue = append(queue, node.Left)
		}
		idx++

		if idx < len(vals) && vals[idx] != Null {
			node.Right = &TreeNode{Val: vals[idx]}
			queue = append(queue, node.Right)
		}
		idx++
	}

	return root
}

// MapValueToCount считает частоту каждого значения в дереве (pre-order обход).
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

// NodeMapping печатает дерево по уровням (для дебага)
func NodeMapping(root *TreeNode) {
	levelsMap := make(map[int][]string)

	var addNodes func(root *TreeNode, depth int)
	addNodes = func(root *TreeNode, depth int) {
		// Инициализируем уровень, если ещё не создан
		if _, exists := levelsMap[depth]; !exists {
			levelsMap[depth] = make([]string, 0)
		}
		levelsMap[depth] = append(levelsMap[depth], nodeStringConversion(root))

		if root == nil {
			return
		}

		addNodes(root.Left, depth+1)
		addNodes(root.Right, depth+1)
	}

	addNodes(root, 0)

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

// nodeStringConversion конвертирует узел в строку для печати.
func nodeStringConversion(node *TreeNode) string {
	if node == nil {
		return EmptyVal
	}
	return strconv.Itoa(node.Val)
}
