package graph

import (
	"fmt"
	"testing"
)

func TestRContractionInput(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")
	// ee182: fix
	if len(vertices) != 200 {
		t.Error("RContractionInput error !")
	}

	if len(edges) != 5034 {
		t.Error("RContractionInput error !")
	}
}

func TestRContraction(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	mincut := RContraction(vertices, edges)

	fmt.Println(mincut)
}
