package easy

import (
	q "dsa/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binaryTreePaths(t *testing.T) {
	type args struct {
		root *q.TreeNode
	}

	q.Null = -101
	null := q.Null

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{q.BuildTree([]int{1, 2, 3, null, 5})},
			want: []string{"1->2->5", "1->3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, binaryTreePaths(tt.args.root), "binaryTreePaths(%v)", tt.args.root)
		})
	}
}
