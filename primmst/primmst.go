package primmst

import (
	"github.com/quantum-craft/go/constant"
	heap "github.com/quantum-craft/go/minheap"
)

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

// GetCost implements the interface of min-heap
func (n heapNode) GetCost() int {
	return *n.minCost
}

// SetCost implements the interface of min-heap
func (n heapNode) SetCost(newCost int) {
	*n.minCost = newCost
}

// GetHeapIdx implements the interface of min-heap
func (n heapNode) GetHeapIdx() int {
	return n.vert.HeapIdx
}

// SetHeapIdx implements the interface of min-heap
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

		newCost := constant.MaxInt
		minheap.Insert(heapNode{vert: &vertices[i], minCost: &newCost})
	}

	v := vertices[startIdx]
	for j := 0; j < len(v.Edges); j++ {
		e := v.Edges[j]
		h := vertices[otherVert(e, v.VIdx)].HeapIdx

		if h != -1 {
			minheap.Update(h, e.Cost)
		}
	}

	n, ok := minheap.ExtractMin().(heapNode)
	for ok == true {
		mstcost += *n.minCost

		for j, v := 0, n.vert; j < len(v.Edges); j++ {
			e := v.Edges[j]
			h := vertices[otherVert(e, v.VIdx)].HeapIdx
			if h != -1 {
				minheap.Update(h, e.Cost)
			}
		}

		n, ok = minheap.ExtractMin().(heapNode)
	}

	return mstcost
}
