// Package traversal реализует алгоритмы обхода графов и деревьев.
//
// Граф передаётся как список смежности: map[node][]neighbors.
// Узлы идентифицируются целыми числами.
package traversal

// BFS (Breadth-First Search) выполняет обход в ширину начиная с узла start.
// Возвращает map, где ключ — узел, значение — минимальное расстояние от start.
//
// # Теория:
//
//	BFS исследует граф слоями: сначала все соседи на расстоянии 1,
//	затем на расстоянии 2 и т.д. Это гарантирует кратчайший путь
//	в невзвешенном графе.
//
//	Очередь (FIFO) обеспечивает порядок обхода по уровням.
//	Узлы помечаются при добавлении в очередь, а не при извлечении —
//	иначе один узел может попасть в очередь несколько раз.
//
//	Сложность: O(V + E), где V — вершины, E — рёбра.
func BFS(start int, graph map[int][]int) map[int]int {
	levels := map[int]int{start: 0}
	queue := []int{start}

	for len(queue) > 0 {
		size := len(queue)
		for range size {
			node := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[node] {
				if _, seen := levels[neighbor]; !seen {
					levels[neighbor] = levels[node] + 1
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return levels
}

// BFSMultiSource выполняет BFS из нескольких стартовых узлов одновременно.
// Все узлы из starts считаются уровнем 0 и стартуют параллельно.
//
// # Теория:
//
//	Multi-source BFS — расширение обычного BFS для случаев,
//	когда источников несколько (например, несколько очагов "заражения").
//
//	Ключевая идея: положить все источники в очередь до начала обхода.
//	BFS сам обеспечит, что каждый узел будет достигнут от ближайшего
//	источника — потому что обход идёт слоями от всех источников сразу.
//
//	Делать отдельный BFS от каждого источника и брать минимум — неверно:
//	это O(S * (V+E)) и не учитывает одновременность распространения.
func BFSMultiSource(starts []int, graph map[int][]int) map[int]int {
	levels := make(map[int]int, len(starts))
	queue := make([]int, len(starts))

	for i, s := range starts {
		levels[s] = 0
		queue[i] = s
	}

	for len(queue) > 0 {
		size := len(queue)
		for range size {
			node := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[node] {
				if _, seen := levels[neighbor]; !seen {
					levels[neighbor] = levels[node] + 1
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return levels
}
