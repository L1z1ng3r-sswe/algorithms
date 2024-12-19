class Solution {
  public:
    bool hasCycle(ListNode *root) {
      if (root == nullptr) {
        return false;
      }

      ListNode *fast = root->next;
      ListNode *slow = root;

      while (fast != nullptr && fast->next != nullptr) {
        if (fast == slow) {
          return true;
        }

        fast = fast->next->next;
        slow = slow->next;
      }

      return false;
    }
};  