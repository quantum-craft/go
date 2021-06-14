package huffman

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type node struct {
	alphabet string
	weight   int
	left     *node
	right    *node
}

func EncodingWithQueue2(file string) *node {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := make([]*node, 0)

	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		if k > 0 {
			data = append(data, &node{
				alphabet: "",
				weight:   i,
				left:     nil,
				right:    nil,
			})
		}

		k++
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].weight < data[j].weight
	})

	queue := NewQueue()

	rp := 0
	var first, second *node
	for rp < len(data) {
		if queue.Len() == 0 {
			first, second = data[rp], data[rp+1]
			rp += 2
		} else if queue.Len() == 1 {
			first = queue.PopFront()
			second = data[rp]
			rp++
		} else {
			first = queue.PopFront()
			second = queue.PopFront()
		}

		newNode := &node{
			weight: first.weight + second.weight,
			left:   first,
			right:  second,
		}

		insert(newNode, data, &rp, queue)
	}

	for !queue.Empty() {
		if queue.Len() == 1 {
			ans := queue.PopFront()

			return ans
		}

		first := queue.PopFront()
		second := queue.PopFront()

		newNode := &node{
			weight: first.weight + second.weight,
			left:   first,
			right:  second,
		}

		queue.PushBack(newNode)
	}

	return nil
}

func insert(n *node, data []*node, rp *int, queue *Queue) {
	for *rp < len(data) && n.weight >= data[*rp].weight {
		queue.PushBack(data[*rp])
		*rp++
	}

	queue.PushBack(n)
}

func Iterate2(root *node) (min, max int) {
	if root.left == nil && root.right == nil {
		return 0, 0
	}

	min1, max1 := Iterate2(root.left)
	min2, max2 := Iterate2(root.right)

	max = maxOf(max1, max2) + 1
	min = minOf(min1, min2) + 1

	return min, max
}

type Queue struct {
	data []*node
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]*node, 0),
	}
}

func (q *Queue) PopFront() *node {
	ret := q.data[0]
	q.data = q.data[1:]

	return ret
}

func (q *Queue) PushBack(i *node) {
	q.data = append(q.data, i)
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) PeekFront() *node {
	return q.data[0]
}
