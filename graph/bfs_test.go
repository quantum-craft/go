package graph

import (
	"testing"
)

func TestBFSSmall(t *testing.T) {
	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	BFS(vertices, edges, 0)

	if CheckAllExplored(vertices) != true {
		t.Error("BFS has error, not all vertices explored")
	}
}

func TestBFS(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	BFS(vertices, edges, 0)

	if CheckAllExplored(vertices) != true {
		t.Error("BFS has error, not all vertices explored")
	}
}

func TestBFSShortestPathSmall(t *testing.T) {
	// vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")

	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")
	dist05 := BFSShortestPath(vertices, edges, 0, 5)
	if dist05 != 3 {
		t.Error("BFSShortestPath has error, dist from v0 to v5 should be 3")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist00 := BFSShortestPath(vertices, edges, 0, 0)
	if dist00 != 0 {
		t.Error("BFSShortestPath has error, dist from v0 to v0 should be 0")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist33 := BFSShortestPath(vertices, edges, 3, 3)
	if dist33 != 0 {
		t.Error("BFSShortestPath has error, dist from v3 to v3 should be 0")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist44 := BFSShortestPath(vertices, edges, 4, 4)
	if dist44 != 0 {
		t.Error("BFSShortestPath has error, dist from v4 to v4 should be 0")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist02 := BFSShortestPath(vertices, edges, 0, 2)
	if dist02 != 1 {
		t.Error("BFSShortestPath has error, dist from v0 to v2 should be 1")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist04 := BFSShortestPath(vertices, edges, 0, 4)
	if dist04 != 2 {
		t.Error("BFSShortestPath has error, dist from v0 to v4 should be 2")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist31 := BFSShortestPath(vertices, edges, 3, 1)
	if dist31 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v1 should be 1")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist32 := BFSShortestPath(vertices, edges, 3, 2)
	if dist32 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v2 should be 1")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist34 := BFSShortestPath(vertices, edges, 3, 4)
	if dist34 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v4 should be 1")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist35 := BFSShortestPath(vertices, edges, 3, 5)
	if dist35 != 1 {
		t.Error("BFSShortestPath has error, dist from v3 to v5 should be 1")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	dist30 := BFSShortestPath(vertices, edges, 3, 0)
	if dist30 != 2 {
		t.Error("BFSShortestPath has error, dist from v3 to v0 should be 2")
	}
}

func TestBFSConnectivity(t *testing.T) {
	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	clusterCnt := BFSConnectivity(vertices, edges)

	if CheckAllExplored(vertices) != true {
		t.Error("BFSConnectivity has error, not all vertices explored")
	}

	if clusterCnt != 1 {
		t.Error("BFSConnectivity has error, clusterCnt should be 1")
	}

	vertices, edges = ConstructGraph("../data/BFSConnectivityTest1.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	clusterCnt = BFSConnectivity(vertices, edges)

	if CheckAllExplored(vertices) != true {
		t.Error("BFSConnectivity has error, not all vertices explored")
	}

	if clusterCnt != 3 {
		t.Error("BFSConnectivity has error, clusterCnt should be 3")
	}

	edgeCnt := CountEdges("../data/BFSConnectivityTest2.txt")
	if edgeCnt != 23 {
		t.Error("Edge count is wrong for file: ../data/BFSConnectivityTest2.txt")
	}

	vertices, edges = ConstructGraph("../data/BFSConnectivityTest2.txt")

	cnt := 0
	for i := 0; i < len(vertices); i++ {
		if vertices[i].edges != nil {
			cnt++
		}
	}

	if cnt != 21 {
		t.Error("Vertices count is wrong for file: ../data/BFSConnectivityTest2.txt")
	}

	// fmt.Println(len(edges))

	if len(edges) != 23 {
		t.Error("ConstructGraph has wrong edge count for file ../data/BFSConnectivityTest2.txt")
	}

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}

	clusterCnt = BFSConnectivity(vertices, edges)

	if CheckAllExplored(vertices) != true {
		t.Error("BFSConnectivity has error, not all vertices explored")
	}

	if clusterCnt != 6 {
		t.Error("BFSConnectivity has error, clusterCnt should be 6")
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
