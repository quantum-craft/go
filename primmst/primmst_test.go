package primmst

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDiffScheduling(t *testing.T) {
	f, _ := os.Open("../data/jobs.txt")
	defer f.Close()

	var weights, lengths []int

	scanner := bufio.NewScanner(f)
	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 1 {
			size, _ := strconv.Atoi(fields[0])
			weights = make([]int, size, size)
			lengths = make([]int, size, size)
		} else {
			w, _ := strconv.Atoi(fields[0])
			l, _ := strconv.Atoi(fields[1])
			weights[k] = w
			lengths[k] = l
			k++
		}
	}

	quickSort(weights, lengths, largerDiff)

	completionTimes := make([]int, len(weights), len(weights))
	completionTimes[0] = lengths[0]
	sum := weights[0] * completionTimes[0]
	for i := 1; i < len(weights); i++ {
		completionTimes[i] = lengths[i] + completionTimes[i-1]
		sum += weights[i] * completionTimes[i]
	}

	if sum != 69119377652 {
		t.Error("TestDiffScheduling error !")
	}
}

func TestRatioScheduling(t *testing.T) {
	f, _ := os.Open("../data/jobs.txt")
	defer f.Close()

	var weights, lengths []int

	scanner := bufio.NewScanner(f)
	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 1 {
			size, _ := strconv.Atoi(fields[0])
			weights = make([]int, size, size)
			lengths = make([]int, size, size)
		} else {
			weights[k], _ = strconv.Atoi(fields[0])
			lengths[k], _ = strconv.Atoi(fields[1])
			k++
		}
	}

	quickSort(weights, lengths, largerRatio)

	completionTimes := make([]int, len(weights), len(weights))
	completionTimes[0] = lengths[0]
	sum := weights[0] * completionTimes[0]
	for i := 1; i < len(weights); i++ {
		completionTimes[i] = lengths[i] + completionTimes[i-1]
		sum += weights[i] * completionTimes[i]
	}

	if sum != 67311454237 {
		t.Error("TestDiffRatio error !")
	}
}

func TestPrimMST(t *testing.T) {
	f, _ := os.Open("../data/edges.txt")
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
			numEdges, _ = strconv.Atoi(fields[1])

			vertices = make([]Vertex, numVertices)
			edges = make([]Edge, numEdges)
		} else {
			vidx1, _ := strconv.Atoi(fields[0])
			vidx2, _ := strconv.Atoi(fields[1])
			cost, _ := strconv.Atoi(fields[2])

			vidx1, vidx2 = vidx1-1, vidx2-1 // convert to zero based

			edges[k].VertIdx[0] = vidx1
			edges[k].VertIdx[1] = vidx2
			edges[k].Cost = cost

			vertices[vidx1].VIdx = vidx1
			vertices[vidx2].VIdx = vidx2
			vertices[vidx1].HeapIdx = -1 // not in heap yet
			vertices[vidx2].HeapIdx = -1 // not in heap yet

			if vertices[vidx1].Edges == nil {
				vertices[vidx1].Edges = make([]*Edge, 0)
			}

			if vertices[vidx2].Edges == nil {
				vertices[vidx2].Edges = make([]*Edge, 0)
			}

			vertices[vidx1].Edges = append(vertices[vidx1].Edges, &edges[k])
			vertices[vidx2].Edges = append(vertices[vidx2].Edges, &edges[k])

			k++
		}
	}

	for i := 0; i < len(vertices); i++ {
		minCost := PrimMST(vertices, edges, i)

		if minCost != -3612829 {
			t.Error("TestPrimMST error !")
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func quickSort(weights []int, lengths []int, larger func(int, int, int, int) bool) {
	if len(weights) <= 1 {
		return
	}

	pivotPos := partition(weights, lengths, r.Intn(len(weights)), larger)
	quickSort(weights[0:pivotPos], lengths[0:pivotPos], larger)
	quickSort(weights[pivotPos+1:], lengths[pivotPos+1:], larger)
}

func partition(weights []int, lengths []int, pivotIdx int, larger func(int, int, int, int) bool) int {
	if len(weights) <= 1 {
		return 0
	}

	swap(weights, lengths, 0, pivotIdx)

	i := 0
	for j := 1; j < len(weights); j++ {
		if larger(weights[j], lengths[j], weights[0], lengths[0]) {
			swap(weights, lengths, i+1, j)
			i++
		}
	}

	swap(weights, lengths, 0, i)

	return i
}

func swap(weights []int, lengths []int, thisIdx int, thatIdx int) {
	weights[thisIdx], weights[thatIdx] = weights[thatIdx], weights[thisIdx]
	lengths[thisIdx], lengths[thatIdx] = lengths[thatIdx], lengths[thisIdx]
}

func largerDiff(w1, l1, w2, l2 int) bool {
	if (w1 - l1) > (w2 - l2) {
		return true
	} else if (w1-l1) == (w2-l2) && w1 > w2 {
		return true
	}

	return false
}

func largerRatio(w1, l1, w2, l2 int) bool {
	if (float64(w1) / float64(l1)) > (float64(w2) / float64(l2)) {
		return true
	}

	return false
}
