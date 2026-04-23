package medium

import "dsa/structs"

type Node = structs.GraphNode

func cloneGraph(head *Node) *Node {
	if head == nil {
		return nil
	}

	links := make(map[int]*Node)

	var trav func(idx int, node *Node) *Node
	trav = func(idx int, node *Node) *Node {
		if links[idx] != nil {
			return links[idx]
		}

		res := new(Node)
		links[idx] = res

		res.Val = node.Val
		res.Neighbors = make([]*Node, len(node.Neighbors))
		for i, cur := range node.Neighbors {
			res.Neighbors[i] = trav(cur.Val, cur)
		}

		return res
	}

	return trav(1, head)
}
