package easy

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		left  int
		right int
		want  int
	}{
		{
			left:  6,
			right: 10,
			want:  4,
		}, {
			left:  10,
			right: 15,
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPrimeSetBits(tt.left, tt.right)
			if tt.want != got {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool 
	}{
		{
			n:    (1 << 63) - 1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPrime(tt.n)
			sum := func() (res int) {
				n := tt.n
				for n > 0 {
					res += n & 1
					n /= 2
				}
				return
			}()
			t.Logf("res: %b (%d) %t", tt.n, sum, got)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
