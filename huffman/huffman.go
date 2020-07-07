package huffman

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"

	heap "github.com/quantum-craft/go/minheap"
	queue "github.com/quantum-craft/go/queue"
)

// HeapNode serves both purposes: node of min-heap and node of Huffman tree
type HeapNode struct {
	alphabet string
	weight   int
	left     *HeapNode
	right    *HeapNode
}

// GetCost implements interface
func (n HeapNode) GetCost() int {
	return n.weight
}

// SetCost implements interface
func (n HeapNode) SetCost(newCost int) {
	// Do nothing
}

// GetHeapIdx implements interface
func (n HeapNode) GetHeapIdx() int {
	// Useless info
	return -1
}

// SetHeapIdx implements interface
func (n HeapNode) SetHeapIdx(idx int) {
	// Do nothing
}

var maxNode HeapNode = HeapNode{
	alphabet: "",
	weight:   queue.GetMaxInt(),
	left:     nil,
	right:    nil,
}

// EncodingWithQueue encodes the input alphabets into a Huffman tree
// Using 1 queue to book-keeping
func EncodingWithQueue(file string) HeapNode {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := []HeapNode{}

	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		if k > 0 {
			data = append(data, HeapNode{
				alphabet: "",
				weight:   i,
				left:     nil,
				right:    nil,
			})
		}

		k++
	}

	quickSort(data)

	minQueue := queue.MakeQueue()
	minQueue.Enqueue(data[0])
	minQueue.Enqueue(data[1])

	k = 2
	var root HeapNode
	for {
		if k == len(data) && minQueue.QueueSize() == 1 {
			root = minQueue.Dequeue().(HeapNode)
			break
		}

		q1, ok1 := minQueue.PeekFront().(HeapNode)
		if !ok1 {
			q1 = maxNode
		}
		q2, ok2 := minQueue.Peek2ndFront().(HeapNode)
		if !ok2 {
			q2 = maxNode
		}

		d1, d2 := maxNode, maxNode
		if k < len(data)-1 {
			d1 = data[k]
			d2 = data[k+1]
		} else if k == len(data)-1 {
			d1 = data[k]
		}

		compareList := []HeapNode{q1, q2, d1, d2}
		minIdx := findMinIdx(compareList)
		n1 := compareList[minIdx]

		if minIdx < 2 {
			minQueue.Dequeue()
		} else {
			k++
		}

		minIdx = findMinIdxExcept(compareList, minIdx)
		n2 := compareList[minIdx]

		if minIdx < 2 {
			minQueue.Dequeue()
		} else {
			k++
		}

		newNode := HeapNode{
			alphabet: "",
			weight:   n1.weight + n2.weight,
			left:     &n1,
			right:    &n2,
		}

		minQueue.Enqueue(newNode)
	}

	return root
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

// QuickSort sorts array in-place with randomized choices of pivot
func quickSort(xs []HeapNode) {
	if len(xs) <= 1 {
		return
	}

	pivotPos := partition(xs, r.Intn(len(xs)))
	quickSort(xs[0:pivotPos])
	quickSort(xs[pivotPos+1:])
}

func partition(xs []HeapNode, pivotIdx int) int {
	if len(xs) <= 1 {
		return 0
	}

	swap(xs, 0, pivotIdx)

	i := 0
	for j := 1; j < len(xs); j++ {
		if xs[j].weight < xs[0].weight {
			swap(xs, i+1, j)
			i++
		}
	}

	swap(xs, 0, i)

	return i
}

func swap(xs []HeapNode, thisIdx int, thatIdx int) {
	xs[thisIdx], xs[thatIdx] = xs[thatIdx], xs[thisIdx]
}

// Encoding encodes the input alphabets into a Huffman tree
// Using min-heap to track 2 smallest members
func Encoding(file string) HeapNode {
	minheap := heap.MakeMinHeap()

	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		if k > 0 {
			heap.Insert(minheap, HeapNode{
				alphabet: string(k),
				weight:   i,
				left:     nil,
				right:    nil,
			})
		}

		k++
	}

	var root HeapNode
	for {
		n1, _ := heap.ExtractMin(minheap).(HeapNode)
		n2, haveTwo := heap.ExtractMin(minheap).(HeapNode)

		if haveTwo == false {
			root = n1
			break
		}

		heap.Insert(minheap, HeapNode{
			alphabet: "",
			weight:   n1.weight + n2.weight,
			left:     &n1,
			right:    &n2,
		})
	}

	return root
}

// Iterate will iterate through all nodes and will calculate the max and min length of paths
func Iterate(root *HeapNode) (min, max int) {
	if root.left == nil && root.right == nil {
		return 0, 0
	}

	min1, max1 := Iterate(root.left)
	min2, max2 := Iterate(root.right)

	max = maxOf(max1, max2) + 1
	min = minOf(min1, min2) + 1

	return min, max
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minOf(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func findMinIdx(xs []HeapNode) int {
	minIdx := -1
	min := queue.GetMaxInt()

	for i := 0; i < len(xs); i++ {
		if xs[i].weight < min {
			min = xs[i].weight
			minIdx = i
		}
	}

	return minIdx
}

func findMinIdxExcept(xs []HeapNode, except int) int {
	minIdx := -1
	min := queue.GetMaxInt()

	for i := 0; i < len(xs); i++ {
		if i != except && xs[i].weight < min {
			min = xs[i].weight
			minIdx = i
		}
	}

	return minIdx
}
