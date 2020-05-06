package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MakeVertex gives you a new Vertex
func MakeVertex(idx int) Vertex {
	return Vertex{
		idx:              idx,
		contracted:       false,
		explored:         false,
		dist:             ^uint64(0), // infinite
		topologicalOrder: -1,         // impossible order
		finishingTime:    -1,         // impossible order
		leader:           -1,         // impossible index
		edges:            make([]*Edge, 0),
	}
}

// MakeEdge gives you a new Edge
func MakeEdge(vertices []Vertex, tailIdx int, headIdx int) Edge {
	return Edge{
		contracted: false,
		head:       &vertices[headIdx],
		tail:       &vertices[tailIdx],
	}
}

// ConstructGraphDirected will construct the adjacency list for directed graph file designated by filePath
func ConstructGraphDirected(filePath string) ([]Vertex, []Edge) {
	lineCnt := CountFileLines(filePath)
	edgeCnt := CountEdgesDirected(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	vertices := make([]Vertex, lineCnt)
	edgesSpace := make([]Edge, edgeCnt)
	edges := edgesSpace[:0]

	for err == nil {
		adjacencyList := strings.Fields(line)

		v0, _ := strconv.Atoi(adjacencyList[0])

		if vertices[v0-1].edges == nil {
			vertices[v0-1] = MakeVertex(v0 - 1)
		}

		for i := 1; i < len(adjacencyList); i++ {
			v, _ := strconv.Atoi(adjacencyList[i])

			if vertices[v-1].edges == nil {
				vertices[v-1] = MakeVertex(v - 1)
			}

			if !repeatedEdges(edges, v0-1, v-1) {
				edges = edgesSpace[:len(edges)+1]
				edges[len(edges)-1] = MakeEdge(vertices, v0-1, v-1)

				vertices[v0-1].edges = append(vertices[v0-1].edges, &edges[len(edges)-1])
				vertices[v-1].edges = append(vertices[v-1].edges, &edges[len(edges)-1])
			}
		}

		line, err = rd.ReadString('\n')
	}

	return vertices, edges
}

// ConstructGraph will construct the adjacency list for undirected graph file designated by filePath
func ConstructGraph(filePath string) ([]Vertex, []Edge) {
	lineCnt := CountFileLines(filePath)
	edgeCnt := CountEdges(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	vertices := make([]Vertex, lineCnt)
	edgesSpace := make([]Edge, edgeCnt)
	edges := edgesSpace[:0]

	for err == nil {
		adjacencyList := strings.Fields(line)

		v0, _ := strconv.Atoi(adjacencyList[0])

		if vertices[v0-1].edges == nil {
			vertices[v0-1] = MakeVertex(v0 - 1)
		}

		for i := 1; i < len(adjacencyList); i++ {
			v, _ := strconv.Atoi(adjacencyList[i])

			if vertices[v-1].edges == nil {
				vertices[v-1] = MakeVertex(v - 1)
			}

			if !repeatedEdges(edges, v0-1, v-1) {
				edges = edgesSpace[:len(edges)+1]
				edges[len(edges)-1] = MakeEdge(vertices, v0-1, v-1)

				vertices[v0-1].edges = append(vertices[v0-1].edges, &edges[len(edges)-1])
				vertices[v-1].edges = append(vertices[v-1].edges, &edges[len(edges)-1])
			}
		}

		line, err = rd.ReadString('\n')
	}

	return vertices, edges
}

func repeatedEdges(edges []Edge, headIdx int, tailIdx int) bool {
	for j := 0; j < len(edges); j++ {
		if edges[j].head.idx == headIdx && edges[j].tail.idx == tailIdx {
			return true
		} else if edges[j].tail.idx == headIdx && edges[j].head.idx == tailIdx {
			return true
		}
	}

	return false
}

// CountFileLines returns the lines of input file,
// which is the number of vertices
func CountFileLines(filePath string) (cnt int) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	_, err = rd.ReadString('\n')

	for err == nil {
		cnt++
		_, err = rd.ReadString('\n')
	}

	return cnt
}

// CountEdges returns the number of edges of the input file
func CountEdges(filePath string) (cnt int) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	cnt = 0
	for err == nil {
		adjacencyList := strings.Fields(line)
		cnt += len(adjacencyList) - 1
		line, err = rd.ReadString('\n')
	}

	return cnt / 2
}

// CountEdgesDirected returns the number of edges of the directed graph input file
func CountEdgesDirected(filePath string) (cnt int) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	cnt = 0
	for err == nil {
		adjacencyList := strings.Fields(line)
		cnt += len(adjacencyList) - 1
		line, err = rd.ReadString('\n')
	}

	return cnt
}

// CheckAllUnexplored checks whether the vertices are all unexplored
func CheckAllUnexplored(vertices []Vertex) bool {
	for i := 0; i < len(vertices); i++ {
		if vertices[i].explored == true {
			return false
		}
	}

	return true
}

// MarkAllUnexplored marks all vertices unexplored
func MarkAllUnexplored(vertices []Vertex) {
	for i := 0; i < len(vertices); i++ {
		vertices[i].explored = false
	}
}

// CheckAllExplored checks whether the vertices are all explored
func CheckAllExplored(vertices []Vertex) bool {
	for i := 0; i < len(vertices); i++ {
		if vertices[i].explored == false {
			return false
		}
	}

	return true
}
