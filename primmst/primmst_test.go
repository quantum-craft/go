package primmst

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDiffScheduling(t *testing.T) {
	f, _ := os.Open("../data/jobs.txt")
	defer f.Close()

	var weights, lengths []int

	scanner := bufio.NewScanner(f)
	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 1 {
			size, _ := strconv.Atoi(fields[0])
			weights = make([]int, size, size)
			lengths = make([]int, size, size)
		} else {
			w, _ := strconv.Atoi(fields[0])
			l, _ := strconv.Atoi(fields[1])
			weights[k] = w
			lengths[k] = l
			k++
		}
	}

	quickSort(weights, lengths, largerDiff)

	completionTimes := make([]int, len(weights), len(weights))
	completionTimes[0] = lengths[0]
	sum := weights[0] * completionTimes[0]
	for i := 1; i < len(weights); i++ {
		completionTimes[i] = lengths[i] + completionTimes[i-1]
		sum += weights[i] * completionTimes[i]
	}

	if sum != 69119377652 {
		t.Error("TestDiffScheduling error !")
	}
}

func TestRatioScheduling(t *testing.T) {
	f, _ := os.Open("../data/jobs.txt")
	defer f.Close()

	var weights, lengths []int

	scanner := bufio.NewScanner(f)
	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 1 {
			size, _ := strconv.Atoi(fields[0])
			weights = make([]int, size, size)
			lengths = make([]int, size, size)
		} else {
			w, _ := strconv.Atoi(fields[0])
			l, _ := strconv.Atoi(fields[1])
			weights[k] = w
			lengths[k] = l
			k++
		}
	}

	quickSort(weights, lengths, largerRatio)

	completionTimes := make([]int, len(weights), len(weights))
	completionTimes[0] = lengths[0]
	sum := weights[0] * completionTimes[0]
	for i := 1; i < len(weights); i++ {
		completionTimes[i] = lengths[i] + completionTimes[i-1]
		sum += weights[i] * completionTimes[i]
	}

	if sum != 67311454237 {
		t.Error("TestDiffRatio error !")
	}
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func quickSort(weights []int, lengths []int, larger func(int, int, int, int) bool) {
	if len(weights) <= 1 {
		return
	}

	pivotPos := partition(weights, lengths, r.Intn(len(weights)), larger)
	quickSort(weights[0:pivotPos], lengths[0:pivotPos], larger)
	quickSort(weights[pivotPos+1:], lengths[pivotPos+1:], larger)
}

func partition(weights []int, lengths []int, pivotIdx int, larger func(int, int, int, int) bool) int {
	if len(weights) <= 1 {
		return 0
	}

	swap(weights, lengths, 0, pivotIdx)

	i := 0
	for j := 1; j < len(weights); j++ {
		if larger(weights[j], lengths[j], weights[0], lengths[0]) {
			swap(weights, lengths, i+1, j)
			i++
		}
	}

	swap(weights, lengths, 0, i)

	return i
}

func swap(weights []int, lengths []int, thisIdx int, thatIdx int) {
	weights[thisIdx], weights[thatIdx] = weights[thatIdx], weights[thisIdx]
	lengths[thisIdx], lengths[thatIdx] = lengths[thatIdx], lengths[thisIdx]
}

func largerDiff(w1, l1, w2, l2 int) bool {
	if (w1 - l1) > (w2 - l2) {
		return true
	} else if (w1-l1) == (w2-l2) && w1 > w2 {
		return true
	}

	return false
}

func largerRatio(w1, l1, w2, l2 int) bool {
	if (float64(w1) / float64(l1)) > (float64(w2) / float64(l2)) {
		return true
	}

	return false
}
