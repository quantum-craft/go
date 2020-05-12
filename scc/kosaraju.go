package scc

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx   int
	Edges []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	Head *Vertex
}

// Kosaraju is refactored from graph package
// It calculates the strongly connected components of the graph
func Kosaraju(file string) []int {
	verticesBackward, _ := ConstructGraphBackward(file)
	orders := KosarajuBackwardLoop(verticesBackward)

	verticesForward, _ := ConstructGraphForward(file)
	leaderCount := KosarajuForwardLoop(verticesForward, orders)

	return leaderCount
}

// KosarajuForwardLoop is the 2nd pass of Kosaraju
func KosarajuForwardLoop(vertices []Vertex, orders []int) []int {
	explored := make([]bool, len(vertices))
	leaderCount := make([]int, len(vertices))

	leader := -1

	for i := len(vertices) - 1; i >= 0; i-- {
		if explored[orders[i]] == false {
			leader = orders[i]
			KosarajuForward(vertices, orders[i], explored, leader, leaderCount, orders)
		}
	}

	return leaderCount
}

// KosarajuForward is the 2nd pass of Kosaraju, using DFS technique
func KosarajuForward(vertices []Vertex, startIdx int, explored []bool, leader int, leaderCount []int, orders []int) {
	v := &vertices[startIdx]
	explored[startIdx] = true
	leaderCount[leader]++

	for i, neighbors := 0, v.Edges; i < len(neighbors); i++ {
		if explored[neighbors[i].Head.idx] == false {
			KosarajuForward(vertices, neighbors[i].Head.idx, explored, leader, leaderCount, orders)
		}
	}
}

// KosarajuBackwardLoop is the 1st pass of Kosaraju
func KosarajuBackwardLoop(vertices []Vertex) []int {
	explored := make([]bool, len(vertices))
	orders := make([]int, len(vertices))

	finishingTime := 0

	for i := len(vertices) - 1; i >= 0; i-- {
		if explored[i] == false {
			KosarajuBackward(vertices, i, explored, &finishingTime, orders)
		}
	}

	return orders
}

// KosarajuBackward is the 1st pass of Kosaraju, using DFS technique
func KosarajuBackward(vertices []Vertex, startIdx int, explored []bool, finishingTime *int, orders []int) {
	v := &vertices[startIdx]
	explored[startIdx] = true

	for i, neighbors := 0, v.Edges; i < len(neighbors); i++ {
		if explored[neighbors[i].Head.idx] == false {
			KosarajuBackward(vertices, neighbors[i].Head.idx, explored, finishingTime, orders)
		}
	}

	*finishingTime++
	orders[*finishingTime-1] = v.idx
}

// DFSLoop runs thorugh entire graph
func DFSLoop(vertices []Vertex, edges []Edge, explored []bool) {
	for i := len(vertices) - 1; i >= 0; i-- {
		if explored[i] == false {
			DFS(vertices, edges, explored, i)
		}
	}
}

// DFS is a test for correctness of ConstructGraphForward/Backward
func DFS(vertices []Vertex, edges []Edge, explored []bool, startIdx int) {
	explored[startIdx] = true

	for i, neighbors := 0, vertices[startIdx].Edges; i < len(neighbors); i++ {
		if explored[neighbors[i].Head.idx] == false {
			DFS(vertices, edges, explored, neighbors[i].Head.idx)
		}
	}
}

// ConstructGraphForward will construct the forward adjacency list for 2nd pass of Kosaraju
func ConstructGraphForward(filePath string) ([]Vertex, []Edge) {
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
		head, _ := strconv.Atoi(fields[1])

		tail, head = tail-1, head-1

		if vertices[tail].Edges == nil {
			vertices[tail].Edges = make([]*Edge, 0)
			vertices[tail].idx = tail
		}

		if vertices[head].Edges == nil {
			vertices[head].Edges = make([]*Edge, 0)
			vertices[head].idx = head
		}

		currentEdges = edges[:len(currentEdges)+1]
		currentEdges[len(currentEdges)-1] = Edge{
			Head: &vertices[head],
		}

		vertices[tail].Edges = append(vertices[tail].Edges, &currentEdges[len(currentEdges)-1])
	}

	return vertices, edges
}

// ConstructGraphBackward will construct the backward adjacency list for 1st pass of Kosaraju
func ConstructGraphBackward(filePath string) ([]Vertex, []Edge) {
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

		// reverse the order
		tail, _ := strconv.Atoi(fields[1])
		head, _ := strconv.Atoi(fields[0])

		tail, head = tail-1, head-1

		if vertices[tail].Edges == nil {
			vertices[tail].Edges = make([]*Edge, 0)
			vertices[tail].idx = tail
		}

		if vertices[head].Edges == nil {
			vertices[head].Edges = make([]*Edge, 0)
			vertices[head].idx = head
		}

		currentEdges = edges[:len(currentEdges)+1]
		currentEdges[len(currentEdges)-1] = Edge{
			Head: &vertices[head],
		}

		vertices[tail].Edges = append(vertices[tail].Edges, &currentEdges[len(currentEdges)-1])
	}

	return vertices, edges
}

// VertexAndEdgeCountFromFile counts vertex counts of a directed edge file
func VertexAndEdgeCountFromFile(filePath string) (int, int) {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line, lineCnt := "", 0

	for scanner.Scan() {
		line = scanner.Text()
		lineCnt++
	}

	fields := strings.Fields(line)

	if len(fields) > 0 {
		i, _ := strconv.Atoi(fields[0])
		return i, lineCnt
	}

	return -1, lineCnt
}
