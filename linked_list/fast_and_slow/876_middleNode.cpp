class Solution {
  public:
    ListNode* middleNode(ListNode *root) {
      if (root == nullptr) {
        return nullptr;
      }

      ListNode *slow = root;
      ListNode *fast = root;  
      
      while (fast != nullptr && fast->next != nullptr) {
        fast = fast->next->next;
        slow = slow->next;
      }

      return slow;
    }
};  