package medium

func maxArea(n []int) int {
	if len(n) == 0 {
		return 0
	}
	var (
		l, r     = 0, len(n) - 1
		max  int = 0
	)

	for l < r {
		if m := min(n[l], n[r]) * (r - l); m > max {
			max = m
		}

		if n[l] < n[r] {
			l++
		} else {
			r--
		}
	}

	return max
}
