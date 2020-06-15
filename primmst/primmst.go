package primmst

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

func haveVert(edge *Edge, vidx int) bool {
	if edge.VertIdx[0] == vidx || edge.VertIdx[1] == vidx {
		return true
	}

	return false
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

	minheap := initHeap()
	for i := 0; i < len(vertices); i++ {
		if i == startIdx {
			continue
		}

		minCost, vEdges := maxInt, vertices[i].Edges
		for j := 0; j < len(vEdges); j++ {
			if haveVert(vEdges[j], startIdx) && vEdges[j].Cost < minCost {
				minCost = vEdges[j].Cost
			}
		}

		insert(minheap, heapNode{vert: &vertices[i], minCost: minCost})
	}

	for n := extractMin(minheap); n.vert != nil; n = extractMin(minheap) {
		mstcost += n.minCost

		for j, v := 0, n.vert; j < len(v.Edges); j++ {
			e := v.Edges[j]
			h := vertices[otherVert(e, v.VIdx)].HeapIdx
			if h != -1 {
				update(minheap, h, e.Cost)
			}
		}
	}

	return mstcost
}

type minHeap struct {
	heap      *[]heapNode
	lastEmpty *int
}

type heapNode struct {
	vert    *Vertex
	minCost int
}

func initHeap() minHeap {
	heap, lastEmpty := make([]heapNode, 0, 0), 0

	return minHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}
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

	bubbleUp(*lastEmpty, minheap)
}

func extractMin(minheap minHeap) heapNode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == 0 {
		return heapNode{vert: nil, minCost: maxInt}
	}

	ret := (*heap)[0]

	*lastEmpty--
	swapNode(heap, 0, *lastEmpty)

	bubbleDown(1, minheap)

	ret.vert.HeapIdx = -1 // bye
	return ret
}

func update(minheap minHeap, heapIdx int, newCost int) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	if newCost < (*heap)[heapIdx].minCost {
		(*heap)[heapIdx].minCost = newCost
		bubbleUp(heapIdx+1, minheap)
	}
}

func peekAt(minheap minHeap, heapIdx int) heapNode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		return heapNode{vert: nil, minCost: maxInt}
	}

	return (*heap)[heapIdx]
}

func peekMin(minheap minHeap) heapNode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == 0 {
		return heapNode{vert: nil, minCost: maxInt}
	}

	return (*heap)[0]
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

// pos is one based index
func bubbleUp(pos int, minheap minHeap) {
	heap := minheap.heap
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
