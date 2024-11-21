std::vector<std::vector<int>> directions = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

int dfs(std::vector<std::vector<int>>& dp, std::vector<std::vector<int>>& matrix, size_t m, size_t n, size_t r, size_t c, int prevVal) {
  if (!(r >= 0 && c >= 0 && r < m && c < n && matrix[r][c] != -1 && matrix[r][c] > prevVal)) {
    return 0;
  }

  if (dp[r][c] != -1) return dp[r][c] + 1;

  int temp = matrix[r][c];
  matrix[r][c] = -1;
  
  int longestPath = 0;
  for (const std::vector<int>& direction : directions) {
    longestPath = std::max(longestPath, dfs(dp, matrix, m, n, r + direction[0], c + direction[1], temp));
  }

  // backtrack
  matrix[r][c] = temp;

  // cache
  dp[r][c] = longestPath;
  
  return ++longestPath;
}

class Solution {
public:
    int longestIncreasingPath(vector<vector<int>>& matrix) {
      size_t m = matrix.size();
      size_t n = matrix[0].size();

      std::vector<std::vector<int>> dp;
      dp.reserve(m);

      for (size_t i = 0; i < m; ++i) {
        std::vector<int> raw;
        raw.reserve(n);

        for (size_t j = 0; j < n; ++j) {
          raw.push_back(-1);
        }

        dp.push_back(raw);
      }

      int longestPath = 0;
      for (size_t r = 0; r < m; ++r) {
        for (size_t c = 0; c < n; ++c) {
          longestPath = std::max(longestPath, dfs(dp, matrix, m, n, r, c, -1));
        }
      }

      return longestPath;
    }
};