package medium

// Input: nums = [2,3,1,1,4]
// Output: 2

func Jump(nums []int) int {
	// если меньше 2 элементов всегда будет 0
	if len(nums) < 2 {
		return 0
	}

	step := 1
	maxReach := nums[0]
	prevReach := 0

	for i := 0; i < len(nums)-1; i++ {
		if i > prevReach {
			step++
			prevReach = maxReach
		}

		if nums[i] != 0 {
			maxReach = max(maxReach, nums[i]+i)
		}

		if maxReach >= len(nums)-1 {
			return step
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
