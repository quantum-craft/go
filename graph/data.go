package graph

// IVertex is the interface for operations on Vertex
type IVertex interface {
	GetEdges() []*Edge
	GetIdx() int
	GetDist() uint64
}

// Vertex is an element of V of a Graph G(V, E)
type Vertex struct {
	idx        int
	contracted bool
	explored   bool
	dist       uint64
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

// GetDist returns the distance from some starting vertex in the vertices array
// Can be infinite (^uint64(0)) if the starting vertex has not been assigned
func (v *Vertex) GetDist() uint64 {
	return v.dist
}
