package heap

// MaxHeap keeps maximum int on the top and also keeps the heap property
type MaxHeap struct {
	lastEmpty *int // zero based
	heap      *[]int
}

// InsertMaxheap will insert the node at the bottom and BUBBLE UP to the proper position
func InsertMaxheap(maxheap MaxHeap, n int) {
	var heap *[]int = maxheap.heap
	var lastEmpty *int = maxheap.lastEmpty

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)
	// pos is one based
	for pos := *lastEmpty; pos > 1 && (*heap)[pos/2-1] < (*heap)[pos-1]; pos = pos / 2 {
		swapNode(heap, pos-1, pos/2-1)
	}
}

// GetMax gives you the maximum member without extracting it
func GetMax(maxheap MaxHeap) int {
	var lastEmpty *int = maxheap.lastEmpty

	if *lastEmpty == 0 {
		return MaxInt
	}

	var heap *[]int = maxheap.heap
	return (*heap)[0]
}

// ExtractMax will extract the maximum member, replace the maximum pos with the last member,
// and BUBBLE DOWN it to the proper position
func ExtractMax(maxheap MaxHeap) int {
	var lastEmpty *int = maxheap.lastEmpty

	if *lastEmpty == 0 {
		return MaxInt
	}

	var heap *[]int = maxheap.heap
	var ret int = (*heap)[0]

	*lastEmpty--

	swapNode(heap, 0, *lastEmpty)
	// pos is one based
	pos := 1
	for {
		if pos-1 >= *lastEmpty || pos*2-1 >= *lastEmpty {
			return ret
		}

		maxPos := -1
		if pos*2 >= *lastEmpty {
			maxPos = findMaxScorePos2(heap, pos, pos*2)
		} else {
			maxPos = findMaxScorePos3(heap, pos, pos*2, pos*2+1)
		}

		if maxPos == pos {
			return ret
		}

		swapNode(heap, pos-1, maxPos-1)

		pos = maxPos
	}
}

// pos is one based
func findMaxScorePos2(heap *[]int, pos1 int, pos2 int) int {
	maxPos := -1

	if (*heap)[pos1-1] >= (*heap)[pos2-1] {
		maxPos = pos1
	} else {
		maxPos = pos2
	}

	return maxPos
}

// pos is one based
func findMaxScorePos3(heap *[]int, pos1 int, pos2 int, pos3 int) int {
	maxPos := -1

	if (*heap)[pos1-1] >= (*heap)[pos2-1] {
		maxPos = pos1
	} else {
		maxPos = pos2
	}

	if (*heap)[pos3-1] >= (*heap)[maxPos-1] {
		maxPos = pos3
	}

	return maxPos
}
