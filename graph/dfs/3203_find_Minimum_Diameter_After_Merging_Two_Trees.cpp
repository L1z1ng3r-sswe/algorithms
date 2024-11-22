#include <iostream>
#include <vector>
#include <algorithm>

int dfs(const std::vector<std::vector<int>>& graph, int curr, int parent, int* diameter) {
  int diameter1 = 0, diameter2 = 0;

  for (const int& neighbor : graph[curr]) {
    if (neighbor != parent) {
      int depth = dfs(graph, neighbor, curr, diameter);
      if (depth > diameter1) {
        diameter2 = diameter1;
        diameter1 = depth;
      } else if (depth > diameter2) {
        diameter2 = depth;
      }
    } 
  }

  if (*diameter < diameter1 + diameter2) {
    *diameter = diameter1 + diameter2;
  }

  return std::max(diameter1, diameter2) + 1;
}

std::vector<std::vector<int>> combineGraph(const std::vector<std::vector<int>>& edges, size_t n) {
  std::vector<std::vector<int>> graph (n);
  for (const std::vector<int>& edge : edges) {
    int node1 = edge[0], node2 = edge[1];

    graph[node1].push_back(node2);
    graph[node2].push_back(node1);
  }

  return graph;
}

class Solution {
public:
    int minimumDiameterAfterMerge(std::vector<std::vector<int>>& edges1, std::vector<std::vector<int>>& edges2) {
      size_t n = edges1.size()+1, m = edges2.size()+1;

      std::vector<std::vector<int>> graph1 = combineGraph(edges1, n), graph2 = combineGraph(edges2, m);
      
      int diameter1 = 0, diameter2 = 0;

      dfs(graph1, 0, -1, &diameter1);
      dfs(graph2, 0, -1, &diameter2);

      int combinedDiameter = (diameter1 + 1) / 2 + (diameter2 + 1) / 2 + 1;
      
      return std::max({diameter1, diameter2, combinedDiameter});
    }
};