package utils

// SliceEqual tests equality of two slices
func SliceEqual(xs []int, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}

// SliceIncreasing tests if the slice is incremental by one
func SliceIncreasing(xs []int) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1]-xs[i] != 1 {
			return false
		}
	}

	return true
}

// SliceDecreasing tests if the slice is decremental by one
func SliceDecreasing(xs []int) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1]-xs[i] != -1 {
			return false
		}
	}

	return true
}
