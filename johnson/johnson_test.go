package johnson

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestJohnson(t *testing.T) {
	// filePathL := "../data/largeG.txt"
	filePath1 := "../data/g1.txt"
	filePath2 := "../data/g2.txt"
	filePath3 := "../data/g3.txt"
	filePath4 := "../data/smallG.txt"

	if JohnsonUtil(filePath3, t) != -19 {
		t.Error("TestJohnson error !")
	}

	if JohnsonUtil(filePath1, t) != math.MaxInt32 {
		t.Error("TestJohnson error !")
	}

	if JohnsonUtil(filePath2, t) != math.MaxInt32 {
		t.Error("TestJohnson error !")
	}

	if JohnsonUtil(filePath4, t) != -6 {
		t.Error("TestJohnson error !")
	}

	// if JohnsonUtil(filePathL, t) != -6 {
	// 	t.Error("TestJohnson error !")
	// }
}

func JohnsonUtil(filePath string, t *testing.T) int {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	n, m := -1, -1

	edges := make([][3]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if n == -1 {
			strs := strings.Split(line, " ")
			n, _ = strconv.Atoi(strs[0])
			m, _ = strconv.Atoi(strs[1])

		} else {
			strs := strings.Split(line, " ")
			u, _ := strconv.Atoi(strs[0])
			v, _ := strconv.Atoi(strs[1])
			w, _ := strconv.Atoi(strs[2])

			edges = append(edges, [3]int{u - 1, v - 1, w})
		}
	}

	if m != len(edges) {
		t.Error("Johnson reading file error !")
	}

	return Johnson(n, m, edges)
}
