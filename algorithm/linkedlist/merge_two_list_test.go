package linkedlist

import (
	"reflect"
	"strconv"
	"testing"
)

var mergeTwoListsTestTable = []struct {
	list1 *ListNode
	list2 *ListNode
	want  *ListNode
}{
	{
		list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
		want:  &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}},
	},
	{
		list1: nil,
		list2: nil,
		want:  nil,
	},
	{
		list1: nil,
		list2: &ListNode{Val: 0, Next: nil},
		want:  &ListNode{Val: 0, Next: nil},
	},
	{
		list1: &ListNode{Val: 1, Next: nil},
		list2: &ListNode{Val: 2, Next: nil},
		want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
	},
	{
		list1: &ListNode{Val: 5, Next: nil},
		list2: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		want:  &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for i, tt := range mergeTwoListsTestTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
