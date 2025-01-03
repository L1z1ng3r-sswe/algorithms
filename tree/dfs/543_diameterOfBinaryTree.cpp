// https://leetcode.com/problems/diameter-of-binary-tree/description/

#include <iostream>

struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;

  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int num, TreeNode* l) : val(num), left(l), right(nullptr) {}
  TreeNode(int num, TreeNode* l, TreeNode* r) : val(num), left(l), right(r) {}
};

int max(int a, int b) {
  if (a > b) {
    return a;
  }
  return b;
}

int dfs(TreeNode* root, int* maxDiameter) {
  if (root == nullptr) {
    return 0;
  }

  int left = dfs(root->left, maxDiameter);
  int right = dfs(root->right, maxDiameter);

  *maxDiameter = max(*maxDiameter, left + right);

  return max(left, right) + 1;
}

class Solution {
 public:
  int diameterOfBinaryTree(TreeNode* root) {
    int maxDiameter = 0;

    dfs(root, &maxDiameter);

    return maxDiameter;
  }
};