package medium

// How can we prove that at least one duplicate number must exist in nums?
// Can you solve the problem in linear runtime complexity?

//? each integer is in the range [1, n] inclusive.
//! You must solve the problem without modifying the array nums and using only constant extra space.

func findDuplicate(n []int) int {
	var fast, slow = n[0], n[0]

	// 1) cycle check (speed 1x + 2x)
	for {
		slow = n[slow]
		fast = n[n[fast]]
		if fast == slow {
			break
		}
	}

	// 2) find the cycle point (with 1x speed)

	fast = n[0]
	for fast != slow {
		fast = n[fast]
		slow = n[slow]
	}

	return fast
}
