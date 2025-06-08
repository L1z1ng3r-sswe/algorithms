func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, reverse(head)}
	curr := dummy

	for curr.Next != nil {
		n--

		if n == 0 {
			curr.Next = curr.Next.Next
			break
		}

		curr = curr.Next
	}

	return reverse(dummy.Next)
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