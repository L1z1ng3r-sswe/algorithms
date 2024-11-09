# Problem Description
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

 
### Example:
- Input: `root = [1,2,3,4,5]`
- Output: `3,4,5`

- Input: `toot = [1,2,3,4,5,6]`
- Output: `4,5,6`

## Initial Approach
fast and slow pointers

## Actual Approach
fast and slow pointers


## What I Didn't Think About The First Time
// i was confused, because i didn't know how to shit left pointer if ll is odd or even.
the answer - just choose the fast position properly.

## Solution (C++)
```cpp
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
```

## Complexity
n - number of nodes
time: O(n)
space: O(1)