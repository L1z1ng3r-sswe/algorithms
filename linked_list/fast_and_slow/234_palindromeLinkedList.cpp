struct ListNode {
  int val;
  ListNode* next;
  ListNode () : val(0), next(nullptr) {}
  ListNode (int val) : val(val), next(nullptr) {}
  ListNode (int val, ListNode* next) : val(val), next(next) {}
};
 
ListNode* mid(ListNode* head) {
  ListNode* fast = head;
  ListNode* slow = head;

  while (fast != nullptr && fast->next != nullptr) {
    fast = fast->next->next;
    slow = slow->next;
  }

  return slow;
}

ListNode* reverse(ListNode* head) {
  ListNode* prev = nullptr;

  while (head != nullptr) {
    ListNode* next = head->next;
    head->next = prev;
    prev = head;
    head = next;
  }

  return prev;
}

class Solution {
public:
    bool isPalindrome(ListNode* head) {
      ListNode* midNode = mid(head);
      ListNode* midNodeReversed = reverse(midNode);

      while (midNodeReversed != nullptr) {
        if (head->val != midNodeReversed->val) {
          return false;
        }

        midNodeReversed = midNodeReversed->next;
        head = head->next;
      }

      return true;
    }
};
