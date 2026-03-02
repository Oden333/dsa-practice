package hard

type lra struct {
	val, idx int
}



func largestRectangleArea(n []int) int {
	var (
		max, val int
		copy     = make([]lra, len(n))
		mStack   = make([]lra, 0)
	)

	for j := len(n) - 1; j >= 0; j-- {
		val = n[j]

		for len(mStack) > 0 {
			if tmp := mStack[len(mStack)-1]; tmp.val < val {
				copy[j] = tmp
				break
			}

			mStack = mStack[:len(mStack)-1]
		}

		if len(mStack) == 0 {
			copy[j] = lra{val, len(n) - 1}

		}

		mStack = append(mStack, lra{val, j})
	}

	for i := 0; i < len(n); i++ {
		
	}

	return max
}

/*
func largestRectangleArea(n []int) int {
	var (
		max  int
		l, r = 0, len(n) - 1
	)

	for l < r {
		if tmp := min(n[l], n[r]); tmp*(r-l) > max {
			max = tmp * tmp * (r - l)
		}

		if n[l] < n[r] {
			l++
		} else {
			r--
		}
	}

	return max
}
*/
