package graph

// BFS is the template of breadth-first-search
func BFS(vertices []Vertex, edges []Edge) {
	queue := NewList()

	exploredAndQueued(&vertices[0], queue)

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
