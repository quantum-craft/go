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

func TestMaxHeapSortLarge(t *testing.T) {
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

	numbers := make([]int, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	maxheap := NewMaxHeap()

	if maxheap.Size() != 0 {
		t.Error("HeapSort error !")
	}

	for i := 0; i < len(numbers); i++ {
		maxheap.Insert(node{Key: &Vertex{Score: numbers[i]}})

		if maxheap.Size() != i+1 {
			t.Error("HeapSort error !")
		}
	}

	ans := make([]int, 0)
	n, ok := maxheap.PopMax().(node)
	for ok == true {
		ans = append(ans, n.Key.Score)
		n, ok = maxheap.PopMax().(node)
	}

	if !SliceDecreasing(ans) {
		t.Error("HeapSort error !")
	}
}

// SliceDecresing tests if the slice is incremental by one
func SliceDecreasing(xs []int) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1]-xs[i] != -1 {
			return false
		}
	}

	return true
}
