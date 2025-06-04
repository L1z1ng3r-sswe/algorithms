func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow := head.Next
	fast := head.Next.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

// time: O(n)

// n = x + l
// l = y + l-y
// slowD = x + y + a*l
// fastD = x + y + b*l
// fastD = slowD*2 ->
// x+y+b*l = 2(x+y+a*l) -> x+y+b*l = 2x+2y+2a*l -> b * l = x+y+2a*l ->
// x+y = (b-2a)*l -> x = (b-2a)*l - y
// x = y ('couse number of loops doesn't matter)