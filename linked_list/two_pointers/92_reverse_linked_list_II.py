# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
      dummy = ListNode(0, head)
      before_left_node = dummy

      for _ in range(left-1):
        before_left_node = before_left_node.next

      left_node = before_left_node.next
      prev = None
      for _ in range(right-left+1):
        nxt = left_node.next
        left_node.next = prev
        prev = left_node
        left_node = nxt

      before_left_node.next.next = left_node
      before_left_node.next = prev

      return dummy.next
     
# time: O(N)
# space: O(1)