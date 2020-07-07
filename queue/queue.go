package queue

import (
	"container/list"
)

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

type DataElement interface {
}

type queue struct {
	data *list.List
}

func MakeQueue() queue {
	return queue{
		data: list.New(),
	}
}

func GetMaxInt() int {
	return maxInt
}

func (q *queue) Enqueue(d DataElement) {
	q.data.PushBack(d)
}

func (q *queue) Dequeue() DataElement {
	if q.data.Len() == 0 {
		return maxInt // data being maxInt means empty
	}

	e := q.data.Front()
	q.data.Remove(e)

	return e.Value
}

func (q *queue) PeekFront() DataElement {
	if q.data.Len() == 0 {
		return maxInt // data being maxInt means empty
	}

	return q.data.Front().Value
}

func (q *queue) Peek2ndFront() DataElement {
	if q.data.Len() <= 1 {
		return maxInt // data being maxInt means empty
	}

	return q.data.Front().Next().Value
}

func (q *queue) QueueSize() int {
	return q.data.Len()
}
