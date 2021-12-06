package twoSat

func TwoSat(n int, clauses [][]int) bool {
	graph := make(map[int][]int)
	graphRev := make(map[int][]int)

	for _, c := range clauses {
		a := c[0]
		b := c[1]

		if a > 0 && b > 0 {
			// !a => b
			graph[a+n] = append(graph[a+n], b)
			graphRev[b] = append(graphRev[b], a+n)
			// !b => a
			graph[b+n] = append(graph[b+n], a)
			graphRev[a] = append(graphRev[a], b+n)
		} else if a < 0 && b > 0 {
			a = -a
			// a => b
			graph[a] = append(graph[a], b)
			graphRev[b] = append(graphRev[b], a)
			// !b => !a
			graph[b+n] = append(graph[b+n], a+n)
			graphRev[a+n] = append(graphRev[a+n], b+n)
		} else if a > 0 && b < 0 {
			b = -b
			// !a => !b
			graph[a+n] = append(graph[a+n], b+n)
			graphRev[b+n] = append(graphRev[b+n], a+n)
			// b => a
			graph[b] = append(graph[b], a)
			graphRev[a] = append(graphRev[a], b)
		} else {
			a = -a
			b = -b
			// a => !b
			graph[a] = append(graph[a], b+n)
			graphRev[b+n] = append(graphRev[b+n], a)
			// b => !a
			graph[b] = append(graph[b], a+n)
			graphRev[a+n] = append(graphRev[a+n], b)
		}
	}

	explored := make([]bool, 2*n+1)
	times := make([]int, 2*n+1)
	t := 0
	for i := 1; i <= 2*n; i++ {
		if !explored[i] {
			dfsTopology(i, graphRev, explored, &t, times)
		}
	}

	explored = make([]bool, 2*n+1)
	roots := make([]int, 2*n+1)
	for t := 2 * n; t >= 1; t-- {
		if !explored[times[t]] {
			dfsScc(times[t], graph, explored, times[t], roots)
		}
	}

	for i := 1; i <= n; i++ {
		if roots[i] == roots[i+n] {
			return false
		}
	}

	return true
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
