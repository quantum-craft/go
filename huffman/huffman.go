package huffman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	heap "github.com/quantum-craft/go/minheap"
	queue "github.com/quantum-craft/go/queue"
	sort "github.com/quantum-craft/go/sort"
)

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

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

// ConcreteData is used as sort.Data
type ConcreteData struct {
	d []HeapNode
}

// Swap implements interface sort.Data
func (a ConcreteData) Swap(i, j int) {
	a.d[i], a.d[j] = a.d[j], a.d[i]
}

// Range implements interface sort.Data
func (a ConcreteData) Range(i, j int) sort.Data {
	return ConcreteData{d: a.d[i:j]}
}

// Len implements interface sort.Data
func (a ConcreteData) Len() int {
	return len(a.d)
}

// LessThan implements interface sort.Data
func (a ConcreteData) LessThan(i, j int) bool {
	if a.d[i].weight < a.d[j].weight {
		return true
	}

	return false
}

// Get implements interface sort.Data
func (a ConcreteData) Get(i int) interface{} {
	return a.d[i]
}

// Set implements interface sort.Data
func (a ConcreteData) Set(i int, data interface{}) {
	a.d[i] = data.(HeapNode)
}

var maxNode HeapNode = HeapNode{
	alphabet: "",
	weight:   maxInt,
	left:     nil,
	right:    nil,
}

// EncodingWithQueue encodes the input alphabets into a Huffman tree
// Using 1 queue to book-keeping
func EncodingWithQueue(file string) HeapNode {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := ConcreteData{d: []HeapNode{}}

	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		if k > 0 {
			data.d = append(data.d, HeapNode{
				alphabet: "",
				weight:   i,
				left:     nil,
				right:    nil,
			})
		}

		k++
	}

	sort.QuickSort(data)

	minQueue := queue.NewQueue()

	k = 0
	var root HeapNode
	for {
		if k == data.Len() && minQueue.QueueSize() == 1 {
			root = minQueue.PopFront().(HeapNode)
			break
		}

		q1, ok1 := minQueue.Front().(HeapNode)
		if !ok1 {
			q1 = maxNode
		}
		q2, ok2 := minQueue.SecondFront().(HeapNode)
		if !ok2 {
			q2 = maxNode
		}

		d1, d2 := maxNode, maxNode
		if k < data.Len()-1 {
			d1 = data.d[k]
			d2 = data.d[k+1]
		} else if k == data.Len()-1 {
			d1 = data.d[k]
		}

		compareList := []HeapNode{q1, q2, d1, d2}
		minIdx := findMinIdx(compareList)
		n1 := compareList[minIdx]

		if minIdx < 2 {
			minQueue.PopFront()
		} else {
			k++
		}

		minIdx = findMinIdxExcept(compareList, minIdx)
		n2 := compareList[minIdx]

		if minIdx < 2 {
			minQueue.PopFront()
		} else {
			k++
		}

		newNode := HeapNode{
			alphabet: "",
			weight:   n1.weight + n2.weight,
			left:     &n1,
			right:    &n2,
		}

		minQueue.PushBack(newNode)
	}

	return root
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

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
			minheap.Insert(HeapNode{
				alphabet: fmt.Sprint(k),
				weight:   i,
				left:     nil,
				right:    nil,
			})
		}

		k++
	}

	var root HeapNode
	for {
		n1, _ := minheap.ExtractMin().(HeapNode)
		n2, haveTwo := minheap.ExtractMin().(HeapNode)

		if haveTwo == false {
			root = n1
			break
		}

		minheap.Insert(HeapNode{
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
	min := maxInt

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
	min := maxInt

	for i := 0; i < len(xs); i++ {
		if i != except && xs[i].weight < min {
			min = xs[i].weight
			minIdx = i
		}
	}

	return minIdx
}
