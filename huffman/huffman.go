package huffman

import (
	"bufio"
	"os"
	"strconv"

	heap "github.com/quantum-craft/go/minheap"
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

// Encoding encodes the input alphabets into a Huffman tree
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
