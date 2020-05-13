package utils

// Partition func will split the input array into two parts:
// one part contains elements which are smaller than the pivot,
// the other part contains elements which are larger than the pivot.
func Partition(xs []int, pivotIdx int) int {
	if len(xs) <= 1 {
		return 0
	}

	Swap(xs, 0, pivotIdx)

	i := 0
	for j := 1; j < len(xs); j++ {
		if xs[j] <= xs[0] {
			Swap(xs, i+1, j)
			i++
		}
	}

	Swap(xs, 0, i)

	return i
}

// PartitionByte func will split the input array into two parts:
// one part contains elements which are smaller than the pivot,
// the other part contains elements which are larger than the pivot.
func PartitionByte(xs []byte, pivotIdx int) int {
	if len(xs) <= 1 {
		return 0
	}

	SwapByte(xs, 0, pivotIdx)

	i := 0
	for j := 1; j < len(xs); j++ {
		if xs[j] <= xs[0] {
			SwapByte(xs, i+1, j)
			i++
		}
	}

	SwapByte(xs, 0, i)

	return i
}
