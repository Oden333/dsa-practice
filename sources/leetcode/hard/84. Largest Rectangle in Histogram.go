package hard

type lra struct {
	val, idx int
}

func largestRectangleArea(n []int) int {
	var (
		max      int
		pse, nse []int
	)

	nse = nextSmallerNumber(n) // считаем next smaller element, где будут лежать id некст уменьшенного элемента.
	pse = prevSmallerNumber(n) // считаем prev smaller element, где будут лежать id предыдущего уменьшенного элемента.

	for i := range len(n) {
		// в формуле берём эти индексы, и сдвигаем nse влево, а pse вправо. Так, получим неуменьшаюшийся относительно h[i] прямоугольник, и проверим все варианты площадей
		if tmp := (1 + (nse[i] - 1) - (pse[i] + 1)) * n[i]; tmp > max {
			max = tmp
		}
	}

	return max
}

func nextSmallerNumber(n []int) []int {
	var (
		val    int
		nse    = make([]int, len(n))
		mStack = make([]lra, 0)
	)

	for j := len(n) - 1; j >= 0; j-- {
		val = n[j]

		for len(mStack) > 0 {
			if tmp := mStack[len(mStack)-1]; tmp.val < val {
				nse[j] = tmp.idx
				break
			}

			mStack = mStack[:len(mStack)-1]
		}

		if len(mStack) == 0 {
			nse[j] = len(n)
		}

		mStack = append(mStack, lra{val, j})
	}

	return nse
}

func prevSmallerNumber(n []int) []int {
	var (
		val    int
		nse    = make([]int, len(n))
		mStack = make([]lra, 0)
	)

	for i := 0; i < len(n); i++ {
		val = n[i]

		for len(mStack) > 0 {
			if tmp := mStack[len(mStack)-1]; tmp.val < val {
				nse[i] = tmp.idx
				break
			}

			mStack = mStack[:len(mStack)-1]
		}

		if len(mStack) == 0 {
			nse[i] = -1
		}

		mStack = append(mStack, lra{val, i})
	}

	return nse
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
