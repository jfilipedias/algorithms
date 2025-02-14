package datastructure

import "testing"

func TestAddToFront(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.AddToFront(10)
	if list.Head == nil || list.Head.Value != 10 {
		t.Errorf("want head with value of 10, but got %d", list.Head.Value)
	}
	if list.Tail == nil || list.Tail.Value != 10 {
		t.Errorf("want tail with value of 10, but got %d", list.Tail.Value)
	}

	list.AddToFront(20)
	if list.Head.Value != 20 {
		t.Errorf("want head with value of 20, but got %d", list.Head.Value)
	}
	if list.Head.Prev.Value != 10 {
		t.Errorf("want head previous value of 10, but got %d", list.Head.Prev.Value)
	}
}

func TestAddToBack(t *testing.T) {
	list := DoublyLinkedList[int]{}

	list.AddToBack(10)
	if list.Head == nil || list.Head.Value != 10 {
		t.Errorf("want head with value of 10, but got %d", list.Head.Value)
	}
	if list.Tail == nil || list.Tail.Value != 10 {
		t.Errorf("want tail with value of 10, but got %d", list.Tail.Value)
	}

	list.AddToBack(20)
	if list.Tail.Value != 20 {
		t.Errorf("want tail with value of 20, but got %d", list.Tail.Value)
	}
	if list.Tail.Next.Value != 10 {
		t.Errorf("want tail previous value of 10, but got %d", list.Tail.Next.Value)
	}
}

func TestRemoveFromFront(t *testing.T) {
	list := DoublyLinkedList[int]{}
	_, ok := list.RemoveFromFront()
	if ok {
		t.Errorf("should not be able to remove from a empty list")
	}

	list.AddToFront(10)
	list.AddToFront(20)

	got, ok := list.RemoveFromFront()
	if got != 20 || !ok {
		t.Errorf("want 20, but got %d", got)
	}
	if list.Head.Value != 10 {
		t.Errorf("want head with value of 10, but got %d", list.Head.Value)
	}
}

func TestRemoveFroBack(t *testing.T) {
	list := DoublyLinkedList[int]{}
	_, ok := list.RemoveFromBack()
	if ok {
		t.Errorf("should not be able to remove from a empty list")
	}

	list.AddToBack(10)
	list.AddToBack(20)

	got, ok := list.RemoveFromBack()
	if got != 20 || !ok {
		t.Errorf("want 20, but got %d", got)
	}
	if list.Head.Value != 10 {
		t.Errorf("want tail with value of 10, but got %d", list.Head.Value)
	}
}
