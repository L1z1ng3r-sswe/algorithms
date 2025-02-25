func generateAbbreviations(word string) []string {
	res := make([]string, 0, int(math.Pow(2, float64(len(word)))))

	var backtrack func(start int, abbreviation []byte, count byte)
	backtrack = func(start int, abbreviation []byte, count byte) {
		if start >= len(word) {
			if count > 1 {
				abbreviation = append(abbreviation, count-1)
			}

			res = append(res, string(abbreviation))
			return
		}

		if count <= 0 {
			backtrack(start+1, append(abbreviation, word[start]), count+1)
		} else {
			backtrack(start+1, append(abbreviation, count, word[start]), 0)
		}

		backtrack(start+1, abbreviation, count+1)
	}
	backtrack(0, []byte{}, 0)

	return res
}

// time: 2^n