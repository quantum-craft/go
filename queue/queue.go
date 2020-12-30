package queue

import (
	"container/list"
)

// DataElement is used as interface{}
type DataElement interface {
}

// Queue is the queue
type Queue struct {
	data *list.List
}

// NewQueue creates a new queue
func NewQueue() *Queue {
	return &Queue{
		data: list.New(),
	}
}

// Empty checks if the queue is empty
func (q *Queue) Empty() bool {
	if q.data.Len() == 0 {
		return true
	}

	return false
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
	if q.Empty() == true {
		return nil
	}

	e := q.data.Front()
	q.data.Remove(e)

	return e.Value
}

// PopBack dequeues from front
func (q *Queue) PopBack() DataElement {
	if q.Empty() == true {
		return nil
	}

	e := q.data.Back()
	q.data.Remove(e)

	return e.Value
}

// Front peeks the first element
func (q *Queue) Front() DataElement {
	if q.Empty() == true {
		return nil
	}

	return q.data.Front().Value
}

// SecondFront peeks the second element
func (q *Queue) SecondFront() DataElement {
	if q.data.Len() <= 1 {
		return nil
	}

	return q.data.Front().Next().Value
}

// QueueSize gives the size of underlying queue
func (q *Queue) QueueSize() int {
	return q.data.Len()
}
