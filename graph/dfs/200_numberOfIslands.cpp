// https://leetcode.com/problems/number-of-islands/description/

#include <iostream>
#include <vector>
#include <functional>

class Solution {
public:
  int numIslands(std::vector<std::vector<char>>& grid) {
    int islands = 0;

    size_t m = grid.size();
    size_t n = grid[0].size();

    std::function<void(int, int)> dfs = [&](int r, int c) {
      if (grid[r][c] == '0') {
        return;
      }

      grid[r][c] = '0';

      if (c+1 < m) dfs(r, c+1);
      if (c-1 >= 0)  dfs(r, c-1);
      if (r+1 < n) dfs(r+1, c);
      if (r-1 >= 0)  dfs(r-1, c);
    };

    for (size_t i = 0; i < m; ++i) {
      for (size_t j = 0; j < n; ++j) {
        if (grid[i][j] == '1') {
          ++islands;
          dfs(i, j);
        }
      }
    }

    return islands;
  }
};