package graph

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix())) // package-wise var

// RContraction use Karger's randomized contraction algorithm
// to find the min-cut of an undirected graph
func RContraction(vertices []Vertex, edges []Edge) int {
	r = rand.New(rand.NewSource(time.Now().Unix())) // new seed in every run

	for i := 0; i < len(vertices)-2; i++ {
		remainingEdges := getRemainingEdges(edges)
		contracted := remainingEdges[r.Intn(len(remainingEdges))]

		contracted.head.edges = append(contracted.head.edges, contracted.tail.edges...)

		// remove tail vertex
		contracted.tail.contracted = true
		for j := 0; j < len(contracted.head.edges); j++ {
			if contracted.head.edges[j].head == contracted.tail {
				contracted.head.edges[j].head = contracted.head
			}

			if contracted.head.edges[j].tail == contracted.tail {
				contracted.head.edges[j].tail = contracted.head
			}
		}

		// remove contracted edge and self-loop
		for j := 0; j < len(contracted.head.edges); j++ {
			if contracted.head.edges[j].head == contracted.head.edges[j].tail {
				contracted.head.edges[j].contracted = true
				contracted.head.edges[j] = nil
			}
		}

		newHeadEdges := make([]*Edge, 0)
		for j := 0; j < len(contracted.head.edges); j++ {
			if contracted.head.edges[j] != nil {
				newHeadEdges = append(newHeadEdges, contracted.head.edges[j])
			}
		}
		contracted.head.edges = newHeadEdges
	}

	return len(getRemainingEdges(edges))
}

func getRemainingEdges(edges []Edge) []Edge {
	ret := make([]Edge, 0)

	for i := 0; i < len(edges); i++ {
		if edges[i].contracted == false {
			ret = append(ret, edges[i])
		}
	}

	return ret
}

func getRemainingVertices(vertices []Vertex) []Vertex {
	ret := make([]Vertex, 0)

	for i := 0; i < len(vertices); i++ {
		if vertices[i].contracted == false {
			ret = append(ret, vertices[i])
		}
	}

	return ret
}
