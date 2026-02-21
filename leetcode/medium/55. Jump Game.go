package medium

// Input: nums = [3,2,1,0,4]
// Output: false
func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	max := 0
	for i := range nums {
		if i > max {
			return false
		}
		if max < (i + nums[i]) {
			max = i + nums[i]
		}
	}
	return true
}
