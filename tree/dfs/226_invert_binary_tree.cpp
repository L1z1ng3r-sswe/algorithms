#include <iostream>

struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
};

class Solution {
 public:
  TreeNode* invertTree(TreeNode* root) {
    if (!root) {
      return nullptr;
    }

    TreeNode* tempLeft = root->left;
    root->left = root->right;
    root->right = tempLeft;

    invertTree(root->left);
    invertTree(root->right);

    return root;
  }
};