package easy

import (
	q "leetcode/structs"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *q.ListNode
		want *q.ListNode
	}{
		{
			head: q.BuildList([]int{1, 2, 3, 4, 5}),
			want: q.BuildList([]int{5, 4, 3, 2, 1}),
		},
		{
			head: q.BuildList([]int{1, 2}),
			want: q.BuildList([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.head)

			if got.String() != tt.want.String() {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList_rec(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *q.ListNode
		want *q.ListNode
	}{
		{
			head: q.BuildList([]int{1, 2, 3, 4, 5}),
			want: q.BuildList([]int{5, 4, 3, 2, 1}),
		},
		{
			head: q.BuildList([]int{1, 2}),
			want: q.BuildList([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList_recursive(tt.head)

			if got.String() != tt.want.String() {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
