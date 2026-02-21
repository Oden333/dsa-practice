package search

func BinarySearch(in []int, searchFor int) (int, bool) {
	if len(in) == 0 {
		return 0, false
	}

	var first, last = 0, len(in) - 1

	for first <= last {
		var mid = ((last - first) / 2) + first

		if in[mid] == searchFor {
			return mid, true
		} else if in[mid] > searchFor {
			last = mid - 1
		} else if in[mid] < searchFor {
			first = mid + 1
		}
	}

	return 0, false
}
