package floydwarshall

import "math"

func Floydwarshall(dist [][]int, n int) int {
	ans := math.MaxInt32
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])

				ans = min(ans, dist[i][j])
			}
		}
	}

	for i := 0; i < n; i++ {
		if dist[i][i] < 0 {
			return math.MaxInt32
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
