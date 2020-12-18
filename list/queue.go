package list

import "container/list"

type queue struct {
	list *list.List
}

func newQueue() *queue {
	return &queue{
		list: list.New(),
	}
}

func (q *queue) front() int {
	return q.list.Front().Value.(int)
}

func (q *queue) back() int {
	return q.list.Back().Value.(int)
}

func (q *queue) popFront() {
	q.list.Remove(q.list.Front())
}

func (q *queue) frontAndPop() int {
	e := q.list.Front()
	q.list.Remove(e)

	return e.Value.(int)
}

func (q *queue) popBack() {
	q.list.Remove(q.list.Back())
}

func (q *queue) backAndPop() int {
	e := q.list.Back()
	q.list.Remove(e)

	return e.Value.(int)
}

func (q *queue) empty() bool {
	if q.list.Len() > 0 {
		return false
	}

	return true
}

func (q *queue) pushBack(d int) {
	q.list.PushBack(d)
}

func (q *queue) pushFront(d int) {
	q.list.PushFront(d)
}

func (q *queue) size() int {
	return q.list.Len()
}
