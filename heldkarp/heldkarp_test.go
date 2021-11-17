package heldkarp

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestHeldKarp(t *testing.T) {
	points := ReadFile("../data/tspSmall.txt")

	ans := heldkarp(points)

	if int(math.Floor(ans)) != 8387 {
		t.Error("HeldKarp error !")
	}

	fmt.Println(ans)
}

func ReadFile(filePath string) [][]float64 {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	n := -1

	points := make([][]float64, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if n == -1 {
			strs := strings.Split(line, " ")
			n, _ = strconv.Atoi(strs[0])
		} else {
			strs := strings.Split(line, " ")
			x, _ := strconv.ParseFloat(strs[0], 64)
			y, _ := strconv.ParseFloat(strs[1], 64)

			points = append(points, []float64{x, y})
		}
	}

	return points
}
