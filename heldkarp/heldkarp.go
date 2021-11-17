package heldkarp

import (
	"math"
)

func heldkarp(points [][]float64) float64 {
	n := len(points)

	fullSet := (1 << n) - 1 - 1

	L := make([][]float64, fullSet+1)

	for i := range L {
		L[i] = make([]float64, n)
	}

	p := PowerSet(toIndices(toDigits(fullSet, n), n))

	subsets := make([][][]int, n)

	for i := range p {
		subsets[len(p[i])] = append(subsets[len(p[i])], p[i])
	}

	for k := 1; k <= n-1; k++ {
		L[1<<k][k] = dist(0, k, points)
	}

	for card := 2; card <= n-1; card++ {
		for _, s := range subsets[card] {
			for _, j := range s {
				sSubJ := toInt(s) & (^(1 << j))

				minimum := math.MaxFloat64
				for _, k := range s {
					if k == j {
						continue
					}

					minimum = min(minimum, L[sSubJ][k]+dist(k, j, points))
				}

				L[toInt(s)][j] = minimum
			}
		}
	}

	minimum := math.MaxFloat64
	for k := 1; k <= n-1; k++ {
		minimum = min(minimum, L[fullSet][k]+dist(k, 0, points))
	}

	return minimum
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func Combinations(L []int, r int) [][]int {
	if r == 1 {
		temp := make([][]int, 0)
		for _, rr := range L {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(L); i++ {
			perms := make([]int, 0)
			perms = append(perms, L[:i]...)
			for _, x := range Combinations(perms, r-1) {
				t := append(x, L[i])
				res = append(res, [][]int{t}...)
			}
		}
		return res
	}
}

func PowerSet(L []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i <= len(L); i++ {
		x := Combinations(L, i)
		res = append(res, x...)
	}
	return res

}

func toDigits(i int, n int) []int {
	ans := make([]int, n)
	k := n - 1
	for i > 0 {
		if i%2 != 0 {
			ans[k] = 1
		} else {
			ans[k] = 0
		}

		k--
		i /= 2
	}

	return ans
}

func toIndices(d []int, n int) []int {
	ans := []int{}

	for i := n - 1; i >= 0; i-- {
		if d[i] == 1 {
			ans = append(ans, n-i-1)
		}
	}

	return ans
}

func toInt(indices []int) int {
	ans := 0
	for i := range indices {
		ans += (1 << indices[i])
	}

	return ans
}

func dist(i, j int, points [][]float64) float64 {
	return math.Sqrt(math.Pow(points[i][0]-points[j][0], 2.0) + math.Pow(points[i][1]-points[j][1], 2.0))
}
