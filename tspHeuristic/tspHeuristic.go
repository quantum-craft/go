package tspHeuristic

import "math"

func TspHeuristic(points map[int][]float64) int {
	firstCity := points[1]

	curr := points[1]
	delete(points, 1)

	ans := 0.0
	for len(points) != 0 {
		next := FindClosest(curr, points)

		ans += math.Sqrt(SquaredDist(curr, next, points))

		curr = points[next]
		delete(points, next)
	}

	ans += math.Sqrt(math.Pow(curr[0]-firstCity[0], 2.0) + math.Pow(curr[1]-firstCity[1], 2.0))

	return int(math.Floor(ans))
}

func SquaredDist(curr []float64, j int, points map[int][]float64) float64 {
	return math.Pow(curr[0]-points[j][0], 2.0) + math.Pow(curr[1]-points[j][1], 2.0)
}

func FindClosest(curr []float64, points map[int][]float64) int {
	ans := -1
	minDist := math.MaxFloat64

	for k := range points {
		dist := SquaredDist(curr, k, points)
		if dist < minDist {
			ans = k
			minDist = dist
		} else if dist == minDist && k < ans {
			ans = k
		}
	}

	return ans
}
