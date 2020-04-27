package graph

// DFS is the template of depth-first-search
func DFS(vertices []Vertex, edges []Edge, startIdx int) {
	v := &vertices[startIdx]
	v.explored = true

	for i, neighbors := 0, v.edges; i < len(neighbors); i++ {
		if neighbors[i].head == v {
			if neighbors[i].tail.explored == false {
				DFS(vertices, edges, neighbors[i].tail.idx)
			}
		} else {
			if neighbors[i].head.explored == false {
				DFS(vertices, edges, neighbors[i].head.idx)
			}
		}
	}
}
