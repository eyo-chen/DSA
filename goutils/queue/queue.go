package queue

import (
	"github.com/OYE0303/DSA/goutils/linkedlist"
)

type myQueue[T any] struct {
	vals linkedlist.MyLinkedList[T]
}

func Constructor[T any]() myQueue[T] {
	return myQueue[T]{
		vals: linkedlist.Constructor[T](),
	}
}

func (q *myQueue[T]) Push(val T) {
	q.vals.InsertAtTail(val)
}

func (q *myQueue[T]) Pop() {
	q.vals.DeleteAtHead()
}

func (q *myQueue[T]) Front() T {
	return q.vals.GetValueAtHead()
}

func (q *myQueue[T]) Empty() bool {
	return q.vals.Empty()
}

func (q *myQueue[T]) Size() int {
	return q.vals.Size
}
