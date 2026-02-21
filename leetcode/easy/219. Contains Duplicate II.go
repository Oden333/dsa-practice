//go:build ignore
// +build ignore

package main

import "fmt"

func containsNearbyDuplicate_slow(nums []int, k int) bool {
	switch len(nums) {
	case 1:
		return false
	case 2:
		return (nums[0] == nums[1]) && (k >= 1)
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if i+k < j-1 {
				break
			}
			if abs(i-j) <= k && nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, num := range nums {
		if j, found := seen[num]; found {
			if i-j <= k {
				return true
			}
		}
		seen[num] = i
	}
	return false
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
