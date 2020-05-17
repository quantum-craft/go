package dijkstra

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx   int
	Score int // Dijkstra's greedy score
	Edges []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	Head   *Vertex
	Weight int
}

// Dijkstra computes shortest path from startIdx to all other reachable nodes
func Dijkstra(vertices []Vertex, edges []Edge, startIdx int) {
	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	vertices[startIdx].Score = 0 // source
	Insert(minheap, Node{Key: &vertices[startIdx]})

}

// ConstructGraph constructs the undirected, weighted graph for Dijkstra algorithm
func ConstructGraph(filePath string) ([]Vertex, []Edge) {
	vertexCnt, edgeCnt := VertexAndEdgeCountFromFile(filePath)

	vertices := make([]Vertex, vertexCnt)
	edges := make([]Edge, edgeCnt)
	currentEdges := edges[:0]

	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		tail, _ := strconv.Atoi(fields[0])
		tail = tail - 1
		for i := 1; i < len(fields); i++ {
			w := strings.Split(fields[i], ",")
			head, _ := strconv.Atoi(w[0])
			head = head - 1
			weight, _ := strconv.Atoi(w[1])

			if vertices[tail].Edges == nil {
				vertices[tail].idx = tail
				vertices[tail].Score = 1000000 // max weight for our application
				vertices[tail].Edges = make([]*Edge, 0)
			}

			if vertices[head].Edges == nil {
				vertices[head].idx = head
				vertices[head].Score = 1000000 // max weight for our application
				vertices[head].Edges = make([]*Edge, 0)
			}

			currentEdges = edges[:len(currentEdges)+1]
			currentEdges[len(currentEdges)-1] = Edge{
				Head:   &vertices[head],
				Weight: weight,
			}

			vertices[tail].Edges = append(vertices[tail].Edges, &currentEdges[len(currentEdges)-1])
		}
	}

	return vertices, edges
}

// VertexAndEdgeCountFromFile counts vertex and edge of a undirected, weighted graph file
func VertexAndEdgeCountFromFile(filePath string) (int, int) {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line, lineCnt, edgeCnt := "", 0, 0

	for scanner.Scan() {
		line = scanner.Text()
		lineCnt++

		fields := strings.Fields(line)

		edgeCnt += (len(fields) - 1)
	}

	return lineCnt, edgeCnt
}

// Node is node unit for Dijkstra algorithm
type Node struct {
	Key *Vertex
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
	for pos := *lastEmpty; pos > 1 && (*heap)[pos/2-1].Key.Score >= (*heap)[pos-1].Key.Score; pos = pos / 2 {
		swapNode(heap, pos-1, pos/2-1)
	}
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and bubble down it to the proper position
func ExtractMin(minheap MinHeap) Node {
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == 0 {
		return Node{Key: nil}
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

	if (*heap)[pos1-1].Key.Score < (*heap)[pos2-1].Key.Score {
		minPos = pos1
	} else {
		minPos = pos2
	}

	return minPos
}

// pos is one based
func findMinScorePos3(heap *[]Node, pos1 int, pos2 int, pos3 int) int {
	minPos := -1

	if (*heap)[pos1-1].Key.Score < (*heap)[pos2-1].Key.Score {
		minPos = pos1
	} else {
		minPos = pos2
	}

	if (*heap)[pos3-1].Key.Score < (*heap)[minPos-1].Key.Score {
		minPos = pos3
	}

	return minPos
}

func swapNode(heap *[]Node, this int, that int) {
	n := (*heap)[this]
	(*heap)[this] = (*heap)[that]
	(*heap)[that] = n
}
