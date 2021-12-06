package twoSat

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestTwoSat(t *testing.T) {
	n, clauses := readFile("../data/2sat1.txt")
	fmt.Println(TwoSat(n, clauses))

	n, clauses = readFile("../data/2sat2.txt")
	fmt.Println(TwoSat(n, clauses))

	n, clauses = readFile("../data/2sat3.txt")
	fmt.Println(TwoSat(n, clauses))

	n, clauses = readFile("../data/2sat4.txt")
	fmt.Println(TwoSat(n, clauses))

	n, clauses = readFile("../data/2sat5.txt")
	fmt.Println(TwoSat(n, clauses))

	n, clauses = readFile("../data/2sat6.txt")
	fmt.Println(TwoSat(n, clauses))

}

func readFile(filePath string) (n int, clauses [][]int) {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	n = -1
	clauses = make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if n == -1 {
			strs := strings.Split(line, " ")
			n, _ = strconv.Atoi(strs[0])
		} else {
			strs := strings.Split(line, " ")

			a, _ := strconv.Atoi(strs[0])
			b, _ := strconv.Atoi(strs[1])

			clauses = append(clauses, []int{a, b})
		}
	}

	return n, clauses
}
