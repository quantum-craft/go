package sorting

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestRSelect(t *testing.T) {
	xs := []int{8, 7, 6, 5, 4, 3, 2, 1}

	if RSelect(xs, len(xs), 1) != 1 {
		t.Error("RSelect error")
	}

	if RSelect(xs, len(xs), len(xs)) != 8 {
		t.Error("RSelect error")
	}

	if RSelect(xs, len(xs), len(xs)/2) != 4 {
		t.Error("RSelect error")
	}

	xs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if RSelect(xs, len(xs), 1) != 1 {
		t.Error("RSelect error")
	}

	if RSelect(xs, len(xs), len(xs)) != 10 {
		t.Error("RSelect error")
	}

	if RSelect(xs, len(xs), len(xs)/2) != 5 {
		t.Error("RSelect error")
	}

	xs = []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}

	if RSelect(xs, len(xs), 1) != -10 {
		t.Error("RSelect error")
	}

	if RSelect(xs, len(xs), len(xs)) != -1 {
		t.Error("RSelect error")
	}

	if RSelect(xs, len(xs), len(xs)/2) != -6 {
		t.Error("RSelect error")
	}

	f, err := os.Open("../data/QuickSortNumbers.txt")
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

	if RSelect(numbers, len(numbers), 1) != 0 {
		t.Error("RSelect error")
	}

	if RSelect(numbers, len(numbers), len(numbers)) != 10000 {
		t.Error("RSelect error")
	}

	if RSelect(numbers, len(numbers), len(numbers)/2) != 4999 {
		t.Error("RSelect error")
	}
}
