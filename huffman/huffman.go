package huffman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	heap "github.com/quantum-craft/go/minheap"
	sort "github.com/quantum-craft/go/sorting"
	stack "github.com/quantum-craft/go/stack"
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

// Encoding2Stacks encodes the input alphabets into a Huffman tree
// Using 2 stacks to track 2 smallest members
func Encoding2Stacks(file string) HeapNode {
	f, _ := os.Open(file)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	data := []int{}

	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		if k > 0 {
			data = append(data, i)
		}

		k++
	}

	sort.QuickSort(data)

	minStack, dataStack := stack.MakeStack(), stack.MakeStack()
	for i := len(data) - 1; i >= 0; i-- {
		dataStack.Push(HeapNode{
			alphabet: "", // no need to store alphabet
			weight:   data[i],
			left:     nil,
			right:    nil,
		})
	}

	var root HeapNode
	for {
		n1 := popMin(&minStack, &dataStack)

		if dataStack.IsEmpty() && minStack.IsEmpty() {
			root = n1
			break
		}

		n2 := popMin(&minStack, &dataStack)

		newNode := HeapNode{
			alphabet: "",
			weight:   n1.weight + n2.weight,
			left:     &n1,
			right:    &n2,
		}

		second, ok := minStack.Peek2ndTop().(HeapNode)

		fmt.Println(minStack.GetData())

		if minStack.IsEmpty() || minStack.PeekTop().(HeapNode).weight > newNode.weight {
			minStack.Push(newNode)
		} else if !ok || second.weight > newNode.weight {
			d1 := minStack.Pop()
			minStack.Push(newNode)
			minStack.Push(d1)
		} else {
			d1 := minStack.Pop()
			d2 := minStack.Pop()
			minStack.Push(newNode)
			minStack.Push(d2)
			minStack.Push(d1)
		}
	}

	return root
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

func popMin(stack1 *stack.Stack, stack2 *stack.Stack) HeapNode {
	d1, ok1 := stack1.PeekTop().(HeapNode)
	if !ok1 {
		return stack2.Pop().(HeapNode)
	}

	d2, ok2 := stack2.PeekTop().(HeapNode)
	if !ok2 {
		return stack1.Pop().(HeapNode)
	}

	if d1.weight < d2.weight {
		return stack1.Pop().(HeapNode)
	}

	return stack2.Pop().(HeapNode)
}
