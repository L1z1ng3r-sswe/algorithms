class Solution {
public:
  bool isSameTree(TreeNode* p, TreeNode* q) {
    if (p == nullptr && q == nullptr) {
      return true;
    } 
    if (p == nullptr || q == nullptr) {
      return false;
    }

    std::vector<TreeNode*> queue1 = {p};
    std::vector<TreeNode*> queue2 = {q}; 

    while (!queue1.empty() && !queue2.empty()) {

      TreeNode* currNode1 = queue1.front();
      TreeNode* currNode2 = queue2.front();

      queue1.erase(queue1.begin());
      queue2.erase(queue2.begin());

      if (currNode1->val != currNode2->val) {
        return false;
      }

      if ((currNode1->left == nullptr) != (currNode2->left == nullptr)) {
        return false;
      }
      if (currNode1->left != nullptr) {
        queue1.push_back(currNode1->left);
        queue2.push_back(currNode2->left);
      }
        
      if ((currNode1->right == nullptr) != (currNode2->right == nullptr )) {
        return false;
      }
      if (currNode1->right != nullptr) {
        queue1.push_back(currNode1->right);
        queue2.push_back(currNode2->right);
      }
    }

    return queue1.empty() && queue2.empty();
  }
};