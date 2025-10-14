# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
      dummy = ListNode(0, head)
      prev = dummy
      curr = head
      while curr != None and curr.next != None:
        new_head = curr.next
        prev.next = new_head
        nxt = new_head.next
        new_head.next = curr
        curr.next = nxt
        prev = curr
        curr = nxt

      return dummy.next
    
# time: O(n)
# space: O(1)