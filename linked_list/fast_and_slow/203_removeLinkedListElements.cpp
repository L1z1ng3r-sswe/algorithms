class Solution {
public:
  ListNode* removeElements(ListNode* head, int val) {
    ListNode* dummy = new ListNode(0, head);
    ListNode* curr = dummy;

    while (curr != nullptr) {
      while (curr->next != nullptr && curr->next->val == val) {
        curr->next = curr->next->next;
      }

      curr = curr->next;
    }
    
    return dummy->next;
  }
};