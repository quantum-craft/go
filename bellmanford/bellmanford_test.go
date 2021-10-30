package bellmanford

import (
	"math"
	"testing"
)

func TestBellmanford(t *testing.T) {
	// Number of vertices
	n := 5

	edges := [][3]int{{0, 1, 2}, {0, 2, 4}, {1, 3, 2}, {1, 2, 1}, {2, 4, 4}, {3, 4, 2}}

	if Bellmanford(n, edges, 0, 4) != 6 {
		t.Error("TestBellmanford has error !")
	}
}

func TestBellmanfordNegative(t *testing.T) {
	// Number of vertices
	n := 6

	edges := [][3]int{{0, 1, 2}, {0, 2, 4}, {1, 3, 2}, {1, 2, 1}, {2, 4, 4}, {3, 4, 2}, {4, 5, 3}, {5, 2, -8}}

	if Bellmanford(n, edges, 0, 4) != math.MaxInt32 {
		t.Error("TestBellmanfordNegative has error !")
	}
}
