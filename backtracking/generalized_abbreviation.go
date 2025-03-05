func generateAbbreviations(word string) []string {
	res := make([]string, 0, int(math.Pow(2, float64(len(word)))))

	var backtrack func(start int, abbreviation []byte, count int)
	backtrack = func(start int, abbreviation []byte, count int) {
		if start == len(word) {
			if count > 0 {
				abbreviation = append(abbreviation, []byte(strconv.Itoa(count))...)
			}
			res = append(res, string(abbreviation))
			return
		}

		if count > 0 {
			backtrack(start+1, append(append(abbreviation, []byte(strconv.Itoa(count))...), word[start]), 0)
		} else {
			backtrack(start+1, append(abbreviation, word[start]), 0)
		}

		backtrack(start+1, abbreviation, count+1)
	}
	backtrack(0, []byte{}, 0)

	return res
}

// time: 2 ^ n where n is the length of word