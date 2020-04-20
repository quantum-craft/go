package graph

import (
	"fmt"
	"math"
	"testing"
)

func TestRContractionInput(t *testing.T) {
	edgeCnt := CountEdges("../data/kargerMinCut.txt")
	if edgeCnt != 2517 {
		t.Error("Edge count is wrong for file: ../data/kargerMinCut.txt")
	}

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
	n := len(vertices)
	lgn := int(math.Log2(float64(n)))

	minOfMincut := 100000 // larger than 2517 :P
	for i := 0; i < n*n*lgn; i++ {
		vertices, edges = ConstructGraph("../data/kargerMinCut.txt")

		mincut := RContraction(vertices, edges)

		if mincut < minOfMincut {
			minOfMincut = mincut
		}

		fmt.Println(mincut)
	}

	fmt.Print("min-cut is")
	fmt.Println(minOfMincut)
}
