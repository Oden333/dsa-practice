package hard

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    []int
		want int
	}{
		{
			n:    []int{2, 1, 5, 6, 2, 3},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.n)
			if got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
