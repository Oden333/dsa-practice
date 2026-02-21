//go:build ignore
// +build ignore

package main

import "fmt"

func numSubarraysWithSum_slow(nums []int, k int) int {
	switch {
	case len(nums) == 1:
		if (nums[0] == 1 && k == 1) || (nums[0] == 0 && k == 0) {
			return 1
		}
		return 0
	case len(nums) == k:
		for _, a := range nums {
			if a == 0 {
				return 0
			}
		}
		return 1
	}
	var counter int
	for i := range nums {
		sum := nums[i]
		if sum == k {
			counter++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				counter++
				continue
			}
			if sum < k {
				continue
			}
			if sum > k {
				break
			}

		}
	}
	return counter
}

// 1⃣Использовать словарь для хранения количества встреченных сумм префиксов.
// Инициализировать текущую сумму и счетчик подмассивов с нулевыми значениями.

// 2⃣Пройти по массиву и обновить текущую сумму.
// Если текущая сумма минус цель уже в словаре, добавить количество таких префиксов к счетчику подмассивов.
// Обновить словарь префиксных сумм.

// 3⃣Вернуть счетчик подмассивов.
func numSubarraysWithSum(nums []int, k int) int {
	switch {
	case len(nums) == 1:
		if (nums[0] == 1 && k == 1) || (nums[0] == 0 && k == 0) {
			return 1
		}
		return 0
	case len(nums) == k:
		for _, a := range nums {
			if a == 0 {
				return 0
			}
		}
		return 1
	}

	counters := make(map[int]int)
	counters[0] = 1

	sum := 0
	prefix := 0

	for _, i := range nums {
		prefix += i

		if val, ok := counters[prefix-k]; ok {
			sum += val
		}

		counters[prefix]++
	}

	return sum
}

func main() {
	fmt.Println(^(1 << 4) - 1)
}
