package graph

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx        int
	contracted bool
	edges      []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	contracted bool
	head       *Vertex
	tail       *Vertex
}

var r = rand.New(rand.NewSource(time.Now().Unix())) // package-wise var

// RContraction use Karger's randomized contraction algorithm
// to find the min-cut of an undirected graph
func RContraction(vertices []Vertex, edges []Edge) int {
	r = rand.New(rand.NewSource(time.Now().Unix())) // new seed in every run

	for i := 0; i < len(vertices)-2; i++ {
		remainingEdges := getRemainingEdges(edges)
		contracted := remainingEdges[r.Intn(len(remainingEdges))]

		contracted.head.edges = append(contracted.head.edges, contracted.tail.edges...)

		// remove tail vertex
		contracted.tail.contracted = true
		for j := 0; j < len(contracted.head.edges); j++ {
			if contracted.head.edges[j].head == contracted.tail {
				contracted.head.edges[j].head = contracted.head
			}

			if contracted.head.edges[j].tail == contracted.tail {
				contracted.head.edges[j].tail = contracted.head
			}
		}

		// remove contracted edge and self-loop
		for j := 0; j < len(contracted.head.edges); j++ {
			if contracted.head.edges[j].head == contracted.head.edges[j].tail {
				// TODO: remove contracted edge from array
				contracted.head.edges[j].contracted = true
				contracted.head.edges[j] = nil
			}
		}

		newHeadEdges := make([]*Edge, 0)
		for j := 0; j < len(contracted.head.edges); j++ {
			if contracted.head.edges[j] != nil {
				newHeadEdges = append(newHeadEdges, contracted.head.edges[j])
			}
		}
		contracted.head.edges = newHeadEdges
	}

	return len(getRemainingEdges(edges))
}

func getRemainingEdges(edges []Edge) []Edge {
	ret := make([]Edge, 0)

	for i := 0; i < len(edges); i++ {
		if edges[i].contracted == false {
			ret = append(ret, edges[i])
		}
	}

	return ret
}

func getRemainingVertices(vertices []Vertex) []Vertex {
	ret := make([]Vertex, 0)

	for i := 0; i < len(vertices); i++ {
		if vertices[i].contracted == false {
			ret = append(ret, vertices[i])
		}
	}

	return ret
}

// ConstructGraph will construct the adjacency list for file designated by filePath
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

// MakeVertex gives you a new Vertex
func MakeVertex(idx int) Vertex {
	return Vertex{
		idx:        idx,
		contracted: false,
		edges:      make([]*Edge, 0),
	}
}

// MakeEdge gives you a new Edge
func MakeEdge(vertices []Vertex, headIdx int, tailIdx int) Edge {
	return Edge{
		contracted: false,
		head:       &vertices[headIdx],
		tail:       &vertices[tailIdx],
	}
}
