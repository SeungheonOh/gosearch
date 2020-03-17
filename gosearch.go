package gosearch

func similarity(original, given string) (int, float64) {
	matching_characters := make([]int, len(original))
	matching := 0
	for i := 0; i < len(given); i++ { // given loop
		for j := matching_characters[matching]; j < len(original); j++ { //original loop
			if given[i] == original[j] {
				if matching > 1 && j == matching_characters[matching-1] {
					continue
				}
				matching_characters[matching] = j
				matching++
				break
			}
		}
	}

	panelty := 0
	for i := 1; i < matching; i++ {
		if matching_characters[i-1]+1 != matching_characters[i] {
			panelty++
		}
	}

	return matching, float64(matching)/float64(len(given)) - float64(panelty)*0.1
}

func search(given string, lines []string, similarity_limit float64) []string {
	ret := make([]string, 1)
	for _, l := range lines {
		_, sim := similarity(l, given)
		if sim >= similarity_limit {
			ret = append(ret, l)
		}
	}

	return ret
}
