package heap

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx      int
	heapIdx  int
	Score    int // Dijkstra's greedy score
	Explored bool
	Edges    []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	Head   *Vertex
	Weight int
}

type node struct {
	Key *Vertex
}

func (n node) GetCost() int {
	return n.Key.Score
}

func (n node) SetCost(newCost int) {
	n.Key.Score = newCost
}

func (n node) GetHeapIdx() int {
	return n.Key.heapIdx
}

func (n node) SetHeapIdx(idx int) {
	n.Key.heapIdx = idx
}

func TestHeapSortLarge(t *testing.T) {
	lineEnding := "\n"

	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	} else {
		lineEnding = "\n"
	}

	f, err := os.Open("../data/QuickSortNumbers.txt")
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')

	numbers := make([]int, 0, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	minheap := NewMinHeap()

	for i := 0; i < len(numbers); i++ {
		minheap.Insert(node{Key: &Vertex{Score: numbers[i]}})
	}

	ans := make([]int, 0)
	n, ok := minheap.PopMin().(node)
	for ok == true {
		ans = append(ans, n.Key.Score)
		n, ok = minheap.PopMin().(node)
	}

	if !SliceIncreasing(ans) {
		t.Error("HeapSort error !")
	}
}

// SliceIncreasing tests if the slice is incremental by one
func SliceIncreasing(xs []int) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1]-xs[i] != 1 {
			return false
		}
	}

	return true
}
