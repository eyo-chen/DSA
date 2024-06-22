package linkedlist

type myListNode[T any] struct {
	Val  T
	Next *myListNode[T]
}

type MyLinkedList[T any] struct {
	Head *myListNode[T]
	Size int
}

func Constructor[T any]() MyLinkedList[T] {
	return MyLinkedList[T]{
		Head: nil,
		Size: 0,
	}
}

func (l *MyLinkedList[T]) InsertAtHead(value T) {
	if l.Head == nil {
		l.Head = &myListNode[T]{
			Val:  value,
			Next: nil,
		}
	} else {
		node := &myListNode[T]{
			Val:  value,
			Next: l.Head,
		}
		l.Head = node
	}

	l.Size++
}

func (l *MyLinkedList[T]) InsertAtTail(value T) {
	node := &myListNode[T]{
		Val:  value,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = node
	} else {
		curNode := l.Head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = node
	}

	l.Size++
}

func (l *MyLinkedList[T]) DeleteAtHead() {
	if l.Empty() {
		return
	}

	newHead := l.Head.Next
	l.Head = newHead

	l.Size--
}

func (l *MyLinkedList[T]) DeleteAtTail() {
	if l.Empty() {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
	} else {
		curNode := l.Head
		for curNode.Next.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = nil
	}

	l.Size--
}

func (l *MyLinkedList[T]) GetValueAtHead() T {
	if l.Empty() {
		var zero T
		return zero
	}

	return l.Head.Val
}

func (l *MyLinkedList[T]) GetValueAtTail() T {
	if l.Empty() {
		var zero T
		return zero
	}

	curNode := l.Head
	for curNode.Next != nil {
		curNode = curNode.Next
	}

	return curNode.Val
}

func (l *MyLinkedList[T]) Empty() bool {
	return l.Size == 0
}
