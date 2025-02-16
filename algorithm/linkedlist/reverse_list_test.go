package linkedlist

import (
	"reflect"
	"strconv"
	"testing"
)

var reverseListTableTest = []struct {
	in   *ListNode
	want *ListNode
}{
	{
		in:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		want: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}},
	},
	{
		in:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
		want: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
	},
	{
		in:   nil,
		want: nil,
	},
}

func TestReserveList(t *testing.T) {
	for i, tt := range reverseListTableTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := reverseList(tt.in)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
