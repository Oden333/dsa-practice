package easy

func FindClosestNumber(nums []int) int {
	minmax := int(1e5)
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) < abs(minmax) {
			minmax = nums[i]
		} else if nums[i] == abs(minmax) {
			minmax = nums[i]
		}
	}
	return int(minmax)
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
