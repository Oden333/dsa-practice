package easy

func Intersection(nums [][]int) []int {

	start := getNumbersToCount(nums[0])
	for i := 0; i < len(nums)-1; i++ {
		numbersMapProcessing(start, getNumbersToCount(nums[i+1]))
	}
	//{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}
	//{3,4}
	fin := make([]int, 0)
	for key := range start {
		fin = append(fin, key)
	}
	quickSort(fin, 0, len(fin)-1)
	return fin
}

func numbersMapProcessing(finalMap, currentMap map[int]int) {
	for key := range finalMap {
		_, isKeyContained := currentMap[key]
		if !isKeyContained {
			delete(finalMap, key)
			continue
		}
	}
}

func getNumbersToCount(slice []int) map[int]int {
	curSliceSet := make(map[int]int)
	for _, number := range slice {
		curSliceSet[number]++
	}
	return curSliceSet
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// Находим индекс опорного элемента после разделения массива
		pivotIndex := partition(arr, low, high)

		// Рекурсивно сортируем подмассивы до и после опорного элемента
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// Функция для обмена элементов в массиве
func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Функция для выбора опорного элемента и разделения массива
func partition(arr []int, low, high int) int {
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
	swap(arr, i+1, high)

	// Возвращаем индекс опорного элемента
	return i + 1
}
