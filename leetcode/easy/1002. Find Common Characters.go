package easy

func CommonChars(words []string) []string {

	curLettersSet := getLetterToCount(words[0])

	for i := 0; i < len(words)-1; i++ {
		nextMap := getLetterToCount(words[i+1])
		lettersMapProcessing(curLettersSet, nextMap)
	}

	var commonLetters []string

	for key, value := range curLettersSet {
		for i := 0; i < value; i++ {
			commonLetters = append(commonLetters, string(key))
		}
	}

	return commonLetters
}

func lettersMapProcessing(source, test map[rune]int) {
	for key := range source {
		testCount, isKeyContained := test[key]
		if !isKeyContained {
			delete(source, key)
			continue
		}

		if testCount < source[key] {
			source[key] = testCount
		}
	}
}

func getLetterToCount(word string) map[rune]int {
	curLettersSet := make(map[rune]int)

	for _, letter := range word {
		curLettersSet[letter]++
	}
	return curLettersSet
}

/* func CommonChars(words []string) []string {
	c := words[0]
	for i := 0; i < len(words); i++ {
		if len(words[i]) < len(c) {
			c = words[i]
		}
	}
	var curByte byte
	var curWord string
	byteMap := make(map[byte]int, len(c))
	for i := 0; i < len(c); i++ {
		curByte = c[i]
		if _, ok := byteMap[curByte]; !ok {
			byteMap[curByte] = 0
		} else {
			continue
		}
		//"acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"
		for n := 0; n < len(words); n++ {
			curWord = words[n]
			//acabcddd
			for l := 0; l < len(curWord); l++ {
				if curWord[l] == curByte {
					byteMap[curByte]++
				}
			}
		}

	}
	var counter int
	commonChars := []string{}
	for key, value := range byteMap {
		counter = int(value / len(words))
		for i := 0; i < counter; i++ {
			commonChars = append(commonChars, string(key))
		}
	}
	return commonChars
} */
