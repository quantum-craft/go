package dijkstra

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx   int
	Edges []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	Head *Vertex
}
