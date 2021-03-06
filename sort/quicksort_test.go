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
	return a.d[i] < a.d[j]
}

func (a ConcreteData) Get(i int) interface{} {
	return a.d[i]
}

func (a ConcreteData) Set(i int, data interface{}) {
	a.d[i] = data.(int)
}

func TestQuickSortSmall(t *testing.T) {
	xs := ConcreteData{d: []int{5, 4, 2, 3, 1, 8, 7, 6}}
	ys := ConcreteData{d: []int{1, 2, 3, 4, 5, 6, 7, 8}}

	QuickSort(xs)

	if !SliceEqual(xs.d, ys.d) {
		t.Error("QuickSort error !")
	}

	if xs.Get(0) != 1 {
		t.Error("QuickSort Get error !")
	}

	if xs.Get(2) != 3 {
		t.Error("QuickSort Get error !")
	}

	if xs.Get(xs.Len()-1) != 8 {
		t.Error("QuickSort Get error !")
	}

	xs.Set(0, 182)
	xs.Set(1, 182)
	xs.Set(2, 182)
	xs.Set(3, 182)
	xs.Set(xs.Len()-1, 182)

	if xs.Get(0) != 182 {
		t.Error("QuickSort Set error !")
	}

	if xs.Get(1) != 182 {
		t.Error("QuickSort Set error !")
	}

	if xs.Get(2) != 182 {
		t.Error("QuickSort Set error !")
	}

	if xs.Get(3) != 182 {
		t.Error("QuickSort Set error !")
	}

	if xs.Get(xs.Len()-1) != 182 {
		t.Error("QuickSort Set error !")
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

func TestQuickSort2Large(t *testing.T) {
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

	numbers := make([]int, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	quickSort(numbers, 0, len(numbers)-1)

	if !SliceIncreasing(numbers) {
		t.Error("QuickSort error !")
	}
}

func TestFindKthLarge(t *testing.T) {
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

	numbers := make([]int, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	kthIdx1 := findKth(numbers, 1023, 0, len(numbers)-1)
	kthIdx2 := findKth(numbers, 2, 0, len(numbers)-1)
	kthIdx3 := findKth(numbers, len(numbers)/2, 0, len(numbers)-1)

	if kthIdx1 != 1022 || numbers[kthIdx1] != 1022 {
		t.Error("findKth error !")
	}

	if kthIdx2 != 1 || numbers[kthIdx2] != 1 {
		t.Error("findKth error !")
	}

	if kthIdx3 != len(numbers)/2-1 || numbers[kthIdx3] != 4999 {
		t.Error("findKth error !")
	}
}
