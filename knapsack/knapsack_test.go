package knapsack

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestKnapsackSmall(t *testing.T) {
	f, _ := os.Open("../data/knapsack1.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	values := []int{}
	weights := []int{}
	W := 0
	n := 0

	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if k == 0 {
			W, _ = strconv.Atoi(fields[0])
			n, _ = strconv.Atoi(fields[1])

			k++
		} else {
			v, _ := strconv.Atoi(fields[0])
			w, _ := strconv.Atoi(fields[1])

			values = append(values, v)
			weights = append(weights, w)
		}
	}

	if len(values) != n {
		t.Error("Input file error !")
	}

	if len(weights) != n {
		t.Error("Input file error !")
	}

	if knapsack(values, weights, W) != 2493893 {
		t.Error("TestKnapsackSmall error !")
	}
}

func TooLongTestKnapsackBig(t *testing.T) {
	f, _ := os.Open("../data/knapsack_big.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	values := []int{}
	weights := []int{}
	W := 0
	n := 0

	k := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if k == 0 {
			W, _ = strconv.Atoi(fields[0])
			n, _ = strconv.Atoi(fields[1])

			k++
		} else {
			v, _ := strconv.Atoi(fields[0])
			w, _ := strconv.Atoi(fields[1])

			values = append(values, v)
			weights = append(weights, w)
		}
	}

	if len(values) != n {
		t.Error("Input file error !")
	}

	if len(weights) != n {
		t.Error("Input file error !")
	}

	if knapsack(values, weights, W) != 4243395 {
		t.Error("TestKnapsackSmall error !")
	}
}
