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
	e := q.list.Front()
	q.list.Remove(e)

	return e.Value.(int)
}

func (q *queue) empty() bool {
	if q.list.Len() > 0 {
		return true
	}

	return false
}

func (q *queue) pushBack(d int) {
	q.list.PushBack(d)
}
