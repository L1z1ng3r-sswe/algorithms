# Problem Description
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

### Example:
- Input: `head = [3,2,0,-4], pos = 1`
- Output: `true`

- Input: `head = [1,2], pos = 0`
- Output: `true`

- Input: `head = [1], pos = -1`
- Output: `false`

## Initial Approach
fast and slow pointer

## Actual Approach
fast and slow pointer

## What I Didn't Think About The First Time
initially fast is the next val and it could be nill, and i checked this case, but i forgot to check the next of the fast, because if we will try to take next next which is nill - error.

## Solution (C++)
```cpp
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
```

## Complexity
n - number of nodes;
space: O(1)
time: O(n)