package kosaraju

func Kosaraju(n int, edges [][]int) []int {
	graphRev := make(map[int][]int)
	graph := make(map[int][]int)

	for _, e := range edges {
		graphRev[e[1]] = append(graphRev[e[1]], e[0])
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	explored := make([]bool, n+1)
	t := 0
	times := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if !explored[i] {
			dfsTopology(i, graphRev, explored, &t, times)
		}
	}

	explored = make([]bool, n+1)
	roots := make([]int, n+1)
	for t := n; t >= 1; t-- {
		if !explored[times[t]] {
			dfsScc(times[t], graph, explored, times[t], roots)
		}
	}

	return roots
}

func dfsTopology(i int, graph map[int][]int, explored []bool, t *int, times []int) {
	explored[i] = true

	for _, head := range graph[i] {
		if !explored[head] {
			dfsTopology(head, graph, explored, t, times)
		}
	}

	(*t)++
	times[*t] = i
}

func dfsScc(i int, graph map[int][]int, explored []bool, r int, roots []int) {
	explored[i] = true
	roots[i] = r

	for _, head := range graph[i] {
		if !explored[head] {
			dfsScc(head, graph, explored, r, roots)
		}
	}
}
