package hard

import (
	"slices"
)

/*
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !slices.Contains(wordList, endWord) {
		return 0
	}

	seen := make(map[string]int)
	seen[beginWord] = 1

	list := make([]string, 0, len(wordList)+1)
	list = append(list, beginWord)

	diff := func(w1, w2 string) bool {
		counter := 0
		r2 := []rune(w2)

		for i, r1 := range w1 {
			if r2[i] != r1 {
				counter++
			}
		}
		return counter == 1
	}

	var (
		level int
		node  string
	)
	for len(list) > 0 {
		node = list[0]
		list = list[1:]

		level = seen[node] + 1
		for _, word := range wordList {
			if _, ok := seen[word]; ok {
				continue
			}
			if !diff(word, node) {
				continue
			}

			if word == endWord {
				return level
			}

			list = append(list, word)
			seen[word] = level
		}
	}

	return 0
}
