package list

import "container/list"

// Queue is Queue
type Queue struct {
	list *list.List
}

// NewQueue is NewQueue
func NewQueue() *Queue {
	return &Queue{
		list: list.New(),
	}
}

// Front is Front
func (q *Queue) Front() int {
	return q.list.Front().Value.(int)
}

// Back is Back
func (q *Queue) Back() int {
	return q.list.Back().Value.(int)
}

// PopFront is PopFront
func (q *Queue) PopFront() {
	q.list.Remove(q.list.Front())
}

// FrontAndPop is FrontAndPop
func (q *Queue) FrontAndPop() int {
	e := q.list.Front()
	q.list.Remove(e)

	return e.Value.(int)
}

// PopBack is PopBack
func (q *Queue) PopBack() {
	q.list.Remove(q.list.Back())
}

// BackAndPop is BackAndPop
func (q *Queue) BackAndPop() int {
	e := q.list.Back()
	q.list.Remove(e)

	return e.Value.(int)
}

// Empty is Empty
func (q *Queue) Empty() bool {
	if q.list.Len() > 0 {
		return false
	}

	return true
}

// PushBack is PushBack
func (q *Queue) PushBack(d int) {
	q.list.PushBack(d)
}

// PushFront is PushFront
func (q *Queue) PushFront(d int) {
	q.list.PushFront(d)
}

// Size is Size
func (q *Queue) Size() int {
	return q.list.Len()
}
