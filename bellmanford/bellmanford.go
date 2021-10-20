package bellmanford

import (
	"math"
)

func Bellmanford(n int, edges [][3]int, src, dest int) int {
	dist := make([]int, n)
	prev := make([]int, n)

	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[src] = 0

	for i := 0; i <= n-1; i++ {
		for _, e := range edges {
			u := e[0]
			v := e[1]
			w := e[2]

			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				prev[v] = u
			}
		}
	}

	// Run one more inner loop
	for _, e := range edges {
		u := e[0]
		v := e[1]
		w := e[2]

		if dist[u]+w < dist[v] {
			// Got improvement
			// Negative cycle
			// prev array is useful
			return -1
		}
	}

	// if dist[dest] != math.MaxInt32 {
	// 	t := dest
	// 	for t != src {
	// 		fmt.Println(t)
	// 		t = prev[t]
	// 	}

	// 	fmt.Println(t)
	// }

	return dist[dest]
}
