package stack

type myStack[T any] struct {
	vals     []T
	length   int
	capacity int
}

func Constructor[T any]() myStack[T] {
	return myStack[T]{
		vals:     make([]T, 1),
		length:   0,
		capacity: 1,
	}
}

func (s *myStack[T]) Push(val T) {
	if s.length == s.capacity {
		s.resize()
	}

	s.vals[s.length] = val
	s.length++
}

func (s *myStack[T]) Pop() {
	if s.length == 0 {
		return
	}

	s.length--
}

func (s *myStack[T]) Peak() T {
	if s.length == 0 {
		var zero T
		return zero
	}

	return s.vals[s.length-1]
}

func (s *myStack[T]) Empty() bool {
	return s.length == 0
}

func (s *myStack[T]) Size() int {
	return s.length
}

func (s *myStack[T]) resize() {
	newVals := make([]T, s.capacity*2)
	copy(newVals, s.vals)
	s.vals = newVals

	s.capacity *= 2
}
