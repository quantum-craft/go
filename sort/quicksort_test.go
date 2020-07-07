package sort

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

type ConcreteData struct {
	d []int
}

func (a ConcreteData) Swap(i, j int) {
	a.d[i], a.d[j] = a.d[j], a.d[i]
}

func (a ConcreteData) Range(i, j int) Data {
	return ConcreteData{d: a.d[i:j]}
}

func (a ConcreteData) Len() int {
	return len(a.d)
}

func (a ConcreteData) LessThan(i, j int) bool {
	if a.d[i] < a.d[j] {
		return true
	}

	return false
}

func TestQuickSortSmall(t *testing.T) {
	xs := ConcreteData{d: []int{5, 4, 2, 3, 1, 8, 7, 6}}
	ys := ConcreteData{d: []int{1, 2, 3, 4, 5, 6, 7, 8}}

	QuickSort(xs)

	if !SliceEqual(xs.d, ys.d) {
		t.Error("QuickSort error !")
	}
}

func TestQuickSortLarge(t *testing.T) {
	lineEnding := "\n"

	if runtime.GOOS == "windows" {
		lineEnding = "\r\n"
	} else {
		lineEnding = "\n"
	}

	f, err := os.Open("../data/QuickSortNumbers.txt")
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')

	numbers := ConcreteData{d: []int{}}

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers.d = append(numbers.d, i)

		line, err = r.ReadString('\n')
	}

	QuickSort(numbers)

	if !SliceIncreasing(numbers.d) {
		t.Error("QuickSort error !")
	}
}