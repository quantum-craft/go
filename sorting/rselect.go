package sorting

import (
	"github.com/quantum-craft/go/utils"
)

// RSelect returns the i-th (one-base) smallest element of an array.
// i = 1 means the smallest element, i = len(xs) means the largest element
func RSelect(xs []int, n int, i int) int {
	if n == 1 {
		return xs[0]
	}

	j := utils.Partition(xs, r.Intn(n))

	if i == j+1 { // super lucky
		return xs[j]
	} else if i < j+1 { // i-th order is in the left side
		return RSelect(xs[0:j], j, i)
	} else { // i-th order is in the right side
		return RSelect(xs[j+1:], n-j-1, i-j-1)
	}
}
