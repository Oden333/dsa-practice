package easy

func LongestCommonPrefix(strs []string) string {
	if strs == nil {
		return ""
	}

	max_prefix := []byte{}

	c := strs[0]
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < len(c) {
			c = strs[i]
		}
	}
	q := 0
	for i := 0; i < len(c); i++ {
		for n := 0; n < len(strs); n++ {
			if strs[n][i] == c[i] {
				q++
			} else {
				return string(max_prefix)
			}
			if q == len(strs) {
				max_prefix = append(max_prefix, c[i])
				q = 0
			}
		}
	}
	return string(max_prefix)
}
