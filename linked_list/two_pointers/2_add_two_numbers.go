func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10 // can only be 0 or 1
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}

	return dummy.Next
}

// time: O(n)