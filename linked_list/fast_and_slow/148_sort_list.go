func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	second := divide(head)
	first := head

	second = sortList(second)
	first = sortList(first)

	return merge(first, second)
}

func divide(head *ListNode) *ListNode {
	var prev *ListNode
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = nil

	return slow
}

// merge merges l1 and l2 into a single node
func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return dummy.Next
}

// time: O(nLogn)