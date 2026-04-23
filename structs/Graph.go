package structs

import (
	"fmt"
)

// GraphNode — нода графа
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

// For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on.
//
// The graph is represented in the test case using an adjacency list.
// ? An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
//
// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
func BuildGraph(arr [][]int) *GraphNode {
	if len(arr) == 0 {
		return nil
	}

	links := make(map[int]*GraphNode)

	var build func(idx int) *GraphNode
	build = func(idx int) *GraphNode {
		if links[idx] != nil {
			return links[idx]
		}

		node := &GraphNode{
			Val:       idx + 1,
			Neighbors: make([]*GraphNode, len(arr[idx])),
		}
		links[idx] = node

		for i, v := range arr[idx] {
			node.Neighbors[i] = build(v - 1)
		}

		return node
	}

	return build(0)
}

func PrintGraph(head *GraphNode) string {
	nodes := make(map[int][]int)

	var trav func(node *GraphNode)
	trav = func(node *GraphNode) {
		if node == nil {
			return
		}
		level := node.Val
		if nodes[level] != nil {
			return
		}

		nodes[level] = make([]int, len(node.Neighbors))
		for i, val := range node.Neighbors {
			nodes[level][i] = val.Val
			trav(val)
		}
	}

	trav(head)

	arr := make([][]int, len(nodes))
	for level, links := range nodes {
		arr[level-1] = links
	}

	// res := strings.Builder{}
	// for _, links := range arr {
	// 	res.WriteString(fmt.Sprintf("%v", links))
	// }

	// return res.String()

	return fmt.Sprintf("%v", arr)

	// return strings.ReplaceAll(fmt.Sprintf("%v", arr)," ",",")
}
