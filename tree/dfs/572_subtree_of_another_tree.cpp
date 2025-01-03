#include <iostream>

struct TreeNode {
  int val = 0;
  TreeNode *left = nullptr;
  TreeNode *right = nullptr;
};

bool areIdentical(TreeNode *root, TreeNode *subTree) {
  if (!root && !subTree) return true;
  if (!root || !subTree) return false;

  return root->val == subTree->val && areIdentical(root->left, subTree->left) &&
         areIdentical(root->right, subTree->right);
}

class Solution {
 public:
  bool isSubtree(TreeNode *root, TreeNode *subTree) {
    if (!subTree) return true;
    if (!root) return false;

    if (areIdentical(root, subTree)) {
      return true;
    }

    return isSubtree(root->left, subTree) || isSubtree(root->right, subTree);
  }
};