package primmst

import "fmt"

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

// Vertex is used for undirected graph in Prim's mst algorithm
type Vertex struct {
	HeapIdx int
}

// Edge is used for undirected graph in Prim's mst algorithm
type Edge struct {
	VertIdx [2]int
	Cost    int
}

// PrimMST using Prim's minimum spanning tree algorithm to find mst cost
func PrimMST(vertices []Vertex, edges []Edge) int {
	heap := make([]heapNode, 0, 0)
	lastEmpty := 0
	minheap := minHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	insert(minheap, heapNode{vert: &vertices[0], minCost: edges[0].Cost})
	insert(minheap, heapNode{vert: &vertices[1], minCost: edges[1].Cost})
	insert(minheap, heapNode{vert: &vertices[2], minCost: edges[2].Cost})
	insert(minheap, heapNode{vert: &vertices[3], minCost: edges[3].Cost})
	insert(minheap, heapNode{vert: &vertices[4], minCost: edges[4].Cost})

	n := extractMin(minheap)
	fmt.Println(n.minCost)
	fmt.Println(n.vert)
	n = extractMin(minheap)
	fmt.Println(n.minCost)
	fmt.Println(n.vert)
	n = extractMin(minheap)
	fmt.Println(n.minCost)
	fmt.Println(n.vert)
	n = extractMin(minheap)
	fmt.Println(n.minCost)
	fmt.Println(n.vert)
	n = extractMin(minheap)
	fmt.Println(n.minCost)
	fmt.Println(n.vert)
	n = extractMin(minheap)
	fmt.Println(n.minCost)
	fmt.Println(n.vert)

	return 0
}

type minHeap struct {
	heap      *[]heapNode
	lastEmpty *int
}

type heapNode struct {
	vert    *Vertex
	minCost int
}

func insert(minheap minHeap, n heapNode) {
	heap, lastEmpty := minheap.heap, minheap.lastEmpty

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)

	bubbleUp(*lastEmpty, heap)
}

func extractMin(minheap minHeap) heapNode {
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == 0 {
		return heapNode{vert: nil, minCost: maxInt}
	}

	heap := minheap.heap
	ret := (*heap)[0]

	*lastEmpty--
	swapNode(heap, 0, *lastEmpty)

	bubbleDown(1, minheap)

	ret.vert.HeapIdx = -1 // bye
	return ret
}

// pos is one based index
func bubbleUp(pos int, heap *[]heapNode) {
	for p := pos; p > 1 && (*heap)[p/2-1].minCost >= (*heap)[p-1].minCost; p = p / 2 {
		swapNode(heap, p-1, p/2-1)
	}
}

// pos is one based index
func bubbleDown(pos int, minheap minHeap) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	p := pos
	for {
		if p-1 >= *lastEmpty || p*2-1 >= *lastEmpty {
			return
		}

		here := downHere(p, minheap)

		if here == p {
			return
		}

		swapNode(heap, p-1, here-1)

		p = here
	}
}

func downHere(p int, minheap minHeap) int {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if p*2 >= *lastEmpty {
		return findMinPos2(heap, p, p*2)
	}

	return findMinPos3(heap, p, p*2, p*2+1)
}

func findMinPos2(heap *[]heapNode, pos1 int, pos2 int) int {
	if (*heap)[pos1-1].minCost < (*heap)[pos2-1].minCost {
		return pos1
	}

	return pos2
}

func findMinPos3(heap *[]heapNode, pos1 int, pos2 int, pos3 int) int {
	return findMinPos2(heap, findMinPos2(heap, pos1, pos2), pos3)
}

func swapNode(heap *[]heapNode, this int, that int) {
	(*heap)[this], (*heap)[that] = (*heap)[that], (*heap)[this]
	(*heap)[this].vert.HeapIdx = this
	(*heap)[that].vert.HeapIdx = that
}
