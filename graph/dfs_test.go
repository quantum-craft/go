package graph

import "testing"

func TestDFS(t *testing.T) {
	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 0)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 3)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 5)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 0)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 99)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 199)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 20)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 50)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFS(vertices, edges, 150)
	if CheckAllExplored(vertices) != true {
		t.Error("DFS has error, not all vertices explored")
	}
}
