class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
      ListNode* dummy = new ListNode(0, head);
      ListNode* curr = dummy;

      while (curr != nullptr && curr->next != nullptr) {
        while (curr->next->next != nullptr && curr->next->val == curr->next->next->val) {
          curr->next = curr->next->next;
        }

        curr = curr->next;
      }

      return dummy->next;
    }
};