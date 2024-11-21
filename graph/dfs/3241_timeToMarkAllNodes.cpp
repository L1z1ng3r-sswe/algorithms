#include <iostream>
#include <vector>

int dfs(const std::vector<std::vector<int>>& graph, int curr, int sum, int parent) {
  if (parent != -1) {
    if (curr % 2 == 0) {
      sum += 2;
    } else {
      sum += 1;
    }
  }

  int maxVal = sum;
  for (const int& node : graph[curr]) {
    if (node != parent) {
      maxVal = std::max(maxVal, dfs(graph, node, sum, curr));
    }
  }

  return maxVal;
}

class Solution {
public:
  std::vector<int> timeTaken(std::vector<std::vector<int>>& edges) {
    size_t n = edges.size();

    std::vector<std::vector<int>> graph(n+1);
    for (const std::vector<int>& edge : edges) {
      int node1 = edge[0];
      int node2 = edge[1];
      graph[node1].push_back(node2);
      graph[node2].push_back(node1);
    }

    std::vector<int> res(n+1);
    for (int i = 0; i <= n; i++) {
      res[i] = dfs(graph, i, 0, -1);
    }

    return res;
  }
};