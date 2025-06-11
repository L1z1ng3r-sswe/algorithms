/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}

	rear := reverse(mid(head))
	curr := head

	for curr.Next != nil {
		next := curr.Next
		curr.Next = rear
		rear = rear.Next
		curr.Next.Next = next
		curr = next
	}

	curr.Next = rear
}

func mid(head *ListNode) *ListNode {
	slow, fast := head, head
	prev := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = nil
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// time: O(n)