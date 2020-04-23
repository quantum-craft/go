package graph

// IVertex is the interface for operations on Vertex
type IVertex interface {
	GetEdges() []*Edge
	GetIdx() int
}

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx        int
	contracted bool
	explored   bool
	edges      []*Edge
}

// Edge is an element of E of a Graph G(V, E)
type Edge struct {
	contracted bool
	head       *Vertex
	tail       *Vertex
}

// GetEdges returns the bound edges of the bound Vertex
func (v *Vertex) GetEdges() []*Edge {
	return v.edges
}

// GetIdx returns the idx (position in the vertices array) of the Vertex
func (v *Vertex) GetIdx() int {
	return v.idx
}
