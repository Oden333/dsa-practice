package traversal

import (
	"testing"

	"dsa/structs"

	"github.com/stretchr/testify/assert"
)

func TestDFSGraph(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "4-node cycle from node 1",
			args: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "linear chain 1-2-3",
			args: [][]int{{2}, {1, 3}, {2}},
			want: []int{1, 2, 3},
		},
		{
			name: "single node",
			args: [][]int{{}},
			want: []int{1},
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
			got := DFSGraph(structs.BuildGraph(tt.args))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDFSTreeOrders(t *testing.T) {
	t.Parallel()

	//       4
	//      / \
	//     2   6
	//    / \ / \
	//   1  3 5  7
	root := structs.BuildTree([]int{4, 2, 6, 1, 3, 5, 7})

	tests := []struct {
		name string
		fn   func(*structs.TreeNode) []int
		want []int
	}{
		{
			// корень → левое → правое
			name: "pre-order",
			fn:   DFSPreOrder,
			want: []int{4, 2, 1, 3, 6, 5, 7},
		},
		{
			// левое → корень → правое (для BST даёт отсортированный порядок)
			name: "in-order",
			fn:   DFSInOrder,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			// левое → правое → корень
			name: "post-order",
			fn:   DFSPostOrder,
			want: []int{1, 3, 2, 5, 7, 6, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.fn(root)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDFSTreeEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *structs.TreeNode
		want []int
	}{
		{name: "nil tree", root: nil, want: nil},
		{name: "single node", root: structs.BuildTree([]int{1}), want: []int{1}},
		{
			//   1
			//    \
			//     2
			//      \
			//       3
			name: "right-skewed tree",
			root: structs.BuildTree([]int{1, structs.Null, 2, structs.Null, 3}),
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := DFSPreOrder(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
