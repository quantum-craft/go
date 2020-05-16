package minheap

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
	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	Insert(minheap, Node{Key: nil, Score: 88})
	Insert(minheap, Node{Key: nil, Score: 32})
	Insert(minheap, Node{Key: nil, Score: 50})
	Insert(minheap, Node{Key: nil, Score: 90})
	Insert(minheap, Node{Key: nil, Score: 23})
	Insert(minheap, Node{Key: nil, Score: 74})
	Insert(minheap, Node{Key: nil, Score: 2})

	fmt.Println(ExtractMin(minheap))
	fmt.Println(ExtractMin(minheap))
	fmt.Println(ExtractMin(minheap))

	Insert(minheap, Node{Key: nil, Score: 5})
	Insert(minheap, Node{Key: nil, Score: 91})
	Insert(minheap, Node{Key: nil, Score: 13})
	Insert(minheap, Node{Key: nil, Score: 1})

	for n := ExtractMin(minheap); n.Score != -1; n = ExtractMin(minheap) {
		fmt.Println(n)
	}
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

	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	for i := 0; i < len(numbers); i++ {
		Insert(minheap, Node{Key: nil, Score: numbers[i]})
	}

	ans := make([]int, 0)
	for n := ExtractMin(minheap); n.Score != -1; n = ExtractMin(minheap) {
		ans = append(ans, n.Score)
	}

	if !utils.SliceIncreasing(ans) {
		t.Error("HeapSort error !")
	}
}
