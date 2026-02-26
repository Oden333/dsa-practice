package medium

import (
	q "leetcode/structs"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *q.ListNode
		want *q.ListNode
	}{

		{
			head: q.BuildList([]int{1, 2, 3, 4}),
			want: q.BuildList([]int{2, 1, 4, 3}),
		},
		{
			head: q.BuildList([]int{}),
			want: q.BuildList([]int{}),
		},
		{
			head: q.BuildList([]int{1, 2, 3}),
			want: q.BuildList([]int{2, 1, 3}),
		},
		{
			head: q.BuildList([]int{1, 2, 3, 4, 5}),
			want: q.BuildList([]int{2, 1, 4, 3, 5}),
		},
		{
			head: q.BuildList([]int{2, 5, 3, 4, 6, 2, 2}),
			want: q.BuildList([]int{5, 2, 4, 3, 2, 6, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.head)
			if got.String() != tt.want.String() {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
