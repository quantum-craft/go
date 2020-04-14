package sorting

import (
	"math/rand"
	"time"

	"github.com/quantum-craft/go/utils"
)

// package-wise var
var r = rand.New(rand.NewSource(time.Now().Unix()))

// QuickSort sorts array in-place with randomized choices of pivot
func QuickSort(xs []int) {
	if len(xs) <= 1 {
		return
	}

	pivotPos := utils.Partition(xs, r.Intn(len(xs)))
	QuickSort(xs[0:pivotPos])
	QuickSort(xs[pivotPos+1:])
}
