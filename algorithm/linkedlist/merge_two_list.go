package linkedlist

// 21. Merge Two Sorted Lists
// You are given the heads of two sorted linked lists `list1` and `list2`.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := list1
	for list1 != nil && list2 != nil {
		tmp := list2.Next

		if list2.Val < list1.Val {
			list2.Next = list1
			list1 = list2
			head = list1
		} else if list1.Next != nil && list2.Val <= list1.Next.Val {
			list2.Next = list1.Next
			list1.Next = list2
		}

		list2 = tmp
		list1 = list1.Next
	}

	return head
}
