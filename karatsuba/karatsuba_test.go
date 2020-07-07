package karatsuba

import (
	"testing"
)

func TestKaratsuba(t *testing.T) {
	// 3141592653589793238462643383279502884197169399375105820974944592
	// 2718281828459045235360287471352662497757247093699959574966967627
	b := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6, 2, 6, 4, 3, 3, 8, 3, 2, 7, 9, 5, 0, 2, 8, 8, 4, 1, 9, 7, 1, 6, 9, 3, 9, 9, 3, 7, 5, 1, 0, 5, 8, 2, 0, 9, 7, 4, 9, 4, 4, 5, 9, 2, 2, 7, 1, 8, 2, 8, 1, 8, 2, 8, 4, 5, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 6, 2, 4, 9, 7, 7, 5, 7, 2, 4, 7, 0, 9, 3, 6, 9, 9, 9}
	a := []int{2, 7, 1, 8, 2, 8, 1, 8, 2, 8, 4, 5, 9, 0, 4, 5, 2, 3, 5, 3, 6, 0, 2, 8, 7, 4, 7, 1, 3, 5, 2, 6, 6, 2, 4, 9, 7, 7, 5, 7, 2, 4, 7, 0, 9, 3, 6, 9, 9, 9, 5, 9, 5, 7, 4, 9, 6, 6, 9, 6, 7, 6, 2, 7, 3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 7, 9, 5, 0, 2, 8, 8, 4, 1, 9, 7, 1, 6, 9, 3, 9, 9, 3}

	karatsubaMul := Multiplication(a, b)
	normalMul := NormalMultiplication(a, b)

	if !SliceEqual(karatsubaMul, normalMul) {
		t.Error("TestKaratsuba error !")
	}
}

// SliceEqual tests equality of two slices
func SliceEqual(xs []int, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}

	return true
}
