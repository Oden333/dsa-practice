package traversal

import "dsa/structs"

// BFSGraph выполняет обход в ширину по графу из structs.GraphNode.
// Возвращает map: Val узла → расстояние от start.
//
// # Теория:
//
//	GraphNode хранит явные указатели на соседей, поэтому отдельный
//	список смежности не нужен — соседи уже встроены в структуру.
//	Роль seen играет map по указателю: два узла с одинаковым Val
//	всё равно различаются, если это разные объекты.
func BFSGraph(start *structs.GraphNode) map[int]int {
	if start == nil {
		return nil
	}

	levels := map[*structs.GraphNode]int{start: 0}
	result := map[int]int{start.Val: 0}
	queue := []*structs.GraphNode{start}

	for len(queue) > 0 {
		size := len(queue)
		for range size {
			node := queue[0]
			queue = queue[1:]
			for _, neighbor := range node.Neighbors {
				if _, seen := levels[neighbor]; !seen {
					levels[neighbor] = levels[node] + 1
					result[neighbor.Val] = levels[neighbor]
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return result
}

// BFSTree выполняет обход бинарного дерева по уровням (level-order traversal).
// Возвращает срез уровней: каждый элемент — значения узлов одного уровня.
//
// # Теория:
//
//	Level-order — это BFS на дереве. В отличие от графа, seen не нужен:
//	в дереве нет циклов и каждый узел имеет ровно одного родителя.
//
//	Возвращаемый [][]int удобен для задач типа "среднее по уровням",
//	"правая сторона дерева", "зигзаг-обход" и т.д.
func BFSTree(root *structs.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	queue := []*structs.TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for range size {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
