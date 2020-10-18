package dijkstra

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	heap "github.com/quantum-craft/go/minheap"
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

// GetCost implements the interface of min-heap
func (n Node) GetCost() int {
	return n.Key.Score
}

// SetCost implements the interface of min-heap
func (n Node) SetCost(newCost int) {
	n.Key.Score = newCost
}

// GetHeapIdx implements the interface of min-heap
func (n Node) GetHeapIdx() int {
	return n.Key.heapIdx
}

// SetHeapIdx implements the interface of min-heap
func (n Node) SetHeapIdx(idx int) {
	n.Key.heapIdx = idx
}

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

// Dijkstra computes shortest path from startIdx to all other reachable nodes
func Dijkstra(vertices []Vertex, edges []Edge, startIdx int) {
	minheap := heap.MakeMinHeap()

	vertices[startIdx].Score = 0 // source
	minheap.Insert(Node{Key: &vertices[startIdx]})

	n, ok := minheap.ExtractMin().(Node)
	for ok == true {
		n.Key.Explored = true
		for i, edges := 0, n.Key.Edges; i < len(n.Key.Edges); i++ {
			if edges[i].Head.Explored == false {

				// check whether Head exists in heap
				if edges[i].Head.heapIdx == -1 {
					edges[i].Head.Score = n.Key.Score + edges[i].Weight
					minheap.Insert(Node{Key: edges[i].Head})
				} else {
					minheap.Update(edges[i].Head.heapIdx, n.Key.Score+edges[i].Weight)
				}
			}
		}

		n, ok = minheap.ExtractMin().(Node)
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
