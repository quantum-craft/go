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
	edges []*Edge
}

type Edge struct {
	idx  int
	head *Vertex
	tail *Vertex
}

// package-wise var
var r = rand.New(rand.NewSource(time.Now().Unix()))

func RContraction(vertices []*Vertex, edges []*Edge) (mincut int) {
	// new seed in every run
	r = rand.New(rand.NewSource(time.Now().Unix()))

	n, _ := len(vertices), len(edges)

	for i := n; i > 2; i-- {
		contract := edges[r.Intn(len(edges))]
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
				// remove contracted edge from array
				// DEBUG codes:
				fmt.Println(contract.head.edges[j].idx)
				fmt.Println(len(edges))

				edges[contract.head.edges[j].idx] = nil
				contract.head.edges[j] = nil
			}
		}

		j := 0
		for j < len(contract.head.edges) {
			if contract.head.edges[j] == nil {
				contract.head.edges = append(contract.head.edges[0:j], contract.head.edges[j+1:]...)
			} else {
				j++
			}
		}

		// shrink and rearrange indices of edges
		j = 0
		for j < len(edges) {
			if edges[j] == nil {
				edges = append(edges[0:j], edges[j+1:]...)
			} else {
				j++
			}
		}

		for j := 0; j < len(edges); j++ {
			edges[j].idx = j
		}
	}

	fmt.Println(len(edges))

	return mincut
}

func MakeVertex() *Vertex {
	return &Vertex{
		edges: make([]*Edge, 0),
	}
}

func MakeEdge(vertices []*Vertex, headIdx int, tailIdx int, idx int) *Edge {
	return &Edge{
		idx:  idx,
		head: vertices[headIdx],
		tail: vertices[tailIdx],
	}
}

func ConstructGraph(filePath string) (vertices []*Vertex, edges []*Edge) {
	lineCnt := CountFileLines(filePath)

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	vertices = make([]*Vertex, lineCnt)
	edges = make([]*Edge, 0)

	for err == nil {
		adjacencyList := strings.Fields(line)

		v0, _ := strconv.Atoi(adjacencyList[0])

		if vertices[v0-1] == nil {
			vertices[v0-1] = MakeVertex()
		}

		for i := 1; i < len(adjacencyList); i++ {
			v, _ := strconv.Atoi(adjacencyList[i])

			if vertices[v-1] == nil {
				vertices[v-1] = MakeVertex()
			}

			edge := MakeEdge(vertices, v0-1, v-1, len(edges))
			edges = append(edges, edge)

			vertices[v0-1].edges = append(vertices[v0-1].edges, edge)
			vertices[v-1].edges = append(vertices[v-1].edges, edge)
		}

		line, err = rd.ReadString('\n')
	}

	return vertices, edges
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

func idxIncreasing(xs []*Edge) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1].idx-xs[i].idx != 1 {
			return false
		}
	}

	return true
}
