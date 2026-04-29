package search

func BinarySearch(in []int, searchFor int) (int, bool) {
	if len(in) == 0 {
		return 0, false
	}

	var left, right = 0, len(in) - 1

	for left <= right {
		var mid = ((right - left) / 2) + left

		if in[mid] == searchFor {
			return mid, true
		} else if in[mid] > searchFor {
			right = mid - 1
		} else if in[mid] < searchFor {
			left = mid + 1
		}
	}

	return 0, false
}
