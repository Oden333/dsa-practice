package easy

func isHappy(n int) bool {

	for {
		switch {
		case n == 1 || n == 7:
			return true
		case n < 10:
			return false
		default:
			n = calc(n)
		}
	}
}

func calc(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
