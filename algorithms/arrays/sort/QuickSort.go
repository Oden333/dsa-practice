package sort

func QuickSort(arr []int64, low, high int) {
	if low < high {
		// Находим индекс опорного элемента после разделения массива
		pivotIndex := Partition(arr, low, high)

		// Рекурсивно сортируем подмассивы до и после опорного элемента
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

// Функция для обмена элементов в массиве
func Swap(arr []int64, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Функция для выбора опорного элемента и разделения массива
func Partition(arr []int64, low, high int) int {
	// Выбираем опорный элемент (в данном случае - последний элемент массива)
	pivot := arr[high]

	// Индекс, который будет использоваться для разделения массива
	i := low - 1

	// Проходим по массиву
	for j := low; j < high; j++ {
		// Если текущий элемент меньше или равен опорному, меняем его местами с элементом на i+1 позиции
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Меняем местами опорный элемент и элемент на позиции i+1
	Swap(arr, i+1, high)

	// Возвращаем индекс опорного элемента
	return i + 1
}
