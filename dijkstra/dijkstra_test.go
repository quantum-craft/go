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
	file := "../data/dijkstraData.txt"

	vertices, edges := ConstructGraph(file)
	Dijkstra(vertices, edges, 0)

	if vertices[6].Score != 2599 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[36].Score != 2610 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[58].Score != 2947 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[81].Score != 2052 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[98].Score != 2367 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[114].Score != 2399 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[132].Score != 2029 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[164].Score != 2442 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[187].Score != 2505 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[196].Score != 3068 {
		t.Error("Dijkstra algorithm has error !")
	}
}

func TestMinHeap(t *testing.T) {
	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	vertex1 := &Vertex{idx: 0, heapIdx: -1, Score: 88}
	vertex2 := &Vertex{idx: 1, heapIdx: -1, Score: 32}
	vertex3 := &Vertex{idx: 2, heapIdx: -1, Score: 50}
	vertex4 := &Vertex{idx: 3, heapIdx: -1, Score: 90}
	vertex5 := &Vertex{idx: 4, heapIdx: -1, Score: 23}
	vertex6 := &Vertex{idx: 5, heapIdx: -1, Score: 74}
	vertex7 := &Vertex{idx: 6, heapIdx: -1, Score: 2}
	vertex8 := &Vertex{idx: 7, heapIdx: -1, Score: 5}
	vertex9 := &Vertex{idx: 8, heapIdx: -1, Score: 91}
	vertex10 := &Vertex{idx: 9, heapIdx: -1, Score: 13}
	vertex11 := &Vertex{idx: 10, heapIdx: -1, Score: 1}

	Insert(minheap, Node{Key: vertex1})
	Insert(minheap, Node{Key: vertex2})
	Insert(minheap, Node{Key: vertex3})
	Insert(minheap, Node{Key: vertex4})
	Insert(minheap, Node{Key: vertex5})
	Insert(minheap, Node{Key: vertex6})
	Insert(minheap, Node{Key: vertex7})
	Insert(minheap, Node{Key: vertex8})
	Insert(minheap, Node{Key: vertex9})
	Insert(minheap, Node{Key: vertex10})
	Insert(minheap, Node{Key: vertex11})

	FindKeyUpdateScore(minheap, vertex8, 1)
	FindKeyUpdateScore(minheap, vertex9, 10)
	FindKeyUpdateScore(minheap, vertex1, 20)
	FindKeyUpdateScore(minheap, vertex11, 0)
	FindKeyUpdateScore(minheap, vertex4, 25)
	FindKeyUpdateScore(minheap, vertex2, 16)
	FindKeyUpdateScore(minheap, vertex10, 12)
	FindKeyUpdateScore(minheap, vertex3, 2)
	FindKeyUpdateScore(minheap, vertex6, 21)

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
