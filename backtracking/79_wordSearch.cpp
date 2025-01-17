// https://leetcode.com/problems/word-search/description/

#include <iostream>
#include <vector>
#include <string>

bool dfs(std::vector<std::vector<char>>& board, size_t m, size_t n, int r, int c, const std::string& word, size_t strIdx) {
  if (strIdx == word.size()) {
    return true;  
  }
  
  if (r-1 >= 0 && board[r-1][c] == word[strIdx]) {
    char temp = board[r-1][c];
    board[r-1][c] = '.';
    if (dfs(board, m, n, r-1, c, word, ++strIdx)) {
      return true;
    }
    board[r-1][c] = temp;
  }

  if (r+1 < m && board[r+1][c] == word[strIdx]) {
    char temp = board[r+1][c];
    board[r+1][c] = '.';
    if (dfs(board, m, n, r+1, c, word, ++strIdx)) {
      return true;
    }
    board[r+1][c] = temp;
  }

  if (c+1 < n && board[r][c+1] == word[strIdx]) {
    char temp = board[r][c+1];
    board[r][c+1] = '.';
    if (dfs(board, m, n, r, c+1, word, ++strIdx)) {
      return true;
    }
    board[r][c+1] = temp;
  }

  if (c-1 >= 0 && board[r][c-1] == word[strIdx]) {
    char temp = board[r][c-1];
    board[r][c-1] = '.';
    if (dfs(board, m, n, r, c-1, word, ++strIdx)) {
      return true;
    }
    board[r][c-1] = temp;
  }

  return false;
}

class Solution {
public:
  bool exist(std::vector<std::vector<char>>& board, std::string word) {
    size_t m = board.size();
    size_t n = board[0].size();

    for (size_t i = 0; i < m; ++i) {
      for (size_t j = 0; j < n; ++j) {
        if (board[i][j] == word[0]) {
          char temp = board[i][j];
          board[i][j] = '.';
          if (dfs(board, m, n, i, j, word, 1)) {
            return true;
          }
          board[i][j] = temp;
        }
      }
    }

    return false;
  } 
};

int main() {
    Solution solution;
    std::vector<std::vector<char>> board = {
        {'A', 'B', 'C', 'E'},
        {'S', 'F', 'C', 'S'},
        {'A', 'D', 'E', 'E'}
    };
    std::string word = "ABCES";
    std::cout << solution.exist(board, word) << std::endl;
    return 0;
}
