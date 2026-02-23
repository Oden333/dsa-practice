package hard

import "testing"

func Test_minWindow(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want string
	}{
		{
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			s:    "a",
			t:    "aa",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
