package graph

import (
	"testing"
)

func TestBFSSmall(t *testing.T) {
	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	BFS(vertices, edges)

	if CheckAllExplored(vertices) != true {
		t.Error("BFS has error, not all vertices explored")
	}
}

func TestBFS(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	BFS(vertices, edges)

	if CheckAllExplored(vertices) != true {
		t.Error("BFS has error, not all vertices explored")
	}
}

func TestBFSInputSmall(t *testing.T) {
	edgeCnt := CountEdges("../data/BFSTestSmall.txt")
	if edgeCnt != 8 {
		t.Error("Edge count is wrong for file: ../data/BFSTestSmall.txt")
	}

	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")

	cnt := 0
	for i := 0; i < len(vertices); i++ {
		if vertices[i].edges != nil {
			cnt++
		}
	}

	if cnt != 6 {
		t.Error("Vertices count is wrong for file: ../data/BFSTestSmall.txt")
	}

	// fmt.Println(len(edges))

	if len(edges) != 8 {
		t.Error("ConstructGraph has wrong edge count for file ../data/BFSTestSmall.txt")
	}
}
