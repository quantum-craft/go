package knapsack

func knapsack(values []int, weights []int, W int) int {
	n := len(values)

	V := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		V[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for x := 0; x <= W; x++ {
			if x < weights[i-1] {
				V[i][x] = V[i-1][x]
			} else {
				V[i][x] = max(V[i-1][x], V[i-1][x-weights[i-1]]+values[i-1])
			}
		}
	}

	return V[n][W]
}

func sliceIncreasing(xs []int) bool {
	if len(xs) <= 1 {
		return true
	}

	for i := 0; i < len(xs)-1; i++ {
		if xs[i+1] < xs[i] {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
