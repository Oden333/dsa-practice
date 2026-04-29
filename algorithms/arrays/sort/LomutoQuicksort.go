package sort

/*
В функции LomutoQuickSort используется рекурсия для сортировки массива.
Функция находит индекс минимального элемента в неотсортированном массиве,
затем меняет его с первым элементом неотсортированного массива.

Наконец, рекурсивно вызывает функцию LomutoQuickSort для сортировки массива,
исключая первый элемент (он уже отсортирован).
*/
func LomutoQuickSort(arr []int64) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// find min value index in the unsorted array
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		// swap the found min value with the first element of the unsorted array
		arr[i], arr[minIdx] = arr[minIdx], arr[i]

		// recursively sort the array excluding the first element (as it is already sorted)
		LomutoQuickSort(arr[i+1:])
	}
}

func mergeSort(arr []int64) []int64 {
	// Step 1: Base case
	if len(arr) <= 1 {
		return arr
	}

	// Step 2: Divide array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Step 3: Merge two halves
	return merge(left, right)
}

func merge(left, right []int64) []int64 {
	// Step 4: Create a result array with combined length of two halves
	result := make([]int64, len(left)+len(right))
	i := 0
	j := 0

	// Step 5: Compare elements in both halves and append the smaller one to the result array
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	// Step 6: If there are remaining elements in the left half, append them to the result array
	for i < len(left) {
		result[i+j] = left[i]
		i++
	}

	// Step 7: If there are remaining elements in the right half, append them to the result array
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}

	return result
}
