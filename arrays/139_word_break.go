func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	set := make(map[string]struct{}, len(s))
	shortest := len(wordDict[0])
	longest := len(wordDict[0])
	for _, word := range wordDict {
		if len(word) < shortest {
			shortest = len(word)
		} else if len(word) > longest {
			longest = len(word)
		}

		set[word] = struct{}{}
	}

	for end := shortest; end <= len(s); end++ {
		for start := 0; start < end; start++ {
			if _, ok := set[s[start:end]]; ok && dp[start] {
				dp[end] = true
			}
		}
	}

	return dp[len(s)]
}

// n = len(s)
// time: O(n^2)
// space: O(n)

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	set := make(map[string]struct{}, len(wordDict))

	minLen := len(s)
	maxLen := 0
	for _, word := range wordDict {
		set[word] = struct{}{}
		l := len(word)
		if l < minLen {
			minLen = l
		}
		if l > maxLen {
			maxLen = l
		}
	}

	for end := minLen; end <= n; end++ {
		// Only consider substring lengths between minLen and maxLen
		for l := minLen; l <= maxLen; l++ {
			start := end - l
			if start < 0 {
				continue
			}
			if dp[start] {
				if _, ok := set[s[start:end]]; ok {
					dp[end] = true
					break
				}
			}
		}
	}

	return dp[n]
}
