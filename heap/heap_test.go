package heap

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/quantum-craft/go/utils"
)

func TestMinHeap(t *testing.T) {
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

	heap := make([]int, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	for i := 0; i < len(numbers); i++ {
		InsertMinheap(minheap, numbers[i])
	}

	if GetMin(minheap) != 0 {
		t.Error("MinHeap GetMin error !")
	}

	ans := make([]int, 0)
	for n := ExtractMin(minheap); n != MaxInt; n = ExtractMin(minheap) {
		ans = append(ans, n)
	}

	if !utils.SliceIncreasing(ans) {
		t.Error("MinHeapSort error !")
	}
}

func TestMaxHeap(t *testing.T) {
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

	heap := make([]int, 0, 0)
	lastEmpty := 0
	maxheap := MaxHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	for i := 0; i < len(numbers); i++ {
		InsertMaxheap(maxheap, numbers[i])
	}

	if GetMax(maxheap) != 10000 {
		t.Error("MaxHeap GetMax error !")
	}

	ans := make([]int, 0)
	for n := ExtractMax(maxheap); n != MinInt; n = ExtractMax(maxheap) {
		ans = append(ans, n)
	}

	if !utils.SliceDecreasing(ans) {
		t.Error("MaxHeapSort error !")
	}
}
