package mwis

// Mwis calculates the maximum-weight independent set
func Mwis(weights []int) (cache []int) {
	cache = make([]int, len(weights))
	cache[0], cache[1] = 0, weights[1]

	for i := 2; i < len(cache); i++ {
		cache[i] = max(cache[i-1], cache[i-2]+weights[i])
	}

	return cache
}

// Reconstruct recovers the maximum-weight independent set from cache
func Reconstruct(cache []int, weights []int) (included []bool) {
	included = make([]bool, len(cache))

	i := len(included) - 1
	for i >= 2 {
		if cache[i-1] > cache[i-2]+weights[i] {
			i = i - 1
		} else {
			included[i] = true
			i = i - 2
		}
	}

	if included[2] == false {
		included[1] = true
	}

	return included
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
