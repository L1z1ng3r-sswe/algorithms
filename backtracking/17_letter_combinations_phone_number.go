var letterMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}

	var backtrack func(start int, comb []byte)
	backtrack = func(start int, comb []byte) {
		if start >= len(digits) {
			res = append(res, string(comb))
			return
		}

		for _, letter := range letterMap[digits[start]] {
			backtrack(start+1, append(comb, letter))
		}
	}
	backtrack(0, []byte{})

	return res
}