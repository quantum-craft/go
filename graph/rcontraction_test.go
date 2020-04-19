package graph

import (
	"fmt"
	"testing"
)

func TestRContractionInput(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	cnt := 0
	for i := 0; i < len(vertices); i++ {
		if vertices[i].edges != nil {
			cnt++
		}
	}

	if cnt != 200 {
		t.Error("RContractionInput error !")
	}

	fmt.Println(len(edges))

	if len(edges) != 2517 {
		t.Error("RContractionInput error !")
	}
}

func TestRContraction(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	mincut := RContraction(vertices, edges)

	fmt.Println(mincut)
}
