package gosearch

import "math"

/*
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
*/

func _similarity(s1, s2 string, i1, i2 int) float64 {
	if i1 < 0 {
		return float64(i2 + 1)
	}
	if i2 < 0 {
		return float64(i1 + 1)
	}

	if s1[i1] == s2[i2] {
		return _similarity(s1, s2, i1-1, i2-1)
	}

	del := _similarity(s1, s2, i1-1, i2)
	ins := _similarity(s1, s2, i1, i2-1)
	sub := _similarity(s1, s2, i1-1, i2-1)

	return math.Min(del, math.Min(ins, sub)) + 1
}

func similarity(s1, s2 string) float64 {
	return _similarity(s1, s2, len(s1)-1, len(s2)-1)
}

func search(given string, lines []string, similarity_limit float64) []string {
	ret := make([]string, 1)
	for _, l := range lines {
		sim := similarity(l, given)
		if sim <= similarity_limit {
			ret = append(ret, l)
		}
	}

	return ret
}
