package linkedlist

import (
	"reflect"
	"strconv"
	"testing"
)

var middleNodeTestTable = []struct {
	in   *ListNode
	want *ListNode
}{
	{
		in:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
		want: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}},
	},
	{
		in:   &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}},
		want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}},
	},
}

func TestMiddleNode(t *testing.T) {
	for i, tt := range middleNodeTestTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := middleNode(tt.in)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
