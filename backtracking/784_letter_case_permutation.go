func letterCasePermutation(s string) []string {
	res := make([]string, 0, len(s))

	dfs(s, 0, make([]byte, 0, len(s)), &res)

	return res
}

func dfs(s string, idx int, currS []byte, res *[]string) {
	if idx == len(s) {
		*res = append(*res, string(currS))
		return
	}

	runeChar := rune(s[idx])

	dfs(s, idx+1, append(currS, s[idx]), res)

	if unicode.IsLetter(runeChar) {
		if unicode.IsUpper(runeChar) {
			dfs(s, idx+1, append(currS, byte(unicode.ToLower(runeChar))), res)
		} else {
			dfs(s, idx+1, append(currS, byte(unicode.ToUpper(runeChar))), res)
		}
	}
}