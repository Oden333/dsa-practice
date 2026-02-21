package easy

func RestoreString(s string, indices []int) string {

	res := make([]rune, len(s))
	for i, n := range s {
		res[indices[i]] = n
	}
	return string(res)

}
