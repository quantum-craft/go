package mwis

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestMwis(t *testing.T) {
	f, _ := os.Open("../data/mwis.txt")
	defer f.Close()

	numVertices := -1
	weights := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.Atoi(line)

		if numVertices == -1 {
			numVertices = i
			weights = append(weights, i)
		} else {
			weights = append(weights, i)
		}
	}

	if len(weights) != numVertices+1 {
		t.Error("Reading file of Mwis has error")
	}

	cache := Mwis(weights)
	included := Reconstruct(cache, weights)

	if included[1] != true {
		t.Error("Mwis has error !")
	}

	if included[2] != false {
		t.Error("Mwis has error !")
	}

	if included[3] != true {
		t.Error("Mwis has error !")
	}

	if included[4] != false {
		t.Error("Mwis has error !")
	}

	if included[17] != false {
		t.Error("Mwis has error !")
	}

	if included[117] != true {
		t.Error("Mwis has error !")
	}

	if included[517] != true {
		t.Error("Mwis has error !")
	}

	if included[997] != false {
		t.Error("Mwis has error !")
	}
}
