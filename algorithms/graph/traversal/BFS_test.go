package traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// linearGraph: 0 — 1 — 2 — 3 — 4
var linearGraph = map[int][]int{
	0: {1},
	1: {0, 2},
	2: {1, 3},
	3: {2, 4},
	4: {3},
}

// branchGraph:
//
//	0 — 1 — 3
//	|       |
//	2       4
var branchGraph = map[int][]int{
	0: {1, 2},
	1: {0, 3},
	2: {0},
	3: {1, 4},
	4: {3},
}

func TestBFS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		start int
		graph map[int][]int
		want  map[int]int
	}{
		{
			name:  "distances from node 0",
			start: 0,
			graph: branchGraph,
			want:  map[int]int{0: 0, 1: 1, 2: 1, 3: 2, 4: 3},
		},
		{
			name:  "distances from node 3",
			start: 3,
			graph: branchGraph,
			want:  map[int]int{3: 0, 1: 1, 4: 1, 0: 2, 2: 3},
		},
		{
			name:  "single node graph",
			start: 0,
			graph: map[int][]int{0: {}},
			want:  map[int]int{0: 0},
		},
		{
			name:  "unreachable nodes excluded",
			start: 0,
			graph: map[int][]int{
				0: {1},
				1: {0},
				2: {3},
				3: {2},
			},
			want: map[int]int{0: 0, 1: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BFS(tt.start, tt.graph)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBFSMultiSource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		starts []int
		graph  map[int][]int
		want   map[int]int
	}{
		{
			// Узел 2 достигается обоими источниками на расстоянии 2 — берётся первый пришедший.
			name:   "two sources on linear graph",
			starts: []int{0, 4},
			graph:  linearGraph,
			want:   map[int]int{0: 0, 4: 0, 1: 1, 3: 1, 2: 2},
		},
		{
			name:   "single source matches BFS",
			starts: []int{0},
			graph:  branchGraph,
			want:   BFS(0, branchGraph),
		},
		{
			name:   "all nodes as sources",
			starts: []int{0, 1, 2, 3, 4},
			graph:  branchGraph,
			want:   map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BFSMultiSource(tt.starts, tt.graph)
			assert.Equal(t, tt.want, got)
		})
	}
}
