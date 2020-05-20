package heap

// MinHeap keeps minimum int on the top and also keeps the heap property
type MinHeap struct {
	lastEmpty *int // zero based
	heap      *[]int
}

// MaxUint is the max uint
const MaxUint = ^int(0) // 11111111
// MinUint is the min uint
const MinUint = int(0) // 00000000
// MaxInt it the max int
const MaxInt = int(MaxUint >> 1) // 01111111
// MinInt is the min int
const MinInt = -MaxInt - 1 // 10000000

// InsertMinheap will insert the node at the bottom and BUBBLE UP to the proper position
func InsertMinheap(minheap MinHeap, n int) {
	var heap *[]int = minheap.heap
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)
	// pos is one based
	for pos := *lastEmpty; pos > 1 && (*heap)[pos/2-1] >= (*heap)[pos-1]; pos = pos / 2 {
		swapNode(heap, pos-1, pos/2-1)
	}
}

// GetMin gives you the minimum member without extracting it
func GetMin(minheap MinHeap) int {
	var heap *[]int = minheap.heap
	return (*heap)[0]
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and BUBBLE DOWN it to the proper position
func ExtractMin(minheap MinHeap) int {
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == 0 {
		return MaxInt // ExtractMin and get MaxInt => wrong
	}

	var heap *[]int = minheap.heap
	var ret int = (*heap)[0]

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
func findMinScorePos2(heap *[]int, pos1 int, pos2 int) int {
	minPos := -1

	if (*heap)[pos1-1] < (*heap)[pos2-1] {
		minPos = pos1
	} else {
		minPos = pos2
	}

	return minPos
}

// pos is one based
func findMinScorePos3(heap *[]int, pos1 int, pos2 int, pos3 int) int {
	minPos := -1

	if (*heap)[pos1-1] < (*heap)[pos2-1] {
		minPos = pos1
	} else {
		minPos = pos2
	}

	if (*heap)[pos3-1] < (*heap)[minPos-1] {
		minPos = pos3
	}

	return minPos
}

func swapNode(heap *[]int, this int, that int) {
	(*heap)[this], (*heap)[that] = (*heap)[that], (*heap)[this]
	// (*heap)[this].Key.heapIdx = this
	// (*heap)[that].Key.heapIdx = that
}
