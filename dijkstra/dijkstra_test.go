package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	file := "../data/dijkstraData.txt"

	ConstructGraph(file)
}
