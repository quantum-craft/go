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

	fmt.Println(vertices[6].Score)   // 7
	fmt.Println(vertices[36].Score)  // 37
	fmt.Println(vertices[58].Score)  // 59
	fmt.Println(vertices[81].Score)  // 82
	fmt.Println(vertices[98].Score)  // 99
	fmt.Println(vertices[114].Score) // 115
	fmt.Println(vertices[132].Score) // 133
	fmt.Println(vertices[164].Score) // 165
	fmt.Println(vertices[187].Score) // 188
	fmt.Println(vertices[196].Score) // 197

	// for i := 0; i < len(vertices); i++ {
	// 	fmt.Println(vertices[i].Score)
	// }
}

func TestMinHeap(t *testing.T) {
	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	vertex1 := &Vertex{idx: 0, Score: 88}
	vertex2 := &Vertex{idx: 1, Score: 32}
	vertex3 := &Vertex{idx: 2, Score: 50}
	vertex4 := &Vertex{idx: 3, Score: 90}
	vertex5 := &Vertex{idx: 4, Score: 23}
	vertex6 := &Vertex{idx: 5, Score: 74}
	vertex7 := &Vertex{idx: 6, Score: 2}
	vertex8 := &Vertex{idx: 7, Score: 5}
	vertex9 := &Vertex{idx: 8, Score: 91}
	vertex10 := &Vertex{idx: 9, Score: 13}
	vertex11 := &Vertex{idx: 10, Score: 1}

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
	FindKeyUpdateScore(minheap, vertex1, 50)
	FindKeyUpdateScore(minheap, vertex11, 0)
	FindKeyUpdateScore(minheap, vertex4, 32)
	FindKeyUpdateScore(minheap, vertex2, 16)

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
