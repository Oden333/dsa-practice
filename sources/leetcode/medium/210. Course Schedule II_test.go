package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "one course no prerequisites",
			args: args{1, [][]int{}},
			want: []int{0},
		},
		{
			name: "two courses linear",
			args: args{2, [][]int{{1, 0}}},
			want: []int{0, 1},
		},
		{
			name: "two courses linear",
			args: args{2, [][]int{{0, 1}}},
			want: []int{1, 0},
		},
		{
			name: "four courses linear chain",
			args: args{4, [][]int{{1, 0}, {2, 1}, {3, 2}}},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "cycle two courses",
			args: args{2, [][]int{{0, 1}, {1, 0}}},
			want: []int{},
		},
		{
			name: "cycle three courses",
			args: args{3, [][]int{{1, 0}, {2, 1}, {0, 2}}},
			want: []int{},
		},
		{
			name: "diamond shared prerequisite",
			args: args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "fan-in multiple prerequisites",
			args: args{5, [][]int{{4, 0}, {4, 1}, {4, 2}, {4, 3}}},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "prerequisite with higher index",
			args: args{4, [][]int{{0, 3}}},
			want: []int{3, 0, 1, 2},
		},
		{
			name: "no prerequisites",
			args: args{3, [][]int{}},
			want: []int{0, 1, 2},
		},
		{
			name: "shared dependency via two paths",
			args: args{6, [][]int{{2, 3}, {1, 2}, {0, 1}, {0, 4}, {4, 5}, {5, 1}}},
			want: []int{3, 2, 1, 5, 4, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, findOrder(tt.args.numCourses, tt.args.prerequisites))
		})
	}
}
