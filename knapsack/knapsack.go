package knapsack

func knapsack(values []int, weights []int, W int) int {
	n := len(values)

	V := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		V[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if j < weights[i-1] {
				V[i][j] = V[i-1][j]
			} else {
				V[i][j] = max(V[i-1][j], V[i-1][j-weights[i-1]]+values[i-1])
			}
		}
	}

	return V[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
