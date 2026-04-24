package medium

import (
	"dsa/structs"
	. "dsa/structs"
	q "dsa/structs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	t.Parallel()
	type args struct {
		root      *TreeNode
		targetSum int
	}

	q.Null = -1
	null := q.Null

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root:      structs.BuildTree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}),
				targetSum: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			args: args{
				root:      structs.BuildTree([]int{1, 2, 3}),
				targetSum: 5,
			},
			want: [][]int{},
		},
		{
			args: args{
				root:      structs.BuildTree([]int{1, 2}),
				targetSum: 5,
			},
			want: [][]int{},
		},
		{
			// промежуточная нода совпадает по сумме, но не является листом
			args: args{
				root:      structs.BuildTree([]int{1, 2, null, 3}),
				targetSum: 3,
			},
			want: [][]int{},
		},
		{
			// сумма совпадает и у промежуточной ноды, и у листа
			args: args{
				root:      structs.BuildTree([]int{1, 1, null, 1}),
				targetSum: 3,
			},
			want: [][]int{{1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, pathSum(tt.args.root, tt.args.targetSum), "pathSum(%v, %v)", tt.args.root, tt.args.targetSum)
		})
	}
}
