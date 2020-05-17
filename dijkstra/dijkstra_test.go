package dijkstra

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

func TestDijkstra(t *testing.T) {
	file := "../data/dijkstraDataSmall.txt"

	vertices, edges := ConstructGraph(file)
	Dijkstra(vertices, edges, 0)
}

func TestMinHeap(t *testing.T) {
	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	Insert(minheap, Node{Key: &Vertex{Score: 88}})
	Insert(minheap, Node{Key: &Vertex{Score: 32}})
	Insert(minheap, Node{Key: &Vertex{Score: 50}})
	Insert(minheap, Node{Key: &Vertex{Score: 90}})
	Insert(minheap, Node{Key: &Vertex{Score: 23}})
	Insert(minheap, Node{Key: &Vertex{Score: 74}})
	Insert(minheap, Node{Key: &Vertex{Score: 2}})

	fmt.Println(ExtractMin(minheap).Key.Score)
	fmt.Println(ExtractMin(minheap).Key.Score)
	fmt.Println(ExtractMin(minheap).Key.Score)

	Insert(minheap, Node{Key: &Vertex{Score: 5}})
	Insert(minheap, Node{Key: &Vertex{Score: 91}})
	Insert(minheap, Node{Key: &Vertex{Score: 13}})
	Insert(minheap, Node{Key: &Vertex{Score: 1}})

	for n := ExtractMin(minheap); n.Key != nil; n = ExtractMin(minheap) {
		fmt.Println(n.Key.Score)
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
		Insert(minheap, Node{Key: &Vertex{Score: numbers[i]}})
	}

	ans := make([]int, 0)
	for n := ExtractMin(minheap); n.Key != nil; n = ExtractMin(minheap) {
		ans = append(ans, n.Key.Score)
	}

	if !utils.SliceIncreasing(ans) {
		t.Error("HeapSort error !")
	}
}
