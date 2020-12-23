package sort

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func quickSort(xs []int, low, high int) {
	if high-low+1 <= 1 {
		return
	}

	pivotPos := partition(xs, low+r.Intn(high-low+1), low, high)
	quickSort(xs, low, pivotPos-1)
	quickSort(xs, pivotPos+1, high)
}

func partition(xs []int, pIdx int, low, high int) int {
	swap(xs, low, pIdx)

	i := low
	for j := low; j <= high; j++ {
		if xs[j] < xs[low] {
			swap(xs, i+1, j)
			i++
		}
	}

	swap(xs, low, i)

	return i
}

func swap(xs []int, i, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}
