func isPalindrome(s string) bool {
	lowerS := strings.ToLower(s)

	for i, j := 0, len(lowerS)-1; i < j; i, j = i+1, j-1 {
		// check i
		for i < j && !isAlphanumerical(lowerS[i]) {
			i++
		}
		// check j
		for i < j && !isAlphanumerical(lowerS[j]) {
			j--
		}

		if i == j {
			return true
		}

		if lowerS[i] != lowerS[j] {
			return false
		}
	}

	return true
}

func isAlphanumerical(char byte) bool {
	if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') {
		return true
	}

	return false
}

// time: O(N)