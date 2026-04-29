package traversal

import "dsa/structs"

// DFSGraph выполняет обход в глубину по графу из structs.GraphNode.
// Возвращает порядок посещения узлов (по Val).
func DFSGraph(start *structs.GraphNode) []int {
	if start == nil {
		return nil
	}

	seen := map[*structs.GraphNode]bool{}
	order := []int{}

	var dfs func(node *structs.GraphNode)
	dfs = func(node *structs.GraphNode) {
		seen[node] = true
		order = append(order, node.Val)
		for _, neighbor := range node.Neighbors {
			if !seen[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
	return order
}

// DFSPreOrder обходит бинарное дерево в порядке: корень → левое → правое.
//
// # Теория:
//
//	Pre-order удобен когда нужно обработать узел до его потомков:
//	копирование дерева, сериализация, вычисление префиксных выражений.
func DFSPreOrder(root *structs.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}

	var dfs func(node *structs.TreeNode)
	dfs = func(node *structs.TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return result
}

// DFSInOrder обходит бинарное дерево в порядке: левое → корень → правое.
//
// # Теория:
//
//	In-order на BST даёт узлы в отсортированном порядке.
//	Именно поэтому in-order — стандартный способ проверить, является ли
//	дерево корректным BST.
func DFSInOrder(root *structs.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}

	var dfs func(node *structs.TreeNode)
	dfs = func(node *structs.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}

	dfs(root)
	return result
}

// DFSPostOrder обходит бинарное дерево в порядке: левое → правое → корень.
//
// # Теория:
//
//	Post-order нужен когда узел обрабатывается после всех потомков:
//	удаление дерева, вычисление высоты/размера поддеревьев,
//	вычисление постфиксных выражений.
func DFSPostOrder(root *structs.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}

	var dfs func(node *structs.TreeNode)
	dfs = func(node *structs.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}

	dfs(root)
	return result
}
