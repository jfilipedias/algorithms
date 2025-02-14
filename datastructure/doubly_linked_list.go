package datastructure

type node[T any] struct {
	Value T
	Next  *node[T]
	Prev  *node[T]
}

type DoublyLinkedList[T any] struct {
	Head *node[T]
	Tail *node[T]
}

func (l *DoublyLinkedList[T]) AddToFront(value T) {
	node := &node[T]{Value: value, Next: nil, Prev: l.Head}
	if l.Head != nil {
		l.Head.Next = node
	} else {
		l.Tail = node
	}
	l.Head = node
}

func (l *DoublyLinkedList[T]) AddToBack(value T) {
	node := &node[T]{Value: value, Next: l.Tail, Prev: nil}
	if l.Tail != nil {
		l.Tail.Prev = node
	} else {
		l.Head = node
	}
	l.Tail = node
}

func (l *DoublyLinkedList[T]) RemoveFromFront() (T, bool) {
	var result T
	if l.Head == nil {
		return result, false
	}

	result = l.Head.Value
	l.Head = l.Head.Prev
	if l.Head != nil {
		l.Head.Next = nil
	} else {
		l.Tail = nil
	}

	return result, true
}

func (l *DoublyLinkedList[T]) RemoveFromBack() (T, bool) {
	var result T
	if l.Tail == nil {
		return result, false
	}

	result = l.Tail.Value
	l.Tail = l.Tail.Next
	if l.Tail != nil {
		l.Tail.Prev = nil
	} else {
		l.Head = nil
	}

	return result, true
}
