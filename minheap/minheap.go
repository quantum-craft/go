package minheap

// HeapNode is the node unit in min-heap
type HeapNode interface {
	GetCost() int
	SetCost(int)
	GetHeapIdx() int
	SetHeapIdx(int)
}

// MinHeap keeps minimum element on the top and also keeps the heap property
type MinHeap struct {
	heap      *[]HeapNode
	lastEmpty *int
}

// MakeMinHeap returns an empty min-heap
func MakeMinHeap() MinHeap {
	return MinHeap{
		heap:      &[]HeapNode{},
		lastEmpty: new(int),
	}
}

// Insert will insert the node at the bottom and re-heapify
func (minheap *MinHeap) Insert(n HeapNode) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)

	minheap.bubbleUp(*lastEmpty)
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and re-heapify
func (minheap *MinHeap) ExtractMin() HeapNode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == 0 {
		var ret HeapNode
		return ret
	}

	ret := (*heap)[0]

	*lastEmpty--
	swapNode(heap, 0, *lastEmpty)

	minheap.bubbleDown(1)

	ret.SetHeapIdx(-1) // bye

	return ret
}

// ForceUpdate updates cost and re-heapify
func (minheap *MinHeap) ForceUpdate(heapIdx int, newCost int) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	if newCost < (*heap)[heapIdx].GetCost() {
		(*heap)[heapIdx].SetCost(newCost)

		// new cost is smaller, re-heapify
		minheap.bubbleUp(heapIdx + 1)
	} else {
		(*heap)[heapIdx].SetCost(newCost)

		// new cost is larger, re-heapify
		minheap.bubbleDown(heapIdx + 1)
	}
}

// Update updates cost when it is smaller
func (minheap *MinHeap) Update(heapIdx int, newCost int) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	if newCost < (*heap)[heapIdx].GetCost() {
		(*heap)[heapIdx].SetCost(newCost)

		// new cost is smaller, re-heapify
		minheap.bubbleUp(heapIdx + 1)
	}
}

// PeekAt provides the element at heapIdx
func (minheap *MinHeap) PeekAt(heapIdx int) HeapNode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		var ret HeapNode
		return ret
	}

	return (*heap)[heapIdx]
}

// PeekMin provides the min element without poping it
func (minheap *MinHeap) PeekMin() HeapNode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == 0 {
		var ret HeapNode
		return ret
	}

	return (*heap)[0]
}

// Delete will delete the element at heapIdx and replace it with the last element
// re-heapify
func (minheap *MinHeap) Delete(heapIdx int) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	*lastEmpty--
	swapNode(heap, heapIdx, *lastEmpty)

	if (*heap)[heapIdx].GetCost() > (*heap)[*lastEmpty].GetCost() {
		minheap.bubbleDown(heapIdx + 1)
	} else {
		minheap.bubbleUp(heapIdx + 1)
	}

	(*heap)[*lastEmpty].SetHeapIdx(-1) // bye
}

// pos is one based index
func (minheap *MinHeap) bubbleUp(pos int) {
	heap := minheap.heap
	for p := pos; p > 1 && (*heap)[p/2-1].GetCost() >= (*heap)[p-1].GetCost(); p = p / 2 {
		swapNode(heap, p-1, p/2-1)
	}
}

// pos is one based index
func (minheap *MinHeap) bubbleDown(pos int) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	p := pos
	for {
		if p-1 >= *lastEmpty || p*2-1 >= *lastEmpty {
			return
		}

		here := minheap.downHere(p)

		if here == p {
			return
		}

		swapNode(heap, p-1, here-1)

		p = here
	}
}

func (minheap *MinHeap) downHere(p int) int {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if p*2 >= *lastEmpty {
		return findMinPos2(heap, p, p*2)
	}

	return findMinPos3(heap, p, p*2, p*2+1)
}

// pos is one based
func findMinPos2(heap *[]HeapNode, pos1 int, pos2 int) int {
	if (*heap)[pos1-1].GetCost() < (*heap)[pos2-1].GetCost() {
		return pos1
	}

	return pos2
}

// pos is one based
func findMinPos3(heap *[]HeapNode, pos1 int, pos2 int, pos3 int) int {
	return findMinPos2(heap, findMinPos2(heap, pos1, pos2), pos3)
}

func swapNode(heap *[]HeapNode, this int, that int) {
	(*heap)[this], (*heap)[that] = (*heap)[that], (*heap)[this]
	(*heap)[this].SetHeapIdx(this)
	(*heap)[that].SetHeapIdx(that)
}
