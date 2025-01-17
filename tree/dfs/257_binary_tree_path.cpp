void dfs(TreeNode* root, std::vector<std::string>& res, std::string path) {
  if (!root->left && !root->right) {
    res.push_back(path);
  }

  if (root->left) {
    dfs(root->left, res, path + "->" + std::to_string(root->left->val));
  }

  if (root->right) {
    dfs(root->right, res, path + "->" + std::to_string(root->right->val));
  }
}

class Solution {
 public:
  std::vector<std::string> binaryTreePaths(TreeNode* root) {
    std::vector<std::string> res;
    dfs(root, res, std::to_string(root->val));
    return res;
  }
};