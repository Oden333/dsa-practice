package traversal

// DFS (Depth-First Search) выполняет рекурсивный обход в глубину начиная с узла start.
// Возвращает порядок посещения узлов.
//
// # Теория:
//
//	DFS уходит как можно глубже по каждой ветке перед тем, как
//	вернуться и исследовать соседей. В отличие от BFS, не гарантирует
//	кратчайший путь — зато полезен для поиска компонент связности,
//	топологической сортировки, обнаружения циклов.
//
//	Рекурсивная реализация естественно отражает стек вызовов.
//	Итеративный аналог использует явный стек ([]int), но обходит
//	соседей в обратном порядке из-за LIFO-природы стека.
//
//	Сложность: O(V + E).
func DFS(start int, graph map[int][]int) []int {
	seen := map[int]bool{}
	order := []int{}

	var dfs func(node int)
	dfs = func(node int) {
		seen[node] = true
		order = append(order, node)
		for _, neighbor := range graph[node] {
			if !seen[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
	return order
}

// DFSMultiSource запускает DFS из нескольких стартовых узлов.
// Источники обходятся последовательно: следующий стартует только если
// все узлы, достижимые из предыдущего, уже посещены.
//
// # Теория:
//
//	В отличие от multi-source BFS, где все источники стартуют одновременно,
//	multi-source DFS обходит компоненты последовательно. Это стандартный
//	паттерн для подсчёта компонент связности или поиска "островов":
//	итерируемся по всем узлам графа, запуская DFS только из непосещённых.
func DFSMultiSource(starts []int, graph map[int][]int) []int {
	seen := map[int]bool{}
	order := []int{}

	var dfs func(node int)
	dfs = func(node int) {
		seen[node] = true
		order = append(order, node)
		for _, neighbor := range graph[node] {
			if !seen[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for _, s := range starts {
		if !seen[s] {
			dfs(s)
		}
	}

	return order
}
