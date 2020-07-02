package primmst

import heap "github.com/quantum-craft/go/minheap"

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

// Vertex is used for undirected graph in Prim's mst algorithm
type Vertex struct {
	VIdx    int
	HeapIdx int
	Edges   []*Edge
}

// Edge is used for undirected graph in Prim's mst algorithm
type Edge struct {
	VertIdx [2]int
	Cost    int
}

type heapNode struct {
	vert    *Vertex
	minCost *int
}

func (n heapNode) GetCost() int {
	return *n.minCost
}

func (n heapNode) SetCost(newCost int) {
	*n.minCost = newCost
}

func (n heapNode) GetHeapIdx() int {
	return n.vert.HeapIdx
}

func (n heapNode) SetHeapIdx(idx int) {
	n.vert.HeapIdx = idx
}

func otherVert(edge *Edge, vidx int) int {
	if edge.VertIdx[0] == vidx {
		return edge.VertIdx[1]
	}

	return edge.VertIdx[0]
}

// PrimMST using Prim's minimum spanning tree algorithm to find mst cost
func PrimMST(vertices []Vertex, edges []Edge, startIdx int) int {
	mstcost := 0

	minheap := heap.MakeMinHeap()
	for i := 0; i < len(vertices); i++ {
		if i == startIdx {
			continue
		}

		newCost := maxInt
		heap.Insert(minheap, heapNode{vert: &vertices[i], minCost: &newCost})
	}

	v := vertices[startIdx]
	for j := 0; j < len(v.Edges); j++ {
		e := v.Edges[j]
		h := vertices[otherVert(e, v.VIdx)].HeapIdx

		if h != -1 {
			heap.Update(minheap, h, e.Cost)
		}
	}

	n, ok := heap.ExtractMin(minheap).(heapNode)
	for ok == true {
		mstcost += *n.minCost

		for j, v := 0, n.vert; j < len(v.Edges); j++ {
			e := v.Edges[j]
			h := vertices[otherVert(e, v.VIdx)].HeapIdx
			if h != -1 {
				heap.Update(minheap, h, e.Cost)
			}
		}

		n, ok = heap.ExtractMin(minheap).(heapNode)
	}

	return mstcost
}

// func delete(minheap minHeap, heapIdx int) {
// 	lastEmpty, heap := minheap.lastEmpty, minheap.heap

// 	if heapIdx >= *lastEmpty {
// 		return
// 	}

// 	*lastEmpty--
// 	swapNode(heap, heapIdx, *lastEmpty)

// 	if (*heap)[heapIdx].minCost > (*heap)[*lastEmpty].minCost {
// 		bubbleDown(heapIdx+1, minheap)
// 	} else {
// 		bubbleUp(heapIdx+1, minheap)
// 	}

// 	(*heap)[*lastEmpty].vert.HeapIdx = -1 // bye
// }
