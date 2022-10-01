package queue

// Front -> B -> C -> D  <- Back
// Enqueue
// Dequeue
// Peek

type item[T any] struct {
	val  T
	next *item[T]
}

type Queue[T any] struct {
	front *item[T]
	back  *item[T]
}

func (q *Queue[T]) Enqueue(val T) {
	item := item[T]{val, nil}
	if q.back == nil {
		q.front = &item
		q.back = &item
		return
	}
	q.back.next = &item
	q.back = q.back.next
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var val T
	if q.front == nil {
		return val, false
	}
	val = q.front.val
	q.front = q.front.next
	if q.front == nil {
		q.back = nil
	}
	return val, true
}

func (q *Queue[T]) Peek() (T, bool) {
	var val T
	if q.front == nil {
		return val, false
	}
	val = q.front.val
	return val, true
}
