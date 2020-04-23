package graph

// BFS is the template of breadth-first-search
func BFS(start *Vertex) {
	queue := NewList()

	exploredAndQueued(start, queue)

	for v := queue.Front(); v != nil; v = queue.Front() {
		for i, neighbors := 0, v.Value.GetEdges(); i < len(neighbors); i++ {
			if neighbors[i].head == v.Value {
				if neighbors[i].tail.explored == false {
					exploredAndQueued(neighbors[i].tail, queue)
				}
			} else {
				if neighbors[i].head.explored == false {
					exploredAndQueued(neighbors[i].head, queue)
				}
			}
		}

		queue.Remove(v)
	}
}

// BFSShortestPath use breadth-first-search to find the shortest path from start to end
func BFSShortestPath(start *Vertex, end *Vertex) uint64 {
	queue := NewList()

	start.dist = 0

	if start == end {
		return start.dist
	}

	exploredAndQueued(start, queue)

	for v := queue.Front(); v != nil; v = queue.Front() {
		for i, neighbors := 0, v.Value.GetEdges(); i < len(neighbors); i++ {
			if neighbors[i].head == v.Value {
				if neighbors[i].tail.explored == false {
					neighbors[i].tail.dist = v.Value.GetDist() + 1

					if neighbors[i].tail == end {
						return neighbors[i].tail.dist
					}

					exploredAndQueued(neighbors[i].tail, queue)
				}
			} else {
				if neighbors[i].head.explored == false {
					neighbors[i].head.dist = v.Value.GetDist() + 1

					if neighbors[i].head == end {
						return neighbors[i].head.dist
					}

					exploredAndQueued(neighbors[i].head, queue)
				}
			}
		}

		queue.Remove(v)
	}

	return ^uint64(0) // infinite, can not find the end vertex
}

func exploredAndQueued(v *Vertex, queue *List) {
	v.explored = true
	queue.PushBack(v)
}

// CheckAllUnexplored checks whether the vertices are all unexplored
func CheckAllUnexplored(vertices []Vertex) bool {
	for i := 0; i < len(vertices); i++ {
		if vertices[i].explored == true {
			return false
		}
	}

	return true
}

// CheckAllExplored checks whether the vertices are all explored
func CheckAllExplored(vertices []Vertex) bool {
	for i := 0; i < len(vertices); i++ {
		if vertices[i].explored == false {
			return false
		}
	}

	return true
}
