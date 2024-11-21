#include <iostream>
#include <vector>
#include <unordered_map>

class Node {
public:
    int val;
    std::vector<Node*> neighbors;
    Node() {
        val = 0;
        neighbors = std::vector<Node*>();
    }
    Node(int _val) {
        val = _val;
        neighbors = std::vector<Node*>();
    }
    Node(int _val, std::vector<Node*> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

Node* dfs(Node* node, std::unordered_map<Node*, Node*>& visited) {
  if (visited.find(node) != visited.end()) {
    return visited.at(node);
  }
  
  Node* root = new Node(node->val);

  if (!node->neighbors.empty()) {
    root->neighbors.reserve(node->neighbors.size());
  }

  visited[node] = root;

  for (Node* neighbor : node->neighbors) {
    root->neighbors.push_back(dfs(neighbor, visited));
  }

  return root;
}

class Solution {
public:
  Node* cloneGraph(Node* node) {
    if (node == nullptr) {
      return nullptr;
    }

    std::unordered_map<Node*, Node*> visited;

    return dfs(node, visited);
  }
};