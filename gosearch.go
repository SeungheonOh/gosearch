package gosearch

const DefaultSimilarity = 0.75

func similarity(original, given []rune, matchingChars []int) (eqCount int, similarity float64) {
	for i := 0; i < len(given); i++ { // given loop
		for j := 0; j < len(original); j++ { //original loop
			if given[i] == original[j] {
				if l := len(matchingChars); l > 1 && j == matchingChars[l-2] {
					continue
				}

				matchingChars = append(matchingChars, j)
				break
			}
		}
	}

	var penalty = 0

	for i := 1; i < len(matchingChars); i++ {
		if matchingChars[i-1]+1 != matchingChars[i] {
			penalty++
		}
	}

	eqCount = len(matchingChars)
	similarity = float64(len(matchingChars))/float64(len(given)) - float64(penalty)*0.1

	return
}

func Search(given string, lines []string) []int {
	return SearchLimit(given, lines, DefaultSimilarity)
}

func SearchLimit(given string, lines []string, similarityLimit float64) []int {
	if len(lines) == 0 {
		return []int{}
	}

	ret := make([]int, 0, len(lines))

	matchingChars := make([]int, 0, len(lines[0]))

	givenRunes := []rune(given)

	for i, l := range lines {
		matchingChars = matchingChars[:0]

		if _, sim := similarity([]rune(l), givenRunes, matchingChars); sim >= similarityLimit {
			ret = append(ret, i)
		}
	}

	return ret
}
