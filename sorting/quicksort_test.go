package sorting

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestQuickSortSmall(t *testing.T) {
	xs := []int{5, 4, 2, 3, 1, 8, 7, 6}
	ys := []int{1, 2, 3, 4, 5, 6, 7, 8}

	QuickSort(xs)

	if !sliceEqual(xs, ys) {
		t.Error("QuickSort error !")
	}
}

func TestQuickSortLarge(t *testing.T) {
	f, err := os.Open("../data/QuickSortNumbers.txt")
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')

	numbers := make([]int, 0, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	QuickSort(numbers)

	if !sliceIncreasing(numbers) {
		t.Error("QuickSort error !")
	}
}

func sliceEqual(xs []int, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}

func sliceIncreasing(xs []int) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1]-xs[i] != 1 {
			return false
		}
	}

	return true
}
