package dijkstra

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

// Node is the node unit in min-heap for Dijkstra algorithm
type Node struct {
	Key *Vertex
}

// MinHeap keeps minimum Score on the top and also keeps the heap property
type MinHeap struct {
	lastEmpty *int
	heap      *[]Node
}

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

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

	for n := ExtractMin(minheap); n.Key != nil; n = ExtractMin(minheap) {
		n.Key.Explored = true
		for i, edges := 0, n.Key.Edges; i < len(n.Key.Edges); i++ {
			if edges[i].Head.Explored == false {
				// check whether Head exists in heap
				if !FindKeyUpdateScore(minheap, edges[i].Head, n.Key.Score+edges[i].Weight) {
					edges[i].Head.Score = n.Key.Score + edges[i].Weight
					Insert(minheap, Node{Key: edges[i].Head})
				}
			}
		}
	}
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
				vertices[tail].heapIdx = -1 // not in heap yet
				vertices[tail].Score = maxInt
				vertices[tail].Explored = false
				vertices[tail].Edges = make([]*Edge, 0)
			}

			if vertices[head].Edges == nil {
				vertices[head].idx = head
				vertices[head].heapIdx = -1 // not in heap yet
				vertices[head].Score = maxInt
				vertices[head].Explored = false
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

// Insert will insert the node at the bottom and re-heapify
func Insert(minheap MinHeap, n Node) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if *lastEmpty == len(*heap) {
		*heap = append(*heap, n)
		*lastEmpty++
	} else {
		(*heap)[*lastEmpty] = n
		*lastEmpty++
	}

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)

	bubbleUp(*lastEmpty, heap)
}

// pos is one based index
func bubbleUp(pos int, heap *[]Node) {
	for p := pos; p > 1 && (*heap)[p/2-1].Key.Score >= (*heap)[p-1].Key.Score; p = p / 2 {
		swapNode(heap, p-1, p/2-1)
	}
}

// pos is one based index
func bubbleDown(pos int, minheap MinHeap) {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	p := pos
	for {
		if p-1 >= *lastEmpty || p*2-1 >= *lastEmpty {
			return
		}

		here := downHere(p, minheap)

		if here == p {
			return
		}

		swapNode(heap, p-1, here-1)

		p = here
	}
}

func downHere(p int, minheap MinHeap) int {
	lastEmpty, heap := minheap.lastEmpty, minheap.heap

	if p*2 >= *lastEmpty {
		return findMinScorePos2(heap, p, p*2)
	}

	return findMinScorePos3(heap, p, p*2, p*2+1)
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and re-heapify
func ExtractMin(minheap MinHeap) Node {
	var lastEmpty *int = minheap.lastEmpty

	if *lastEmpty == 0 {
		return Node{Key: nil}
	}

	heap := minheap.heap
	ret := (*heap)[0]

	*lastEmpty--
	swapNode(heap, 0, *lastEmpty)

	bubbleDown(1, minheap)

	return ret
}

// FindKeyUpdateScore will find the vertex in the heap and update score if it is smaller
func FindKeyUpdateScore(minheap MinHeap, key *Vertex, score int) bool {
	var heap *[]Node = minheap.heap

	if key.heapIdx >= 0 && (*heap)[key.heapIdx].Key.idx == key.idx {
		if score < (*heap)[key.heapIdx].Key.Score {
			(*heap)[key.heapIdx].Key.Score = score

			// new score is smaller, re-heapify
			bubbleUp(key.heapIdx+1, heap)
		}

		return true
	}

	return false
}

// pos is one based
func findMinScorePos2(heap *[]Node, pos1 int, pos2 int) int {
	if (*heap)[pos1-1].Key.Score < (*heap)[pos2-1].Key.Score {
		return pos1
	}

	return pos2
}

// pos is one based
func findMinScorePos3(heap *[]Node, pos1 int, pos2 int, pos3 int) int {
	return findMinScorePos2(heap, findMinScorePos2(heap, pos1, pos2), pos3)
}

func swapNode(heap *[]Node, this int, that int) {
	(*heap)[this], (*heap)[that] = (*heap)[that], (*heap)[this]
	(*heap)[this].Key.heapIdx = this
	(*heap)[that].Key.heapIdx = that
}
