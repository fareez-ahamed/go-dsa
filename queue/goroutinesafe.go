package queue

import "sync"

type SafeQueue[T any] struct {
	front *item[T]
	back  *item[T]
	mutex sync.RWMutex
}

func (q *SafeQueue[T]) Enqueue(val T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	item := item[T]{val, nil}
	if q.back == nil {
		q.front = &item
		q.back = &item
		return
	}
	q.back.next = &item
	q.back = q.back.next
}

func (q *SafeQueue[T]) Dequeue() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
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

func (q *SafeQueue[T]) Peek() (T, bool) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	var val T
	if q.front == nil {
		return val, false
	}
	val = q.front.val
	return val, true
}
