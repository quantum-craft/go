package dijkstra

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	file := "../data/dijkstraData.txt"

	vertices, edges := ConstructGraph(file)
	Dijkstra(vertices, edges, 0)

	if vertices[6].Score != 2599 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[36].Score != 2610 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[58].Score != 2947 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[81].Score != 2052 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[98].Score != 2367 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[114].Score != 2399 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[132].Score != 2029 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[164].Score != 2442 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[187].Score != 2505 {
		t.Error("Dijkstra algorithm has error !")
	}

	if vertices[196].Score != 3068 {
		t.Error("Dijkstra algorithm has error !")
	}
}
