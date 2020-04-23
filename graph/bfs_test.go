package graph

import (
	"testing"
)

func TestBFSSmall(t *testing.T) {
	// vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")
	vertices, _ := ConstructGraph("../data/BFSTestSmall.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	BFS(&vertices[0])

	if CheckAllExplored(vertices) != true {
		t.Error("BFS has error, not all vertices explored")
	}
}

func TestBFS(t *testing.T) {
	// vertices, edges := ConstructGraph("../data/kargerMinCut.txt")
	vertices, _ := ConstructGraph("../data/kargerMinCut.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	BFS(&vertices[0])

	if CheckAllExplored(vertices) != true {
		t.Error("BFS has error, not all vertices explored")
	}
}

func TestBFSShortestPathSmall(t *testing.T) {
	// vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")

	vertices, _ := ConstructGraph("../data/BFSTestSmall.txt")
	dist05 := BFSShortestPath(&vertices[0], &vertices[5])
	if dist05 != 3 {
		t.Error("BFSShortestPath has error, dist from v0 to v5 should be 3")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist00 := BFSShortestPath(&vertices[0], &vertices[0])
	if dist00 != 0 {
		t.Error("BFSShortestPath has error, dist from v0 to v0 should be 0")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist33 := BFSShortestPath(&vertices[3], &vertices[3])
	if dist33 != 0 {
		t.Error("BFSShortestPath has error, dist from v3 to v3 should be 0")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist44 := BFSShortestPath(&vertices[4], &vertices[4])
	if dist44 != 0 {
		t.Error("BFSShortestPath has error, dist from v4 to v4 should be 0")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist02 := BFSShortestPath(&vertices[0], &vertices[2])
	if dist02 != 1 {
		t.Error("BFSShortestPath has error, dist from v0 to v2 should be 1")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist04 := BFSShortestPath(&vertices[0], &vertices[4])
	if dist04 != 2 {
		t.Error("BFSShortestPath has error, dist from v0 to v4 should be 2")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist31 := BFSShortestPath(&vertices[3], &vertices[1])
	if dist31 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v1 should be 1")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist32 := BFSShortestPath(&vertices[3], &vertices[2])
	if dist32 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v2 should be 1")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist34 := BFSShortestPath(&vertices[3], &vertices[4])
	if dist34 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v4 should be 1")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist35 := BFSShortestPath(&vertices[3], &vertices[5])
	if dist35 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v5 should be 1")
	}

	vertices, _ = ConstructGraph("../data/BFSTestSmall.txt")
	dist30 := BFSShortestPath(&vertices[3], &vertices[0])
	if dist30 != 2 {
		t.Error("BFSShortestPath has error, dist from v3 to v0 should be 2")
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
