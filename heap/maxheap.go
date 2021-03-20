package heap

// MaxHeap keeps maximum element on the top and also keeps the heap property
type MaxHeap struct {
	heap      *[]INode
	lastEmpty *int
}

// NewMaxHeap returns an empty max-heap
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap:      &[]INode{},
		lastEmpty: new(int),
	}
}

// Empty returns whether heap is empty
func (maxheap *MaxHeap) Empty() bool {
	lastEmpty := maxheap.lastEmpty

	return *lastEmpty == 0
}

// Size returns the heap's size
func (maxheap *MaxHeap) Size() int {
	lastEmpty := maxheap.lastEmpty
	return *lastEmpty
}

// Insert will insert the node at the bottom and re-heapify
func (maxheap *MaxHeap) Insert(n INode) {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)

	maxheap.bubbleUp(*lastEmpty)
}

// PopMax will pop the maximum member, replace the maximum pos with the last member,
// and re-heapify
func (maxheap *MaxHeap) PopMax() INode {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if *lastEmpty == 0 {
		var ret INode
		return ret
	}

	ret := (*heap)[0]

	*lastEmpty--
	swapNode(heap, 0, *lastEmpty)

	maxheap.bubbleDown(1)

	ret.SetHeapIdx(-1) // bye

	return ret
}

// ForceUpdate updates cost and re-heapify
func (maxheap *MaxHeap) ForceUpdate(heapIdx int, newCost int) {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	if newCost > (*heap)[heapIdx].GetCost() {
		(*heap)[heapIdx].SetCost(newCost)

		// new cost is smaller, re-heapify
		maxheap.bubbleUp(heapIdx + 1)
	} else {
		(*heap)[heapIdx].SetCost(newCost)

		// new cost is larger, re-heapify
		maxheap.bubbleDown(heapIdx + 1)
	}
}

// Update updates cost when it is smaller
func (maxheap *MaxHeap) Update(heapIdx int, newCost int) {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	if newCost > (*heap)[heapIdx].GetCost() {
		(*heap)[heapIdx].SetCost(newCost)

		// new cost is smaller, re-heapify
		maxheap.bubbleUp(heapIdx + 1)
	}
}

// PeekAt provides the element at heapIdx
func (maxheap *MaxHeap) PeekAt(heapIdx int) INode {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if heapIdx >= *lastEmpty {
		var ret INode
		return ret
	}

	return (*heap)[heapIdx]
}

// PeekMax provides the min element without poping it
func (maxheap *MaxHeap) PeekMax() INode {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if *lastEmpty == 0 {
		var ret INode
		return ret
	}

	return (*heap)[0]
}

// Delete will delete the element at heapIdx and replace it with the last element
// re-heapify
func (maxheap *MaxHeap) Delete(heapIdx int) {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if heapIdx >= *lastEmpty {
		return
	}

	*lastEmpty--
	swapNode(heap, heapIdx, *lastEmpty)

	if (*heap)[heapIdx].GetCost() < (*heap)[*lastEmpty].GetCost() {
		maxheap.bubbleDown(heapIdx + 1)
	} else {
		maxheap.bubbleUp(heapIdx + 1)
	}

	(*heap)[*lastEmpty].SetHeapIdx(-1) // bye
}

// pos is one based index
func (maxheap *MaxHeap) bubbleUp(pos int) {
	heap := maxheap.heap
	for p := pos; p > 1 && (*heap)[p/2-1].GetCost() <= (*heap)[p-1].GetCost(); p = p / 2 {
		swapNode(heap, p-1, p/2-1)
	}
}

// pos is one based index
func (maxheap *MaxHeap) bubbleDown(pos int) {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	p := pos
	for {
		if p-1 >= *lastEmpty || p*2-1 >= *lastEmpty {
			return
		}

		here := maxheap.downHere(p)

		if here == p {
			return
		}

		swapNode(heap, p-1, here-1)

		p = here
	}
}

func (maxheap *MaxHeap) downHere(p int) int {
	lastEmpty, heap := maxheap.lastEmpty, maxheap.heap

	if p*2 >= *lastEmpty {
		return findMaxPos2(heap, p, p*2)
	}

	return findMaxPos3(heap, p, p*2, p*2+1)
}

// pos is one based
func findMaxPos2(heap *[]INode, pos1 int, pos2 int) int {
	if (*heap)[pos1-1].GetCost() > (*heap)[pos2-1].GetCost() {
		return pos1
	}

	return pos2
}

// pos is one based
func findMaxPos3(heap *[]INode, pos1 int, pos2 int, pos3 int) int {
	return findMaxPos2(heap, findMaxPos2(heap, pos1, pos2), pos3)
}
