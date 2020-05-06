package graph

// DFSUndirected is the template of depth-first-search for undirected graph
func DFSUndirected(vertices []Vertex, edges []Edge, startIdx int) {
	v := &vertices[startIdx]
	v.explored = true

	for i, neighbors := 0, v.edges; i < len(neighbors); i++ {
		if neighbors[i].tail.explored == false {
			DFSUndirected(vertices, edges, neighbors[i].tail.idx)
		}
		if neighbors[i].head.explored == false {
			DFSUndirected(vertices, edges, neighbors[i].head.idx)
		}
	}
}

// DFSDirected is the template of depth-first-search for directed graph
// DFSDirected is also used as a sub-routine of TopologicalSort
func DFSDirected(vertices []Vertex, edges []Edge, startIdx int, currentLabel *int) {
	v := &vertices[startIdx]
	v.explored = true

	for i, neighbors := 0, v.edges; i < len(neighbors); i++ {
		if neighbors[i].head.explored == false {
			DFSDirected(vertices, edges, neighbors[i].head.idx, currentLabel)
		}
	}

	if currentLabel != nil {
		v.topologicalOrder = *currentLabel
		*currentLabel--
	}
}

// TopologicalSort uses a DFS sub-routine to solve topological sorting problem
func TopologicalSort(vertices []Vertex, edges []Edge) {
	MarkAllUnexplored(vertices)

	currentLabel := len(vertices)
	for i := 0; i < len(vertices); i++ {
		if vertices[i].explored == false {
			DFSDirected(vertices, edges, vertices[i].idx, &currentLabel)
		}
	}
}
