package stack

type item[T any] struct {
	val  T
	next *item[T]
}

type Stack[T any] struct {
	top *item[T]
}

func (s *Stack[T]) Push(val T) {
	newItem := item[T]{
		val:  val,
		next: s.top,
	}
	s.top = &newItem
}

func (s *Stack[T]) Peek() (T, bool) {
	var val T
	if s.top == nil {
		return val, false
	}
	return s.top.val, true
}

func (s *Stack[T]) Pop() (T, bool) {
	var val T
	if s.top == nil {
		return val, false
	}
	val = s.top.val
	s.top = s.top.next
	return val, true
}
