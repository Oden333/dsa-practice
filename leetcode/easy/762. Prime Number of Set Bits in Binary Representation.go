package easy

func countPrimeSetBits(left int, right int) int {
	var res int
	for left <= right {
		if isPrime(left) {
			res++
		}
		left++
	}
	return res
}

func isPrime(n int) bool {
	var res int
	for n > 0 {
		res += n & 1
		n /= 2
	}
	switch res {
	default:
		return false
	case 2, 3, 5, 7, 11, 13, 17, 19:
		// 2, 3, 5, 7, 11, 13, 17, 19, 23, 29,  31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97
		return true
	}
}
