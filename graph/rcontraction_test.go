package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRContraction(t *testing.T) {
	f, err := os.Open("../data/kargerMinCut.txt")
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	rd := bufio.NewReader(f)
	line, err := rd.ReadString('\n')

	numbers := make([]int, 0, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, "\n"))

		numbers = append(numbers, i)

		line, err = rd.ReadString('\n')
	}
}
