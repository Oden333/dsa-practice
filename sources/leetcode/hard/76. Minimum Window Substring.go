package hard

import (
	"maps"
	"strings"
)

// https://www.youtube.com/watch?v=jSto0O4AJbM

func minWindow1(s string, t string) string { // O(m+n)
	tmap := make(map[rune]int)
	for _, r := range []rune(t) {
		tmap[r]++
	}

	tmp := []rune(s)
	copy := maps.Clone(tmap)
	var l, r = 0, len(tmp) - 1
	for l < r {
		if _, ok := copy[tmp[l]]; !ok {
			continue
		}
	}

	return ""
}

func minWindow(s string, t string) string {
	tmap := make(map[rune]int)
	for _, r := range []rune(t) {
		tmap[r]++
	}

	minStr := ""
	var (
		smap   map[rune]int
		window strings.Builder
	)

	str := []rune(s)
	for i, f := range str {
		smap = maps.Clone(tmap)
		if _, ok := smap[f]; !ok {
			continue
		}

		window = strings.Builder{}

		for _, r := range str[i:] {
			window.WriteRune(r)
			if _, ok := smap[r]; !ok {
				continue
			}

			smap[r]--
			if smap[r] == 0 {
				delete(smap, r)
			}

			if len(smap) == 0 {
				break
			}
		}
		if len(smap) != 0 {
			continue
		}

		if minStr == "" || len(window.String()) < len(minStr) {
			minStr = window.String()
		}

		if len(minStr) == len(t) {
			return minStr
		}

		window.Reset()
	}

	return minStr
}
