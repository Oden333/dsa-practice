package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		args  [][]int
		setup func() *GraphNode
	}{
		{
			name: "4 nodes cycle",
			args: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			setup: func() *GraphNode {
				n1, n2, n3, n4 := &GraphNode{Val: 1}, &GraphNode{Val: 2}, &GraphNode{Val: 3}, &GraphNode{Val: 4}
				n1.Neighbors = []*GraphNode{n2, n4}
				n2.Neighbors = []*GraphNode{n1, n3}
				n3.Neighbors = []*GraphNode{n2, n4}
				n4.Neighbors = []*GraphNode{n1, n3}
				return n1
			},
		},
		{
			name: "empty graph",
			args: [][]int{},
			setup: func() *GraphNode {
				return nil
			},
		},
		{
			name: "single node",
			args: [][]int{{}},
			setup: func() *GraphNode {
				return &GraphNode{Val: 1, Neighbors: []*GraphNode{}}
			},
		},
		{
			name: "two nodes connected",
			args: [][]int{{2}, {1}},
			setup: func() *GraphNode {
				a, b := &GraphNode{Val: 1}, &GraphNode{Val: 2}
				a.Neighbors = []*GraphNode{b}
				b.Neighbors = []*GraphNode{a}
				return a
			},
		},
		{
			name: "linear chain 1-2-3",
			args: [][]int{{2}, {1, 3}, {2}},
			setup: func() *GraphNode {
				n1, n2, n3 := &GraphNode{Val: 1}, &GraphNode{Val: 2}, &GraphNode{Val: 3}
				n1.Neighbors = []*GraphNode{n2}
				n2.Neighbors = []*GraphNode{n1, n3}
				n3.Neighbors = []*GraphNode{n2}
				return n1
			},
		},
		{
			name: "star graph (1 connected to 2,3,4)",
			args: [][]int{{2, 3, 4}, {1}, {1}, {1}},
			setup: func() *GraphNode {
				n1, n2, n3, n4 := &GraphNode{Val: 1}, &GraphNode{Val: 2}, &GraphNode{Val: 3}, &GraphNode{Val: 4}
				n1.Neighbors = []*GraphNode{n2, n3, n4}
				n2.Neighbors = []*GraphNode{n1}
				n3.Neighbors = []*GraphNode{n1}
				n4.Neighbors = []*GraphNode{n1}
				return n1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BuildGraph(tt.args)
			want := tt.setup()
			assert.Equal(t, PrintGraph(want), PrintGraph(got), "graph structure mismatch")
		})
	}
}

func TestPrintGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func() *GraphNode // Функция создаёт граф для теста
		expected string
	}{
		{
			name: "4 nodes cycle",
			setup: func() *GraphNode {
				n1, n2, n3, n4 := &GraphNode{Val: 1}, &GraphNode{Val: 2}, &GraphNode{Val: 3}, &GraphNode{Val: 4}
				n1.Neighbors = []*GraphNode{n2, n4}
				n2.Neighbors = []*GraphNode{n1, n3}
				n3.Neighbors = []*GraphNode{n2, n4}
				n4.Neighbors = []*GraphNode{n1, n3}
				return n1
			},
			expected: "[[2 4] [1 3] [2 4] [1 3]]",
		},
		{
			name: "single node",
			setup: func() *GraphNode {
				return &GraphNode{Val: 1, Neighbors: []*GraphNode{}}
			},
			expected: "[[]]",
		},
		{
			name: "two nodes connected",
			setup: func() *GraphNode {
				a, b := &GraphNode{Val: 1}, &GraphNode{Val: 2}
				a.Neighbors = []*GraphNode{b}
				b.Neighbors = []*GraphNode{a}
				return a
			},
			expected: "[[2] [1]]",
		},
		{
			name: "empty graph",
			setup: func() *GraphNode {
				return nil
			},
			expected: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := PrintGraph(tt.setup())
			assert.Equal(t, tt.expected, got)
		})
	}
}
