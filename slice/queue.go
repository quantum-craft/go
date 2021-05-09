package slice

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) PopFront() int {
	ret := q.data[0]
	q.data = q.data[1:]

	return ret
}

func (q *Queue) PushBack(i int) {
	q.data = append(q.data, i)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) PushFront(i int) {
	q.data = append([]int{i}, q.data...)
}

func (q *Queue) PeekFront() int {
	return q.data[0]
}

func (q *Queue) PopBack() int {
	ret := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]

	return ret
}

func (q *Queue) PeekBack() int {
	return q.data[len(q.data)-1]
}
