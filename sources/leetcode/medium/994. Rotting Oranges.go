package medium

/*
You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
*/
func orangesRotting(grid [][]int) int {

	type node struct {
		i, j int
	}

	nodes := make([]node, 0)
	levels := make(map[node]int)

	var col []int
	for i := 0; i < len(grid); i++ {
		col = grid[i]
		for j, val := range col {
			if val == 2 {
				root := node{i, j}
				nodes = append(nodes, root)
				levels[root] = 0
			}
		}
	}

	getCh := func(i, j int) (node, bool) {
		var res node
		if i >= 0 && i < len(grid) &&
			(j >= 0 && j < len(grid[0])) &&
			grid[i][j] == 1 {
			res = node{i, j}
			return res, true
		}
		return res, false
	}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		pLvl := levels[node]

		if ch1, ok := getCh(node.i+1, node.j); ok {
			if _, ok := levels[ch1]; !ok {
				nodes = append(nodes, ch1)
				levels[ch1] = pLvl + 1
			}
		}
		if ch2, ok := getCh(node.i, node.j+1); ok {
			if _, ok := levels[ch2]; !ok {
				nodes = append(nodes, ch2)
				levels[ch2] = pLvl + 1
			}
		}
		if ch3, ok := getCh(node.i-1, node.j); ok {
			if _, ok := levels[ch3]; !ok {
				nodes = append(nodes, ch3)
				levels[ch3] = pLvl + 1
			}
		}
		if ch4, ok := getCh(node.i, node.j-1); ok {
			if _, ok := levels[ch4]; !ok {
				nodes = append(nodes, ch4)
				levels[ch4] = pLvl + 1
			}
		}

	}

	for i := 0; i < len(grid); i++ {
		col = grid[i]
		for j, val := range col {
			if val == 1 {
				node := node{i, j}
				if _, ok := levels[node]; !ok {
					return -1
				}
			}
		}
	}

	max := 0
	for _, l := range levels {
		if max < l {
			max = l
		}
	}

	return max
}
