package knapsack

type item struct {
	value  int
	weight int
}

var items []item
var table map[[2]int]int

func knapsack(input []item, W int) int {
	if W == 0 {
		return 0
	}

	n := len(input)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i < len(dp); i++ {
		for w := 1; w <= W; w++ {
			item := input[i-1]
			dp[i][w] = dp[i-1][w]

			if w-item.weight >= 0 {
				dp[i][w] = max(dp[i][w], dp[i-1][w-item.weight]+item.value)
			}
		}
	}

	return dp[len(input)][W]
}

func knapsack2(input []item, W int) int {
	items = input
	table = make(map[[2]int]int)
	n := len(items)

	return V(n, W)
}

// V is our simulated table V[i][x]
func V(i, x int) int {
	if i == 0 {
		return 0
	}

	wi, vi := items[i-1].weight, items[i-1].value

	if x < wi {
		return V(i-1, x)
	}

	v, ok := table[[2]int{i, x}]
	if !ok {
		table[[2]int{i, x}] = max(V(i-1, x), V(i-1, x-wi)+vi)
		v = table[[2]int{i, x}]
	}

	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
