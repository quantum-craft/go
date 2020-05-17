package dijkstra

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
	Head   *Vertex
	Weight int
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
				vertices[tail].Edges = make([]*Edge, 0)
				vertices[tail].idx = tail
			}

			if vertices[head].Edges == nil {
				vertices[head].Edges = make([]*Edge, 0)
				vertices[head].idx = head
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
