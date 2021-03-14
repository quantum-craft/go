package slice

type Queue struct {
	data []int
}

func (q *Queue) NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) PushBack(i int) {
	q.data = append(q.data, i)
}

func (q *Queue) PopFront() int {
	ret := q.data[0]
	q.data = q.data[1:]

	return ret
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}
