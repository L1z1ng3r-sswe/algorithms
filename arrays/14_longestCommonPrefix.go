func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		currChar := strs[0][i]

		for _, str := range strs {
			if i >= len(str) || str[i] != currChar {
				return str[:i]
			}
		}
	}

	return strs[0]
}

// time: O(n)
// time: O(n)