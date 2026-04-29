package traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		start int
		graph map[int][]int
		want  []int
	}{
		{
			// 0 — 1 — 3
			// |       |
			// 2       4
			name:  "branch graph from node 0",
			start: 0,
			graph: branchGraph,
			want:  []int{0, 1, 3, 4, 2},
		},
		{
			name:  "branch graph from node 3",
			start: 3,
			graph: branchGraph,
			want:  []int{3, 1, 0, 2, 4},
		},
		{
			name:  "single node",
			start: 0,
			graph: map[int][]int{0: {}},
			want:  []int{0},
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
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := DFS(tt.start, tt.graph)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDFSMultiSource(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		starts []int
		graph  map[int][]int
		want   []int
	}{
		{
			// Два источника из разных компонент: 0→1 и 2→3.
			// DFS обходит их последовательно, а не одновременно.
			name:   "two disconnected components",
			starts: []int{0, 2},
			graph: map[int][]int{
				0: {1},
				1: {0},
				2: {3},
				3: {2},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			// Второй источник уже посещён первым — не добавляет дубликатов.
			name:   "second source already visited",
			starts: []int{0, 1},
			graph:  branchGraph,
			want:   []int{0, 1, 3, 4, 2},
		},
		{
			name:   "single source matches DFS",
			starts: []int{0},
			graph:  branchGraph,
			want:   DFS(0, branchGraph),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := DFSMultiSource(tt.starts, tt.graph)
			assert.Equal(t, tt.want, got)
		})
	}
}
