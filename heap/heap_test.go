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

// func TestRanged2Sum(t *testing.T) {
// 	f, _ := os.Open("../data/Ranged2Sum.txt")
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	array := make([]int, 0, 0)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		x, _ := strconv.Atoi(line)
// 		array = append(array, x)
// 	}

// 	total := 0
// 	keeper := make(map[int]int)
// 	tKeeper := make(map[int]bool)
// 	for i := 0; i < len(array); i++ {
// 		x := array[i]

// 		for k := range keeper {
// 			if x != k && k >= -10000-x && k < 10000-x {
// 				_, exist := tKeeper[x+k]
// 				if !exist {
// 					tKeeper[x+k] = true
// 					total++
// 				}
// 			}
// 		}

// 		keeper[x] = i
// 	}

// 	fmt.Println(total)
// }

// func TestRanged2Sum(t *testing.T) {
// 	f, _ := os.Open("../data/Ranged2Sum.txt")
// 	defer f.Close()

// 	scanner := bufio.NewScanner(f)
// 	array := make([]int, 0, 0)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		x, _ := strconv.Atoi(line)
// 		array = append(array, x)
// 	}

// 	tKeeper := make(map[int]bool)
// 	total := 0
// 	for i := 0; i < len(array); i++ {
// 		for j := i + 1; j < len(array); j++ {
// 			if array[i] != array[j] {
// 				if array[i]+array[j] <= 10000 && array[i]+array[j] >= -10000 {
// 					_, exist := tKeeper[array[i]+array[j]]
// 					if !exist {
// 						tKeeper[array[i]+array[j]] = true
// 						total++
// 					}
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println(total)
// 	// Answer should be 427
// }
