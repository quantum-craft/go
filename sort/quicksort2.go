package sort

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func quickSort(xs []int) {
	if len(xs) <= 1 {
		return
	}

	pivotPos := partition(xs, r.Intn(len(xs)))
	quickSort(xs[0:pivotPos])
	quickSort(xs[pivotPos+1 : len(xs)])
}

func partition(xs []int, pivotIdx int) int {
	if len(xs) <= 1 {
		return 0
	}

	swap(xs, 0, pivotIdx)

	i := 0
	for j := 1; j < len(xs); j++ {
		if xs[j] < xs[0] {
			swap(xs, i+1, j)
			i++
		}
	}

	swap(xs, 0, i)

	return i
}

func swap(xs []int, i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}
