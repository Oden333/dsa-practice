package traversal

import (
	"testing"

	"dsa/structs"

	"github.com/stretchr/testify/assert"
)

func TestBFSGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args [][]int
		want map[int]int
	}{
		{
			name: "4-node cycle",
			args: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			want: map[int]int{1: 0, 2: 1, 4: 1, 3: 2},
		},
		{
			name: "linear chain 1-2-3",
			args: [][]int{{2}, {1, 3}, {2}},
			want: map[int]int{1: 0, 2: 1, 3: 2},
		},
		{
			name: "single node",
			args: [][]int{{}},
			want: map[int]int{1: 0},
		},
		{
			name: "nil graph",
			args: [][]int{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BFSGraph(structs.BuildGraph(tt.args))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBFSTree(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *structs.TreeNode
		want [][]int
	}{
		{
			//     1
			//    / \
			//   2   3
			//  / \
			// 4   5
			name: "balanced tree",
			root: structs.BuildTree([]int{1, 2, 3, 4, 5}),
			want: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			//   1
			//    \
			//     2
			//      \
			//       3
			name: "right-skewed tree",
			root: structs.BuildTree([]int{1, structs.Null, 2, structs.Null, 3}),
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "single node",
			root: structs.BuildTree([]int{42}),
			want: [][]int{{42}},
		},
		{
			name: "nil tree",
			root: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := BFSTree(tt.root)
			assert.Equal(t, tt.want, got, "tree levels mismatch")
		})
	}
}
