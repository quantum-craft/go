package sort

// MergeSort is a basic nlgn sorting algorithm
func MergeSort(xs []int) []int {
	if len(xs) <= 1 {
		return xs
	}

	left := MergeSort(xs[0 : len(xs)/2])
	right := MergeSort(xs[len(xs)/2:])
	return Merge(left, right)
}

// Merge will merge two sorted arrays into one sorted array
func Merge(xs []int, ys []int) []int {
	res := make([]int, len(xs)+len(ys))

	i, j, k := 0, 0, 0
	for i < len(xs) && j < len(ys) {
		if xs[i] < ys[j] {
			res[k] = xs[i]
			i++
		} else {
			res[k] = ys[j]
			j++
		}

		k++
	}

	if i != len(xs) {
		copy(res[k:], xs[i:])
	} else if j != len(ys) {
		copy(res[k:], ys[j:])
	}

	return res
}
