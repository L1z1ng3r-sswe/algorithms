# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head or not head.next:
            return head

        n, old_tail = list_node_len(head)

        k %= n
        if k == 0:
            return head

        dummy = ListNode(0, head)
        new_tail = dummy
        for _ in range(n - k):
            new_tail = new_tail.next

        new_head = new_tail.next
        new_tail.next = None
        old_tail.next = head

        return new_head


def list_node_len(head: Optional[ListNode]) -> tuple[int, Optional[ListNode]]:
    length = 0
    current = head

    while current and current.next:
        length += 1
        current = current.next
    length += 1

    return length, current

# time: O(n)
# time: O(1)