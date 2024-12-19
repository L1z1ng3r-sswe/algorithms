func longestCommonPrefix(strs []string) string {
	var res string
	minLen := smallestStr(strs)

	for i := range minLen {
		currSymbol := strs[0][i]

		for _, word := range strs {
			if word[i] != currSymbol {
				return res
			}
		}
		res += string(currSymbol)
	}

	return res
}

func smallestStr(strs []string) int {
	res := math.MaxInt64

	for _, word := range strs {
		res = min(res, len(word))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}