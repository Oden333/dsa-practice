package medium

import (
	q "dsa/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *q.TreeNode
		k    int
	}

	q.Null = -1
	null := q.Null

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{q.BuildTree([]int{5, 3, 6, 2, 4, null, null, 1}), 3},
			want: 3,
		}, {
			args: args{q.BuildTree([]int{1, null, 2}), 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kthSmallest(tt.args.root, tt.args.k), "kthSmallest(%v, %v)", tt.args.root, tt.args.k)
		})
	}
}
