func numDecodings(s string) int {
	dp := make(map[int]int, len(s))

	var dfs func(idx int) int
	dfs = func(idx int) int {
		if idx == len(s) {
			return 1
		}

		if val, ok := dp[idx]; ok {
			return val
		}

		var res int

		if _, ok := dict[s[idx:idx+1]]; ok {
			res += dfs(idx + 1)
		}

		if idx+1 < len(s) {
			if _, ok := dict[s[idx:idx+2]]; ok {
				res += dfs(idx + 2)
			}
		}

		dp[idx] = res

		return res
	}

	return dfs(0)
}

// n = len(s)
// time: O(n)
// space: recursive stack + dp lenght = O(n)

var dict = map[string]struct{}{
	"1":  {},
	"2":  {},
	"3":  {},
	"4":  {},
	"5":  {},
	"6":  {},
	"7":  {},
	"8":  {},
	"9":  {},
	"10": {},
	"11": {},
	"12": {},
	"13": {},
	"14": {},
	"15": {},
	"16": {},
	"17": {},
	"18": {},
	"19": {},
	"20": {},
	"21": {},
	"22": {},
	"23": {},
	"24": {},
	"25": {},
	"26": {},
}