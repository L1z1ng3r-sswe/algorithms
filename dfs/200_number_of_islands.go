var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var zero, one byte = '0', '1'

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == zero {
			return
		}

		grid[r][c] = zero // mark as visited

		for _, dir := range dirs {
			dfs(r+dir[0], c+dir[1])
		}
	}

	var res int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == one {
				res++
				dfs(r, c)
			}
		}
	}

	return res
}

// time: O(n)
// space: O(1)