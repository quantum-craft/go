package minheap

import "github.com/quantum-craft/go/dijkstra"

// Node is node unit for Dijkstra algorithm
type Node struct {
	Key   *dijkstra.Vertex
	Value int // Dijkstra's greedy score
}

// MinHeap keeps minimum member on the top and keeps the heap property
type MinHeap struct {
	lastEmpty *int // zero based
	heap      *[]Node
}

// Insert will insert the node at the bottom and bubble up to the proper position
func Insert(minheap MinHeap, n Node) {
	var heap *[]Node = minheap.heap
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	// pos is one based
	for pos := *lastEmpty; pos > 1 && (*heap)[pos/2-1].Value >= (*heap)[pos-1].Value; pos = pos / 2 {
		n := (*heap)[pos-1]
		(*heap)[pos-1] = (*heap)[pos/2-1]
		(*heap)[pos/2-1] = n
	}
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and bubble down it to the proper position
func ExtractMin(heap *[]Node) {

}
