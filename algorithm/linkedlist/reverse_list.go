package linkedlist

// 206. Reverse Linked List
// Given the `head` of a singly linked list, reverse the list, and return the reversed list.
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode

	for head != nil {
		nextNode := head.Next
		head.Next = newHead
		newHead = head
		head = nextNode
	}

	return newHead
}
