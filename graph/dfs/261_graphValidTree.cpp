#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>

bool dfs(const std::unordered_map<int, std::vector<int>>& graph, std::unordered_set<int>& visited, int curr, int parent) {
  if (visited.find(curr) != visited.end()) {
    return false;
  }
  
  visited.insert(curr);

  for (const int& neighbor : graph.at(curr)) {
    if (neighbor != parent) {
      if (!dfs(graph, visited, neighbor, curr)) {
        return false;
      }
    }
  }
  
  return true;
}

bool isValid(int n, std::vector<std::vector<int>>& edges) {
  if (edges.size() + 1 != n) {
    return false;
  } 

  std::unordered_set<int> visited;
  visited.reserve(n);

  std::unordered_map<int, std::vector<int>> graph;
  graph.reserve(n);

  for (const std::vector<int>& edge : edges) {
    graph[edge[0]].push_back(edge[1]);
    graph[edge[1]].push_back(edge[0]);
  }

  if (edges.empty()) {
    return n == 1 || n == 0;
  }

  if (!dfs(graph, visited, edges[0][0], -1)) {
    return false;
  }

  // ensure all nodes are connected
  return n == visited.size();
}