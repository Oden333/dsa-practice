package medium

import "testing"

func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{

		{
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			nums: []int{3, 3, 3, 3, 3},
			want: 3,
		},
		{
			nums: []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
