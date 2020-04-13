package misc

import (
	"github.com/cznic/mathutil"
)

// Find2ndLargest finds the second-largest number in the array,
// and that uses at most n + lg n - 2 comparisons.
// 1.
func Find2ndLargest(a []int) int {
	cnt := len(a) - 1
	larges := make([]int, 0, cnt)
	smalls := make([]int, 0, cnt)
	backs := make([]int, 0, cnt)

	depth := mathutil.Log2Uint32(uint32(len(a)))

	splitTarget := a
	prevBase := 0
	for j, base, length := 0, 0, len(a)/2; j < depth; j, base, length = j+1, base+length, length/2 {
		largePiece, smallPiece, backPiece := splitLargeSmall(splitTarget, prevBase)
		larges = append(larges, largePiece...)
		smalls = append(smalls, smallPiece...)
		backs = append(backs, backPiece...)

		splitTarget = larges[base : base+length]
		prevBase = base
	}

	seconds := make([]int, depth)
	seconds[0] = smalls[cnt-1]
	backIdx := backs[cnt-1]
	for i := 1; i < depth; i++ {
		seconds[i] = smalls[backIdx]
		backIdx = backs[backIdx]
	}

	return findMax(seconds)
}

func findMax(a []int) int {
	max := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}

func splitLargeSmall(a []int, prevBase int) ([]int, []int, []int) {
	larges := make([]int, len(a)/2)
	smalls := make([]int, len(a)/2)
	backs := make([]int, len(a)/2)

	for i, k := 0, 0; i < len(a); i, k = i+2, k+1 {
		if a[i] > a[i+1] {
			larges[k] = a[i]
			smalls[k] = a[i+1]
			backs[k] = i + prevBase
		} else {
			larges[k] = a[i+1]
			smalls[k] = a[i]
			backs[k] = i + 1 + prevBase
		}
	}

	return larges, smalls, backs
}

// ComputeModalElement computes the maximum element of unimodal array (elements are distinct)
// in O(lg n) time
// 2.
func ComputeModalElement(a []int) int {
	n := len(a)

	if n == 1 {
		return a[0]
	}

	if n == 2 {
		if a[0] > a[1] {
			return a[0]
		}

		return a[1]
	}

	mid := n / 2

	if a[mid] > a[mid+1] {
		return ComputeModalElement(a[0 : mid+1])
	}

	return ComputeModalElement(a[mid+1 : n])
}

// ValueEqIndex checks whether or not there is an index i such that A[i] = i
// 3.
func ValueEqIndex(a []int, idx []int) (bool, int) {
	if len(a) == 0 {
		return false, -1
	}

	if len(a) == 1 {
		return a[0] == idx[0], idx[0]
	}

	var middleI = len(a) / 2

	if a[middleI] > idx[middleI] {
		return ValueEqIndex(a[0:middleI], idx[0:middleI])
	}

	if a[middleI] < idx[middleI] {
		return ValueEqIndex(a[middleI+1:], idx[middleI+1:len(a)])
	}

	return true, idx[middleI]
}

// ComputeLocalMinimum computes one local minimum in a n by n grid
// 4.
func ComputeLocalMinimum(fullArr [][]int) []int {
	coords := makeCoordinate(fullArr)

	return computeLocalMinimum(fullArr, coords)
}

func computeLocalMinimum(fullArr [][]int, coords [][]coordinate) []int {
	if len(coords[0]) == 1 {
		list := compareList(coords[0][0], fullArr)
		value := fullArr[coords[0][0].x][coords[0][0].y]

		ret := []int{}

		if smallerThanList(value, list, fullArr) {
			ret = append(ret, value)
		}

		return ret
	}

	e1, o1 := splitEvenOdd(coords)
	e2, o2 := splitOddEven(coords)

	even1 := computeLocalMinimum(fullArr, e1)
	odd1 := computeLocalMinimum(fullArr, o1)
	even2 := computeLocalMinimum(fullArr, e2)
	odd2 := computeLocalMinimum(fullArr, o2)

	ret := []int{}
	all := [][]int{
		even1,
		odd1,
		even2,
		odd2,
	}

	for _, m := range all {
		ret = append(ret, m...)
	}

	return ret
}

// func reallyComputeLocalMinimum(fullArr [][]int, coords [][]coordinate) []int {
// 	ret := make([]int, 0, 0)

// 	for i := range coords {
// 		for j := range coords[i] {
// 			list := compareList(coords[i][j], fullArr)
// 			value := fullArr[coords[i][j].x][coords[i][j].y]

// 			if smallerThanList(value, list, fullArr) {
// 				ret = append(ret, value)
// 			}
// 		}
// 	}

// 	return ret
// }

type coordinate struct {
	x int
	y int
}

func smallerThanList(value int, list []coordinate, fullArr [][]int) bool {
	for k := range list {
		compareValue := fullArr[list[k].x][list[k].y]

		if value > compareValue {
			return false
		}
	}

	return true
}

func splitEvenOdd(a [][]coordinate) (e [][]coordinate, o [][]coordinate) {
	e = make([][]coordinate, len(a)/2)
	o = make([][]coordinate, len(a)/2)

	for i := 0; i < len(a)/2; i++ {
		e[i] = make([]coordinate, len(a[0])/2)
		o[i] = make([]coordinate, len(a[0])/2)
	}

	for i := 0; i < len(a); i = i + 2 {
		for j := 0; j < len(a[0]); j = j + 2 {
			o[i/2][j/2] = a[i][j]
		}
	}

	for i := 0; i < len(a); i = i + 2 {
		for j := 0; j < len(a[0]); j = j + 2 {
			e[i/2][j/2] = a[i+1][j+1]
		}
	}

	return e, o
}

func splitOddEven(a [][]coordinate) (o [][]coordinate, e [][]coordinate) {
	e = make([][]coordinate, len(a)/2)
	o = make([][]coordinate, len(a)/2)

	for i := 0; i < len(a)/2; i++ {
		e[i] = make([]coordinate, len(a[0])/2)
		o[i] = make([]coordinate, len(a[0])/2)
	}

	for i := 0; i < len(a); i = i + 2 {
		for j := 0; j < len(a[0]); j = j + 2 {
			o[i/2][j/2] = a[i][j+1]
		}
	}

	for i := 0; i < len(a); i = i + 2 {
		for j := 0; j < len(a[0]); j = j + 2 {
			e[i/2][j/2] = a[i+1][j]
		}
	}

	return o, e
}

func makeCoordinate(a [][]int) [][]coordinate {
	coords := make([][]coordinate, len(a))

	for i := range a {
		coords[i] = make([]coordinate, len(a[0]))

		for j := range a[i] {
			coords[i][j] = coordinate{x: i, y: j}
		}
	}

	return coords
}

func compareList(coord coordinate, fullArr [][]int) []coordinate {
	dx := len(fullArr)
	dy := len(fullArr[0])

	coords := make([]coordinate, 0, 4)

	if coord.y-1 >= 0 {
		coords = append(coords, coordinate{x: coord.x, y: coord.y - 1}) // up
	}

	if coord.y+1 < dy {
		coords = append(coords, coordinate{x: coord.x, y: coord.y + 1}) // down
	}

	if coord.x-1 >= 0 {
		coords = append(coords, coordinate{x: coord.x - 1, y: coord.y}) // left
	}

	if coord.x+1 < dx {
		coords = append(coords, coordinate{x: coord.x + 1, y: coord.y}) // right
	}

	return coords
}
