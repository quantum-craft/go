package kruskalmst

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMaxSpacingClustering(t *testing.T) {
	f, _ := os.Open("../data/four_clustering1.txt")
	defer f.Close()

	var numVertices, numEdges int
	var vertices []Vertex
	var edges []Edge

	k := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 1 {
			numVertices, _ = strconv.Atoi(fields[0])
			numEdges = numVertices * (numVertices - 1) / 2 // n choose 2

			if numVertices != 500 {
				t.Error("TestMaxSpacingClustering error! Vertex count is wrong.")
			}

			if numEdges != 124750 {
				t.Error("TestMaxSpacingClustering error! Edge count is wrong.")
			}

			vertices = make([]Vertex, numVertices)
			edges = make([]Edge, numEdges)
		} else {
			vidx1, _ := strconv.Atoi(fields[0])
			vidx2, _ := strconv.Atoi(fields[1])
			cost, _ := strconv.Atoi(fields[2])

			vidx1, vidx2 = vidx1-1, vidx2-1 // convert to zero based

			edges[k].VertIdx[0] = vidx1
			edges[k].VertIdx[1] = vidx2
			edges[k].Cost = cost

			vertices[vidx1].VIdx = vidx1
			vertices[vidx2].VIdx = vidx2
			vertices[vidx1].GroupLeader = &vertices[vidx1] // assign self as leader
			vertices[vidx2].GroupLeader = &vertices[vidx2] // assign self as leader
			vertices[vidx1].GroupSize = 1                  // only self
			vertices[vidx2].GroupSize = 1                  // only self
			vertices[vidx1].Added = false
			vertices[vidx2].Added = false

			// if vertices[vidx1].Edges == nil {
			// 	vertices[vidx1].Edges = make([]*Edge, 0)
			// }

			// if vertices[vidx2].Edges == nil {
			// 	vertices[vidx2].Edges = make([]*Edge, 0)
			// }

			// vertices[vidx1].Edges = append(vertices[vidx1].Edges, &edges[k])
			// vertices[vidx2].Edges = append(vertices[vidx2].Edges, &edges[k])

			k++
		}
	}

	maxSpacing := MaxSpacingClustering(vertices, edges, 4)

	if maxSpacing != 106 {
		t.Error("TestMaxSpacingClustering error!")
	}

	if maxDepth > 5 {
		t.Error("TestMaxSpacingClustering error, leader chain is too long!")
	}
}

func TestKruskalMST(t *testing.T) {
	f, _ := os.Open("../data/edges.txt")
	defer f.Close()

	var numVertices, numEdges int
	var vertices []Vertex
	var edges []Edge

	k := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 2 {
			numVertices, _ = strconv.Atoi(fields[0])
			numEdges, _ = strconv.Atoi(fields[1])

			vertices = make([]Vertex, numVertices)
			edges = make([]Edge, numEdges)
		} else {
			vidx1, _ := strconv.Atoi(fields[0])
			vidx2, _ := strconv.Atoi(fields[1])
			cost, _ := strconv.Atoi(fields[2])

			vidx1, vidx2 = vidx1-1, vidx2-1 // convert to zero based

			edges[k].VertIdx[0] = vidx1
			edges[k].VertIdx[1] = vidx2
			edges[k].Cost = cost

			vertices[vidx1].VIdx = vidx1
			vertices[vidx2].VIdx = vidx2
			vertices[vidx1].GroupLeader = &vertices[vidx1] // assign self as leader
			vertices[vidx2].GroupLeader = &vertices[vidx2] // assign self as leader
			vertices[vidx1].GroupSize = 1                  // only self
			vertices[vidx2].GroupSize = 1                  // only self
			vertices[vidx1].Added = false
			vertices[vidx2].Added = false

			// if vertices[vidx1].Edges == nil {
			// 	vertices[vidx1].Edges = make([]*Edge, 0)
			// }

			// if vertices[vidx2].Edges == nil {
			// 	vertices[vidx2].Edges = make([]*Edge, 0)
			// }

			// vertices[vidx1].Edges = append(vertices[vidx1].Edges, &edges[k])
			// vertices[vidx2].Edges = append(vertices[vidx2].Edges, &edges[k])

			k++
		}
	}

	for i := 0; i < len(vertices); i++ {
		if findLeader(&vertices[i]) != &vertices[i] {
			t.Error("TestKruskalMST error! Initially GroupLeader of each vertex should be itself.")
		}

		if findGroupSize(&vertices[i]) != 1 {
			t.Error("TestKruskalMST error! Initially GroupSize of each group should be 1.")
		}
	}

	minCost := KruskalMST(vertices, edges)

	if minCost != -3612829 {
		t.Error("TestKruskalMST error !")
	}

	if maxDepth > 5 {
		t.Error("TestKruskalMST error, leader chain is too long!")
	}
}
