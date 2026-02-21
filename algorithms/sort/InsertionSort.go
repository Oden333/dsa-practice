package sort

//До 32 элементов в массиве выполняется за O(n)
func InsertionSort(values []int64) {
	var value int64
	var newIndex int

	// Get each element
	for i := 1; i < len(values); i++ {
		value = values[i]
		newIndex = i - 1

		// Find a new idx position
		for newIndex >= 0 && value < values[newIndex] {
			values[newIndex+1] = values[newIndex]
			newIndex--
		}

		// Insert the value in the correct position
		values[newIndex+1] = value
	}
}
