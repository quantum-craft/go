package scc

import (
	"fmt"
	"testing"

	"github.com/quantum-craft/go/sorting"
)

func TestKosaraju(t *testing.T) {
	// file := "../data/SCC.txt"
	// file := "../data/KosarajuTestSmallEdges.txt"
	file := "../data/KosarajuTestMediumEdges.txt"

	leaderCount := Kosaraju(file)

	leaderCount = sorting.MergeSort(leaderCount)

	fmt.Println(leaderCount[len(leaderCount)-1])
	fmt.Println(leaderCount[len(leaderCount)-2])
	fmt.Println(leaderCount[len(leaderCount)-3])
	fmt.Println(leaderCount[len(leaderCount)-4])
	fmt.Println(leaderCount[len(leaderCount)-5])
}

func testAllTrue(slice []bool) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == false {
			return false
		}
	}

	return true
}
