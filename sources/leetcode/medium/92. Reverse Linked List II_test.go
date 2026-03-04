package medium

import (
	q "dsa/structs"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head  *q.ListNode
		left  int
		right int
		want  *q.ListNode
	}{
		{
			head:  q.BuildList([]int{1, 2, 3, 4, 5}),
			left:  2,
			right: 4,
			want:  q.BuildList([]int{1, 4, 3, 2, 5}),
		},
		{
			head:  q.BuildList([]int{5}),
			left:  1,
			right: 1,
			want:  q.BuildList([]int{5}),
		},
		{
			head:  q.BuildList([]int{3, 5}),
			left:  1,
			right: 2,
			want:  q.BuildList([]int{5, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.head, tt.left, tt.right)
			if got.String() != tt.want.String() {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
