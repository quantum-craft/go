package minheap

import "github.com/quantum-craft/go/dijkstra"

// Node is node unit for Dijkstra algorithm
type Node struct {
	Key   *dijkstra.Vertex
	Score int // Dijkstra's greedy score
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
	for pos := *lastEmpty; pos > 1 && (*heap)[pos/2-1].Score >= (*heap)[pos-1].Score; pos = pos / 2 {
		swapNode(heap, pos-1, pos/2-1)
	}
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and bubble down it to the proper position
func ExtractMin(minheap MinHeap) Node {
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == 0 {
		return Node{Key: nil, Score: -1}
	}

	var heap *[]Node = minheap.heap
	var ret Node = (*heap)[0]

	*lastEmpty--
	swapNode(heap, 0, *lastEmpty)

	// pos is one based
	pos := 1
	for {
		if pos-1 >= *lastEmpty || pos*2-1 >= *lastEmpty {
			return ret
		}

		minPos := -1
		if pos*2 >= *lastEmpty {
			minPos = findMinScorePos2(heap, pos, pos*2)
		} else {
			minPos = findMinScorePos3(heap, pos, pos*2, pos*2+1)
		}

		if minPos == pos {
			return ret
		}

		swapNode(heap, pos-1, minPos-1)

		pos = minPos
	}
}

// pos is one based
func findMinScorePos2(heap *[]Node, pos1 int, pos2 int) int {
	minPos := -1

	if (*heap)[pos1-1].Score < (*heap)[pos2-1].Score {
		minPos = pos1
	} else {
		minPos = pos2
	}

	return minPos
}

// pos is one based
func findMinScorePos3(heap *[]Node, pos1 int, pos2 int, pos3 int) int {
	minPos := -1

	if (*heap)[pos1-1].Score < (*heap)[pos2-1].Score {
		minPos = pos1
	} else {
		minPos = pos2
	}

	if (*heap)[pos3-1].Score < (*heap)[minPos-1].Score {
		minPos = pos3
	}

	return minPos
}

func swapNode(heap *[]Node, this int, that int) {
	n := (*heap)[this]
	(*heap)[this] = (*heap)[that]
	(*heap)[that] = n
}
