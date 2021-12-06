package kosaraju

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestKosaraju(t *testing.T) {
	fileName := "../data/SCC.txt"

	n, edges := readFile(fileName)

	roots := Kosaraju(n, edges)

	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		counts[roots[i]]++
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
}

func readFile(fileName string) (int, [][]int) {
	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	edges := [][]int{}
	n := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		tail, _ := strconv.Atoi(fields[0])
		head, _ := strconv.Atoi(fields[1])

		n = max(max(n, tail), head)

		edges = append(edges, []int{tail, head})
	}

	return n, edges
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
