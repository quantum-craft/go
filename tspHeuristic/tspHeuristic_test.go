package tspHeuristic

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestTspHeuristic(t *testing.T) {
	// points := ReadFile("../data/tspLarge.txt")

	// ans := TspHeuristic(points)

	// fmt.Println(ans)
}

func ReadFile(filePath string) map[int][]float64 {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	n := -1

	points := make(map[int][]float64)

	for scanner.Scan() {
		line := scanner.Text()

		if n == -1 {
			strs := strings.Split(line, " ")
			n, _ = strconv.Atoi(strs[0])
		} else {
			strs := strings.Split(line, " ")
			idx, _ := strconv.Atoi(strs[0])
			x, _ := strconv.ParseFloat(strs[1], 64)
			y, _ := strconv.ParseFloat(strs[2], 64)

			points[idx] = []float64{x, y}
		}
	}

	return points
}
