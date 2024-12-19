
class Solution {
public:
  int minDepth(TreeNode* root) {
    if (root == nullptr) {
      return 0;
    }

    std::vector<TreeNode*> queue = {root};
    int depth = 0;
    while (!queue.empty()) {
      int currLen = queue.size();
      depth++;

      for (int i = 0; i < currLen; ++i) {
        TreeNode* currNode = queue[i];

        if (currNode->left == nullptr && currNode->right == nullptr) {
          return depth;
        }

        if (currNode->left != nullptr) {
          queue.push_back(currNode->left);
        }

        if (currNode->right != nullptr) {
          queue.push_back(currNode->right);
        }
      }

      queue.erase(queue.begin(), queue.begin()+currLen);
    }

    return depth;
  }
};