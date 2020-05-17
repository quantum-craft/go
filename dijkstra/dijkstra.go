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
	lastEmpty *int // zero based
	heap      *[]Node
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
				vertices[tail].heapIdx = -1    // not in heap yet
				vertices[tail].Score = 1000000 // max weight for our application
				vertices[tail].Explored = false
				vertices[tail].Edges = make([]*Edge, 0)
			}

			if vertices[head].Edges == nil {
				vertices[head].idx = head
				vertices[head].heapIdx = -1    // not in heap yet
				vertices[head].Score = 1000000 // max weight for our application
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

// Insert will insert the node at the bottom and BUBBLE UP to the proper position
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

	swapNode(heap, *lastEmpty-1, *lastEmpty-1)
	// pos is one based
	for pos := *lastEmpty; pos > 1 && (*heap)[pos/2-1].Key.Score >= (*heap)[pos-1].Key.Score; pos = pos / 2 {
		swapNode(heap, pos-1, pos/2-1)
	}
}

// ExtractMin will extract the minimum member, replace the minimum pos with the last member,
// and BUBBLE DOWN it to the proper position
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

// FindKeyUpdateScore will find the vertex in the heap and update score if it is smaller
func FindKeyUpdateScore(minheap MinHeap, key *Vertex, score int) bool {
	var heap *[]Node = minheap.heap

	if key.heapIdx >= 0 && (*heap)[key.heapIdx].Key.idx == key.idx {
		if score < (*heap)[key.heapIdx].Key.Score {
			(*heap)[key.heapIdx].Key.Score = score
			// pos is one based
			// score is smaller => BUBBLE UP
			for pos := key.heapIdx + 1; pos > 1 && (*heap)[pos/2-1].Key.Score >= (*heap)[pos-1].Key.Score; pos = pos / 2 {
				swapNode(heap, pos-1, pos/2-1)
			}
		}
		return true
	}

	return false
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
	(*heap)[this], (*heap)[that] = (*heap)[that], (*heap)[this]
	(*heap)[this].Key.heapIdx = this
	(*heap)[that].Key.heapIdx = that
}
