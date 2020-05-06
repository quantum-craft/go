package graph

import (
	"testing"
)

func TestKosaraju(t *testing.T) {
	vertices, edges := ConstructGraphDirected("../data/KosarajuTestSmall.txt")

	Kosaraju(vertices, edges)

	if vertices[0].leader != 6 {
		t.Error("Kosaraju has error")
	}

	if vertices[3].leader != 6 {
		t.Error("Kosaraju has error")
	}

	if vertices[6].leader != 6 {
		t.Error("Kosaraju has error")
	}

	if vertices[1].leader != 7 {
		t.Error("Kosaraju has error")
	}

	if vertices[4].leader != 7 {
		t.Error("Kosaraju has error")
	}

	if vertices[7].leader != 7 {
		t.Error("Kosaraju has error")
	}

	if vertices[2].leader != 8 {
		t.Error("Kosaraju has error")
	}

	if vertices[5].leader != 8 {
		t.Error("Kosaraju has error")
	}

	if vertices[8].leader != 8 {
		t.Error("Kosaraju has error")
	}
}

func TestDFSDirected(t *testing.T) {
	vertices, edges := ConstructGraphDirected("../data/DFSDirectedTest1.txt")

	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraphDirected has error, not all vertices unexplored")
	}
	DFSDirected(vertices, edges, 0, nil)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSDirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraphDirected("../data/DFSDirectedTest1.txt")
	n := len(vertices)
	DFSDirected(vertices, edges, 0, &n)

	if CheckAllExplored(vertices) != true {
		t.Error("DFSDirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraphDirected("../data/DFSDirectedTest1.txt")
	TopologicalSort(vertices, edges)
	if CheckAllExplored(vertices) != true {
		t.Error("TopologicalSort has error, not all vertices explored")
	}

	if vertices[0].topologicalOrder != 1 {
		t.Error("TopologicalSort has error, v0's order should be 1")
	}

	if vertices[1].topologicalOrder != 3 {
		t.Error("TopologicalSort has error, v1's order should be 3")
	}

	if vertices[2].topologicalOrder != 2 {
		t.Error("TopologicalSort has error, v2's order should be 2")
	}

	if vertices[3].topologicalOrder != 4 {
		t.Error("TopologicalSort has error, v3's order should be 4")
	}

	vertices, edges = ConstructGraphDirected("../data/DFSDirectedTest1.txt")
	n = len(vertices)
	DFSDirected(vertices, edges, 2, &n)
	DFSDirected(vertices, edges, 0, &n)

	if CheckAllExplored(vertices) != true {
		t.Error("DFSDirected has error, not all vertices explored")
	}

	if vertices[0].topologicalOrder != 1 {
		t.Error("TopologicalSort has error, v0's order should be 1")
	}

	if vertices[1].topologicalOrder != 2 {
		t.Error("TopologicalSort has error, v1's order should be 2")
	}

	if vertices[2].topologicalOrder != 3 {
		t.Error("TopologicalSort has error, v2's order should be 3")
	}

	if vertices[3].topologicalOrder != 4 {
		t.Error("TopologicalSort has error, v3's order should be 4")
	}
}

func TestDFSUndirected(t *testing.T) {
	vertices, edges := ConstructGraph("../data/BFSTestSmall.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 0)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 3)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/BFSTestSmall.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 5)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 0)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 99)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 199)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 20)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 50)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}

	vertices, edges = ConstructGraph("../data/kargerMinCut.txt")
	if CheckAllUnexplored(vertices) != true {
		t.Error("ConstructGraph has error, not all vertices unexplored")
	}
	DFSUndirected(vertices, edges, 150)
	if CheckAllExplored(vertices) != true {
		t.Error("DFSUndirected has error, not all vertices explored")
	}
}
