package queue

import (
	"container/list"
)

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

// DataElement is used as interface{}
type DataElement interface {
}

// Queue is the queue
type Queue struct {
	data *list.List
}

// MakeQueue creates a new queue
func MakeQueue() Queue {
	return Queue{
		data: list.New(),
	}
}

// GetMaxInt returns maxInt of queue package
func GetMaxInt() int {
	return maxInt
}

// PushBack enqueues from back
func (q *Queue) PushBack(d DataElement) {
	q.data.PushBack(d)
}

// PushFront enqueues from front
func (q *Queue) PushFront(d DataElement) {
	q.data.PushFront(d)
}

// PopFront dequeues from front
func (q *Queue) PopFront() DataElement {
	if q.data.Len() == 0 {
		return maxInt // data being maxInt means empty
	}

	e := q.data.Front()
	q.data.Remove(e)

	return e.Value
}

// PopBack dequeues from front
func (q *Queue) PopBack() DataElement {
	if q.data.Len() == 0 {
		return maxInt // data being maxInt means empty
	}

	e := q.data.Back()
	q.data.Remove(e)

	return e.Value
}

// PeekFront peeks the first element
func (q *Queue) PeekFront() DataElement {
	if q.data.Len() == 0 {
		return maxInt // data being maxInt means empty
	}

	return q.data.Front().Value
}

// Peek2ndFront peeks the second element
func (q *Queue) Peek2ndFront() DataElement {
	if q.data.Len() <= 1 {
		return maxInt // data being maxInt means empty
	}

	return q.data.Front().Next().Value
}

// QueueSize returns the size of queue
func (q *Queue) QueueSize() int {
	return q.data.Len()
}
