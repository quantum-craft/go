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
				contract.head.edges[j] = nil
			}
		}

		// shrink edges
		// remove tail vertex from array
		// remove contracted edge from array
	}

	return mincut
}

func MakeVertex() *Vertex {
	return &Vertex{
		edges: make([]*Edge, 0),
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

			edge := &Edge{
				head: vertices[v0-1],
				tail: vertices[v-1],
			}
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
