package medium

import (
	"fmt"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{

		{
			nums1: []int{1, 7, 11},
			nums2: []int{2, 4, 6},
			k:     3,
			want:  [][]int{{1, 2}, {1, 4}, {1, 6}},
		},

		{
			nums1: []int{1, 1, 2},
			nums2: []int{1, 2, 3},
			k:     2,
			want:  [][]int{{1, 1}, {1, 1}},
		},

		{
			nums1: []int{1, 2, 4, 5, 6},
			nums2: []int{3, 5, 7, 9},
			k:     3,
			want:  [][]int{{1, 3}, {2, 3}, {1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kSmallestPairs(tt.nums1, tt.nums2, tt.k)
			if fmt.Sprintf("%v", tt.want) != fmt.Sprintf("%v", got) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
