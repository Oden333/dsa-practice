package medium

import "slices"

func findOrder(numCourses int, prerequisites [][]int) []int {
	type node struct {
		num int
		ch  []*node
	}

	deps := make(map[int]*node)
	var cur, child *node

	for _, req := range prerequisites {
		if cur = deps[req[0]]; cur == nil {
			cur = new(node)
			cur.num = req[0]
			deps[req[0]] = cur
		}

		if child = deps[req[1]]; child == nil {
			child = new(node)
			child.num = req[1]
			deps[req[1]] = child
		}

		cur.ch = append(cur.ch, child)
	}

	seen := make(map[int]struct{})
	res := make([]int, 0, numCourses)

	var trav func(node *node, list *[]int) bool
	trav = func(node *node, list *[]int) bool {
		if node == nil {
			return false
		}

		if _, ok := seen[node.num]; ok {
			return false
		}

		*list = append(*list, node.num)

		for _, v := range node.ch {
			if slices.Contains(*list, v.num) {
				return true
			}

			if trav(v, list) {
				return true
			}
		}

		if !slices.Contains(res, node.num) {
			res = append(res, node.num)
		}
		seen[node.num] = struct{}{}
		*list = (*list)[:len(*list)-1]

		return false
	}

	for i := range numCourses {
		cur = deps[i]

		if cur == nil {
			res = append(res, i)
			seen[i] = struct{}{}
			continue
		}

		if _, ok := seen[i]; ok {
			continue
		}

		list := []int{}
		if trav(cur, &list) && len(list) != 0 {
			return []int{}
		}
	}

	return res
}
