package floydwarshall

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestFloydwarshall(t *testing.T) {
	filePath1 := "../data/g1.txt"
	filePath2 := "../data/g2.txt"
	filePath3 := "../data/g3.txt"

	ans1 := FloydwarshallUtil(filePath1)
	ans2 := FloydwarshallUtil(filePath2)
	ans3 := FloydwarshallUtil(filePath3)

	ans := min(ans1, min(ans2, ans3))

	if ans != -19 {
		t.Error("TestFloydwarshall error !")
	}
}

func FloydwarshallUtil(filePath string) int {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	n := -1

	var dist [][]int

	for scanner.Scan() {
		line := scanner.Text()

		if n == -1 {
			strs := strings.Split(line, " ")
			n, _ = strconv.Atoi(strs[0])
			// m, _ = strconv.Atoi(strs[1])

			dist = make([][]int, n)
			for i := range dist {
				dist[i] = make([]int, n)

				for j := range dist[i] {
					dist[i][j] = math.MaxInt32
				}

				dist[i][i] = 0
			}

		} else {
			strs := strings.Split(line, " ")
			u, _ := strconv.Atoi(strs[0])
			v, _ := strconv.Atoi(strs[1])
			w, _ := strconv.Atoi(strs[2])

			dist[u-1][v-1] = w
		}
	}

	ans := Floydwarshall(dist, n)

	return ans
}
