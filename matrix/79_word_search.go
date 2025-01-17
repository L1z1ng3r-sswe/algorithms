func exist(board [][]byte, word string) bool { // for ascii only
	m, n := len(board), len(board[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if dfs(board, word, 0, r, c) {
				return true
			}
		}
	}

	return false
}

var directions = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} // left, right, top, bottom

func dfs(board [][]byte, word string, idx int, r, c int) bool {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) || board[r][c] == '#' || board[r][c] != word[idx] {
		return false
	}
	if idx == len(word)-1 {
		return true
	}

	charTp := board[r][c]
	board[r][c] = '#'

	for _, direction := range directions {
		newR := r + direction[0]
		newC := c + direction[1]

		if dfs(board, word, idx+1, newR, newC) {
			return true
		}
	}

	board[r][c] = charTp // backtrack
	return false
}

// m, n, l = len(mx), len(mx[0]), len(word)
// time = n*m * (4^l)
// space = O(n*m)