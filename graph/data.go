package graph

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx        int
	contracted bool
	edges      []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	contracted bool
	head       *Vertex
	tail       *Vertex
}
