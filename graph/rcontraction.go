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

type Vertex struct {
	idx   int
	edges []*Edge
}

type Edge struct {
	contracted bool
	head       *Vertex
	tail       *Vertex
}

// package-wise var
var r = rand.New(rand.NewSource(time.Now().Unix()))

func RContraction(vertices []Vertex, edges []Edge) (mincut int) {
	// new seed in every run
	r = rand.New(rand.NewSource(time.Now().Unix()))

	n, _ := len(vertices), len(edges)

	// ee182: optimize codes
	for i := n; i > 2; i-- {
		remainingEdges := getRemainingEdges(edges)

		// ee182
		fmt.Println(len(remainingEdges))

		contract := remainingEdges[r.Intn(len(remainingEdges))]

		contract.head.edges = append(contract.head.edges, contract.tail.edges...)

		// remove tail vertex
		for j := 0; j < len(contract.head.edges); j++ {
			if contract.head.edges[j].head == contract.tail {
				contract.head.edges[j].head = contract.head
			}

			if contract.head.edges[j].tail == contract.tail {
				contract.head.edges[j].tail = contract.head
			}
		}

		// remove contracted edge and self-loop
		for j := 0; j < len(contract.head.edges); j++ {
			if contract.head.edges[j].head == contract.head.edges[j].tail {
				// TODO: remove contracted edge from array
				contract.head.edges[j].contracted = true
				contract.head.edges[j] = nil
			}
		}

		newHeadEdges := make([]*Edge, 0)
		for j := 0; j < len(contract.head.edges); j++ {
			if contract.head.edges[j] != nil {
				newHeadEdges = append(newHeadEdges, contract.head.edges[j])
			}
		}
		contract.head.edges = newHeadEdges
	}

	return mincut
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

func ConstructGraph(filePath string) (vertices []Vertex, edges []Edge) {
	lineCnt := CountFileLines(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	vertices = make([]Vertex, lineCnt)
	edges = make([]Edge, 0)

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
				edges = append(edges, MakeEdge(vertices, v0-1, v-1))
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

func MakeVertex(idx int) Vertex {
	return Vertex{
		idx:   idx,
		edges: make([]*Edge, 0),
	}
}

func MakeEdge(vertices []Vertex, headIdx int, tailIdx int) Edge {
	return Edge{
		contracted: false,
		head:       &vertices[headIdx],
		tail:       &vertices[tailIdx],
	}
}
