package kruskalmst

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

// TestMaxSpacing is using brute-force to solve big Hamming problem
func TestMaxSpacing() {
	f, _ := os.Open("../data/four_clustering_big.txt")
	defer f.Close()

	var numVertices, numEdges int
	var vertices []Vertex
	var edges []Edge

	k := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 2 {
			numVertices, _ = strconv.Atoi(fields[0])
			numEdges = numVertices * (numVertices - 1) / 2 // n choose 2
			vertices = make([]Vertex, numVertices)
			edges = make([]Edge, numEdges)

			fmt.Println(numEdges)
		} else {
			vertices[k].VIdx = k
			vertices[k].GroupLeader = &vertices[k] // assign self as leader
			vertices[k].GroupSize = 1              // only self
			vertices[k].Added = false
			vertices[k].Code = streamToUint(fields)

			k++
		}
	}

	k = 0
	for i := 0; i < len(vertices); i++ {
		for j := i + 1; j < len(vertices); j++ {
			edges[k].VertIdx[0] = i
			edges[k].VertIdx[1] = j
			edges[k].Cost = hammingDist(vertices[i].Code, vertices[j].Code)

			k++

			if k%10000000 == 0 {
				fmt.Println(k)
			}
		}
	}

	fmt.Println("start KruskalMST ...")

	groups := MaxHammingClustering(vertices, edges, 3)

	fmt.Println(groups)
}

// MaxHammingClustering is very similar to KruskalMST
// Used to calculate k-clustering with expected min-distance of groups
func MaxHammingClustering(vertices []Vertex, edges []Edge, spacing int) int {
	groupCnt := len(vertices)

	quickSort(edges)

	for i := 0; i < len(edges); i++ {
		vert0, vert1 := &vertices[edges[i].VertIdx[0]], &vertices[edges[i].VertIdx[1]]
		leader0, leader1 := findLeader(vert0), findLeader(vert1)

		if leader0 != leader1 {
			if edges[i].Cost != spacing {
				if leader0.GroupSize > leader1.GroupSize {
					bigEatsSmall(leader0, leader1)
				} else {
					bigEatsSmall(leader1, leader0)
				}

				groupCnt--
			} else {
				return groupCnt
			}
		}
	}

	return -1
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

var maxDepth int = 0

func findLeader(vert *Vertex) *Vertex {
	v := vert
	debug := 0

	for v.GroupLeader.VIdx != v.VIdx {
		v = v.GroupLeader
		debug++
	}

	if debug > maxDepth {
		maxDepth = debug
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
