var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	atlantic := make([][]bool, m)
	pacific := make([][]bool, m)

	for r := range m {
		pacific[r] = make([]bool, n)
		atlantic[r] = make([]bool, n)
	}

	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		visited[r][c] = true

		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]

			if (nr >= 0 && nr < m && nc >= 0 && nc < n) && // 1
				(!visited[nr][nc]) && // 2
				(heights[nr][nc] >= heights[r][c]) { // 3
				dfs(nr, nc, visited)
			}
		}
	}

	for c := range n {
		mMinusOne := m - 1 // not to recalculate
		dfs(0, c, pacific)
		dfs(mMinusOne, c, atlantic)
	}

	for r := range m {
		nMinusOne := n - 1 // not to recalculate
		dfs(r, 0, pacific)
		dfs(r, nMinusOne, atlantic)
	}

	// intersection check
	res := [][]int{}
	for r := range m {
		for c := range n {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

// time: O(m*n)
// space: O(m*n)