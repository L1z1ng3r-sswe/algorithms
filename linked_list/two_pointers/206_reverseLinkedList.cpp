#include <iostream>

struct ListNode {
  int val;
  ListNode *next;
  ListNode () : val(0), next(nullptr) {}
  ListNode (int val) : val(val), next(nullptr) {}
  ListNode (int val, ListNode* next) : val(val), next(next) {}
};

class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode* prev = nullptr;

        while (head != nullptr) {
          ListNode* next = head->next;
          head->next = prev; 
          prev = head;
          head = next;
        }

        return prev;
    }
};