package johnson

import (
	"math"

	"github.com/quantum-craft/go/bellmanford"
)

func Johnson(n int, m int, edges [][3]int) int {
	N := n + 1

	for h := 0; h < n; h++ {
		edges = append(edges, [3]int{n, h, 0})
	}

	// n is the virtual vertex
	dist := bellmanford.BellmanfordDist(N, edges, n)

	if dist == nil {
		return math.MaxInt32
	}

	edges = edges[:m]
	graph := make(map[int][][2]int)

	for i := range edges {
		edges[i][2] += (dist[edges[i][0]] - dist[edges[i][1]])
		graph[edges[i][0]] = append(graph[edges[i][0]], [2]int{edges[i][1], edges[i][2]})
	}

	ans := math.MaxInt32
	for src := 0; src < n; src++ {
		d, _ := Dijkstra(graph, src, n, dist)
		ans = min(ans, d)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Dijkstra(graph map[int][][2]int, src int, n int, dist []int) (int, int) {
	keeper := make([]*node, n)

	for i := range keeper {
		keeper[i] = &node{
			idx:     i,
			cost:    math.MaxInt32,
			heapIdx: -1,
			popped:  false,
		}
	}

	minheap := NewMinHeap()

	keeper[src].idx = src
	keeper[src].cost = 0
	keeper[src].heapIdx = -1
	keeper[src].popped = false

	minheap.Insert(keeper[src])

	minIdx := -1
	minDist := math.MaxInt32

	for !minheap.Empty() {

		curr := minheap.PopMin().(*node)
		curr.popped = true

		adjCost := curr.cost - (dist[src] - dist[curr.idx])
		if adjCost < minDist {
			minDist = adjCost
			minIdx = curr.idx
		}

		for _, head := range graph[curr.idx] {
			v := head[0]
			w := head[1]

			if !keeper[v].popped {
				if keeper[v].heapIdx == -1 {
					keeper[v].idx = v
					keeper[v].cost = curr.cost + w
					keeper[v].heapIdx = -1
					keeper[v].popped = false

					minheap.Insert(keeper[v])
				} else {
					minheap.UpdateIfSmaller(keeper[v].heapIdx, curr.cost+w)
				}
			}
		}
	}

	return minDist, minIdx
}

type node struct {
	idx     int
	cost    int
	heapIdx int
	popped  bool
}

func (n *node) GetCost() int {
	return n.cost
}

func (n *node) SetCost(c int) {
	n.cost = c
}

func (n *node) GetHeapIdx() int {
	return n.heapIdx
}

func (n *node) SetHeapIdx(i int) {
	n.heapIdx = i
}

// INode is the node interface in min-heap
type INode interface {
	GetCost() int
	SetCost(int)
	GetHeapIdx() int
	SetHeapIdx(int)
}

// MinHeap keeps minimum element on the top and also keeps the heap property
type MinHeap struct {
	heap      *[]INode
	lastEmpty *int
}

// NewMinHeap returns an empty min-heap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		heap:      &[]INode{},
		lastEmpty: new(int),
	}
}

// Empty returns whether heap is empty
func (minheap *MinHeap) Empty() bool {
	lastEmpty := minheap.lastEmpty

	return *lastEmpty == 0
}

// Size returns the heap's size
func (minheap *MinHeap) Size() int {
	lastEmpty := minheap.lastEmpty
	return *lastEmpty
}

// Insert will insert the node at the bottom and re-heapify
func (minheap *MinHeap) Insert(n INode) {
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

// PopMin will pop the minimum member, replace the minimum pos with the last member,
// and re-heapify
func (minheap *MinHeap) PopMin() INode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == 0 {
		var ret INode
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

// UpdateIfSmaller updates cost when it is smaller
func (minheap *MinHeap) UpdateIfSmaller(heapIdx int, newCost int) {
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
func (minheap *MinHeap) PeekAt(heapIdx int) INode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if heapIdx >= *lastEmpty {
		var ret INode
		return ret
	}

	return (*heap)[heapIdx]
}

// PeekMin provides the min element without poping it
func (minheap *MinHeap) PeekMin() INode {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == 0 {
		var ret INode
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
func findMinPos2(heap *[]INode, pos1 int, pos2 int) int {
	if (*heap)[pos1-1].GetCost() < (*heap)[pos2-1].GetCost() {
		return pos1
	}

	return pos2
}

// pos is one based
func findMinPos3(heap *[]INode, pos1 int, pos2 int, pos3 int) int {
	return findMinPos2(heap, findMinPos2(heap, pos1, pos2), pos3)
}

func swapNode(heap *[]INode, this int, that int) {
	(*heap)[this], (*heap)[that] = (*heap)[that], (*heap)[this]
	(*heap)[this].SetHeapIdx(this)
	(*heap)[that].SetHeapIdx(that)
}
