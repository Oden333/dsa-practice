package hard

import (
	q "dsa/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *q.TreeNode
	}
	q.Null = -1001
	null := q.Null

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{q.BuildTree([]int{1, 2, 3})},
			want: 6,
		}, {
			args: args{q.BuildTree([]int{-10, 9, 20, null, null, 15, 7})},
			want: 42,
		}, {
			args: args{q.BuildTree([]int{0})},
			want: 0,
		}, {
			args: args{q.BuildTree([]int{-2, 1})},
			want: 1,
		}, {
			args: args{q.BuildTree([]int{2, -1, -2})},
			want: 2,
		}, {
			args: args{q.BuildTree([]int{-3})},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxPathSum(tt.args.root), "maxPathSum(%v)", tt.args.root)
		})
	}
}
