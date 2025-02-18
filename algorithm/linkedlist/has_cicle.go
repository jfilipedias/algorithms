package linkedlist

import "reflect"

// 141. Linked List Cycle
// Given `head`, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer.
// Return `true` if there is a cycle in the linked list. Otherwise, return `false`.
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if reflect.DeepEqual(slow, fast) {
			return true
		}
	}

	return false
}
