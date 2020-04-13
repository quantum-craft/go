package utils

// Swap func swap the elements indicated by two indices
func Swap(xs []int, thisIdx int, thatIdx int) {
	xs[thisIdx], xs[thatIdx] = xs[thatIdx], xs[thisIdx]
}
