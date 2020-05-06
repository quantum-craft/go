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

// Kosaraju calculates the strongly connected components of the graph
func Kosaraju(vertices []Vertex, edges []Edge) {
	newOrders := make([]int, len(vertices))

	// first pass: reversed
	DFSLoopKosaraju(vertices, edges, true, newOrders)

	// second pass: forword
	DFSLoopKosaraju(vertices, edges, false, newOrders)
}

// DFSLoopKosaraju will calculate the finishingTime for the first time called (reversed == true)
// and assign leaders for the scond time called (reversed == false)
func DFSLoopKosaraju(vertices []Vertex, edges []Edge, reversed bool, newOrders []int) {
	MarkAllUnexplored(vertices)

	finishingTime := 0
	leader := -1

	for i := len(vertices) - 1; i >= 0; i-- {
		// first pass
		if reversed == true {
			if vertices[i].explored == false {
				DFSKosaraju(vertices, edges, i, reversed, &finishingTime, leader, newOrders)
			}
		} else { // second pass
			if vertices[newOrders[i]].explored == false {
				leader = newOrders[i]
				DFSKosaraju(vertices, edges, newOrders[i], reversed, &finishingTime, leader, newOrders)
			}
		}
	}
}

// DFSKosaraju is the DFS sub-routine for Kosaraju algorithm
func DFSKosaraju(vertices []Vertex, edges []Edge, startIdx int, reversed bool, finishingTime *int, leader int, newOrders []int) {
	v := &vertices[startIdx]
	v.explored = true

	// second pass
	if reversed == false {
		v.leader = leader
	}

	for i, neighbors := 0, v.edges; i < len(neighbors); i++ {
		if reversed == false {
			if neighbors[i].head.explored == false {
				DFSKosaraju(vertices, edges, neighbors[i].head.idx, reversed, finishingTime, leader, newOrders)
			}
		} else {
			if neighbors[i].tail.explored == false {
				DFSKosaraju(vertices, edges, neighbors[i].tail.idx, reversed, finishingTime, leader, newOrders)
			}
		}
	}

	// first pass
	if reversed == true {
		*finishingTime++
		v.finishingTime = *finishingTime
		newOrders[v.finishingTime-1] = v.idx
	}
}
