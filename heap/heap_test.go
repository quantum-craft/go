package heap

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/quantum-craft/go/utils"
)

func TestMinHeap(t *testing.T) {
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

	numbers := make([]int, 0, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	heap := make([]int, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	for i := 0; i < len(numbers); i++ {
		InsertMinheap(minheap, numbers[i])
	}

	if GetMin(minheap) != 0 {
		t.Error("MinHeap GetMin error !")
	}

	ans := make([]int, 0)
	for n := ExtractMin(minheap); n != MinInt; n = ExtractMin(minheap) {
		ans = append(ans, n)
	}

	if !utils.SliceIncreasing(ans) {
		t.Error("MinHeapSort error !")
	}
}

func TestMaxHeap(t *testing.T) {
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

	numbers := make([]int, 0, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, lineEnding))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	heap := make([]int, 0, 0)
	lastEmpty := 0
	maxheap := MaxHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	for i := 0; i < len(numbers); i++ {
		InsertMaxheap(maxheap, numbers[i])
	}

	if GetMax(maxheap) != 10000 {
		t.Error("MaxHeap GetMax error !")
	}

	ans := make([]int, 0)
	for n := ExtractMax(maxheap); n != MaxInt; n = ExtractMax(maxheap) {
		ans = append(ans, n)
	}

	if !utils.SliceDecreasing(ans) {
		t.Error("MaxHeapSort error !")
	}
}

func TestMedianMaintenance(t *testing.T) {
	// for lower half
	heapMax := make([]int, 0, 0)
	lastEmptyMax := 0
	maxheap := MaxHeap{
		heap:      &heapMax,
		lastEmpty: &lastEmptyMax,
	}

	// for upper half
	heapMin := make([]int, 0, 0)
	lastEmptyMin := 0
	minheap := MinHeap{
		heap:      &heapMin,
		lastEmpty: &lastEmptyMin,
	}

	f, _ := os.Open("../data/Median.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	k := 0
	median := 0
	for scanner.Scan() {
		line := scanner.Text()
		score, _ := strconv.Atoi(line)

		if score < GetMax(maxheap) {
			InsertMaxheap(maxheap, score)
			k++
		} else {
			InsertMinheap(minheap, score)
			k++
		}

		if k%2 == 0 { // k is even
			if lastEmptyMax > k/2 {
				for (lastEmptyMax - k/2) != 0 {
					toMin := ExtractMax(maxheap)
					InsertMinheap(minheap, toMin)
				}
			} else if lastEmptyMax < k/2 {
				for (k/2 - lastEmptyMax) != 0 {
					toMax := ExtractMin(minheap)
					InsertMaxheap(maxheap, toMax)
				}
			}
		} else { // k is odd
			if lastEmptyMax > (k+1)/2 {
				for (lastEmptyMax - (k+1)/2) != 0 {
					toMin := ExtractMax(maxheap)
					InsertMinheap(minheap, toMin)
				}
			} else if lastEmptyMax < (k+1)/2 {
				for ((k+1)/2 - lastEmptyMax) != 0 {
					toMax := ExtractMin(minheap)
					InsertMaxheap(maxheap, toMax)
				}
			}
		}
		median += GetMax(maxheap)
	}

	if k != 10000 {
		t.Error("MedianMaintenance error !")
	}

	if median%10000 != 1213 {
		t.Error("MedianMaintenance average error !")
	}
}
