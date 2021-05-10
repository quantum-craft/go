package segment

import (
	"math"
)

// Tree implementation of segment tree
type Tree struct {
	mins     []int //elements of the tree
	origSize int   //size number of elements in the original array
}

// NewTree constructs a segment tree and allows to perform RMQ on provided targetArray
func NewTree(input []int) *Tree {
	nodes := make([]int, calcTreeSize(len(input)))

	t := &Tree{nodes, len(input)}

	t.build(input, 0, 0, len(input)-1)

	return t
}

// RangeMinQuery returns minimum element in the [left,right] slice of the original array
func (t *Tree) RangeMinQuery(left, right int) int {
	if left > right {
		left, right = right, left
	}
	return (&query{left: left, right: right, mins: t.mins}).rangeMinimum(0, 0, t.origSize-1)
}

func (t *Tree) Add(pos int, x int) {
	t.addUtil(pos, x, 0, 0, t.origSize-1)
}

// Below are utils
func (t *Tree) build(input []int, curr int, leftBound, rightBound int) {
	if leftBound == rightBound {
		t.mins[curr] = input[leftBound]
		return
	}

	mid := (leftBound + rightBound) / 2
	t.build(input, 2*curr+1, leftBound, mid)
	t.build(input, 2*curr+2, mid+1, rightBound)

	t.mins[curr] = min(t.mins[2*curr+1], t.mins[2*curr+2])
}

func (q *query) rangeMinimum(curr int, leftBound, rightBound int) int {
	if q.left > rightBound || q.right < leftBound {
		return math.MaxInt32
	}

	if q.left <= leftBound && q.right >= rightBound {
		return q.mins[curr]
	}

	mid := (leftBound + rightBound) / 2
	return min(q.rangeMinimum(2*curr+1, leftBound, mid), q.rangeMinimum(2*curr+2, mid+1, rightBound))
}

func (t *Tree) addUtil(pos int, x int, curr int, leftBound, rightBound int) {
	if leftBound == rightBound {
		t.mins[curr] += x
	} else {
		mid := (leftBound + rightBound) / 2

		if pos <= mid {
			t.addUtil(pos, x, curr*2+1, leftBound, mid)
		} else {
			t.addUtil(pos, x, curr*2+2, mid+1, rightBound)
		}

		t.mins[curr] = min(t.mins[curr*2+1], t.mins[curr*2+2])
	}
}

type query struct {
	left, right int
	mins        []int
}

func calcTreeSize(originalSize int) int {
	return 1<<uint(math.Ceil(math.Log2(float64(originalSize)))+1) - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
