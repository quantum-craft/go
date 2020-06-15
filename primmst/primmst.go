package primmst

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

// Vertex is used for undirected graph in Prim's mst algorithm
type Vertex struct {
	HeapIdx int
}

// Edge is used for undirected graph in Prim's mst algorithm
type Edge struct {
	VertIdx [2]int
	Cost    int
}

func testPrim() {

}
