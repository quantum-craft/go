package kruskalmst

import (
	"math/rand"
	"strconv"
	"time"
)

// Vertex is used for undirected graph in Kruskal's mst algorithm
type Vertex struct {
	VIdx        int
	GroupLeader *Vertex
	GroupSize   int
	Added       bool
	Code        uint32 // Hamming Code
	// Edges       []*Edge
}

// Edge is used for undirected graph in Kruskal's mst algorithm
type Edge struct {
	VertIdx [2]int
	Cost    int
}

func groupingHammingDist(vertices []Vertex, weights []uint32, bookKeeper map[uint32][]int, groupCnt *int) {
	for i := 0; i < len(vertices); i++ {
		for k := 0; k < len(weights); k++ {
			idxs, ok := bookKeeper[vertices[i].Code^weights[k]]
			if ok == true {
				for j := 0; j < len(idxs); j++ {
					if idxs[j] != i {
						grouping(vertices, i, idxs[j], groupCnt)
					}
				}
			}
		}
	}
}

func grouping(vertices []Vertex, idx0, idx1 int, groupCnt *int) {
	vert0, vert1 := &vertices[idx0], &vertices[idx1]
	leader0, leader1 := findLeader(vert0), findLeader(vert1)

	if leader0 != leader1 {
		if leader0.GroupSize > leader1.GroupSize {
			bigEatsSmall(leader0, leader1)
		} else {
			bigEatsSmall(leader1, leader0)
		}

		*groupCnt--
	}
}

// MaxHammingClustering is very similar to KruskalMST
// Used to calculate k-clustering with expected min-distance of groups
// Expected min-distance is 3 for this func
func MaxHammingClustering(vertices []Vertex, bookKeeper map[uint32][]int) int {
	groupCnt := len(vertices)

	var weightZeros [1]uint32
	weightZeros[0] = uint32(0)

	var weightOnes [24]uint32
	for i := 0; i < 24; i++ {
		weightOnes[i] = uint32(1 << i)
	}

	var weightTwos [276]uint32 // 24 choose 2
	k := 0
	for i := 0; i < 24; i++ {
		for j := i + 1; j < 24; j++ {
			weightTwos[k] = uint32(1<<i) | uint32(1<<j)
			k++
		}
	}

	weights := make([]uint32, 0)
	weights = append(weights, weightZeros[:]...)
	weights = append(weights, weightOnes[:]...)
	weights = append(weights, weightTwos[:]...)

	// grouping Hamming distance of zero, one, and two
	groupingHammingDist(vertices, weights[:], bookKeeper, &groupCnt)

	return groupCnt
}

func hammingDist(x uint32, y uint32) int {
	return bitCount(x ^ y)
}

func bitCount(x uint32) int {
	cnt := 0

	for x != 0 {
		x = x & (x - 1)
		cnt++
	}

	return cnt
}

func streamToUint(s []string) uint32 {
	var acc uint32 = 0

	for i := 0; i < len(s); i++ {
		digit, _ := strconv.Atoi(s[i])
		acc = acc*2 + uint32(digit)
	}

	return acc
}

// MaxSpacingClustering is very similar to KruskalMST
// Used to calculate k-clustering with maximized min-distance of groups
func MaxSpacingClustering(vertices []Vertex, edges []Edge, k int) int {
	groupCnt := len(vertices)

	quickSort(edges)

	for i := 0; i < len(edges); i++ {
		vert0, vert1 := &vertices[edges[i].VertIdx[0]], &vertices[edges[i].VertIdx[1]]
		leader0, leader1 := findLeader(vert0), findLeader(vert1)

		if leader0 != leader1 {
			if groupCnt != k {
				if leader0.GroupSize > leader1.GroupSize {
					bigEatsSmall(leader0, leader1)
				} else {
					bigEatsSmall(leader1, leader0)
				}

				groupCnt--
			} else {
				return edges[i].Cost
			}
		}
	}

	return -1
}

// KruskalMST using Kruskal's minimum spanning tree algorithm to find mst cost
func KruskalMST(vertices []Vertex, edges []Edge) int {
	minCost, vertCnt := 0, 0

	quickSort(edges)

	for i := 0; i < len(edges); i++ {
		if vertCnt == len(vertices) {
			break
		}

		vert0, vert1 := &vertices[edges[i].VertIdx[0]], &vertices[edges[i].VertIdx[1]]
		leader0, leader1 := findLeader(vert0), findLeader(vert1)

		if leader0 != leader1 {
			minCost += edges[i].Cost

			if vert0.Added == false {
				vert0.Added = true
				vertCnt++
			}

			if vert1.Added == false {
				vert1.Added = true
				vertCnt++
			}

			if leader0.GroupSize > leader1.GroupSize {
				bigEatsSmall(leader0, leader1)
			} else {
				bigEatsSmall(leader1, leader0)
			}
		}
	}

	return minCost
}

func bigEatsSmall(big *Vertex, small *Vertex) {
	small.GroupLeader = big
	big.GroupSize += small.GroupSize
}

func findLeader(vert *Vertex) *Vertex {
	v := vert

	for v.GroupLeader.VIdx != v.VIdx {
		v = v.GroupLeader
	}

	return v
}

func findGroupSize(vert *Vertex) int {
	return findLeader(vert).GroupSize
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func quickSort(xs []Edge) {
	if len(xs) <= 1 {
		return
	}

	pivotPos := partition(xs, r.Intn(len(xs)))
	quickSort(xs[0:pivotPos])
	quickSort(xs[pivotPos+1:])
}

func partition(xs []Edge, pivotIdx int) int {
	if len(xs) <= 1 {
		return 0
	}

	swap(xs, 0, pivotIdx)

	i := 0
	for j := 1; j < len(xs); j++ {
		if xs[j].Cost < xs[0].Cost {
			swap(xs, i+1, j)
			i++
		}
	}

	swap(xs, 0, i)

	return i
}

func swap(xs []Edge, thisIdx int, thatIdx int) {
	xs[thisIdx], xs[thatIdx] = xs[thatIdx], xs[thisIdx]
}
