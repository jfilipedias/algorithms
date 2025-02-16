package linkedlist

func middleNode(head *ListNode) *ListNode {
	ahead := head

	for ahead != nil && ahead.Next != nil {
		ahead = ahead.Next.Next
		head = head.Next
	}

	return head
}
