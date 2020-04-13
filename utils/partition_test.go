package utils

import (
	"testing"
)

func TestPartition(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8}

	pivotPos := Partition(xs, 4)

	if pivotPos != 4 {
		t.Errorf("Partition a sorted array gives wrong pivot position.")
	}

	pivotPos = Partition(xs, 7)

	if pivotPos != 7 {
		t.Errorf("Partition a sorted array gives wrong pivot position.")
	}

	pivotPos = Partition(xs, 0)

	if pivotPos != 0 {
		t.Errorf("Partition a sorted array gives wrong pivot position.")
	}

	ys := []int{7, 2, 4, 5, 3, 6, 1, 8}

	pivotPos = Partition(ys, 0)

	if pivotPos != 6 {
		t.Errorf("Partition a jumbled array gives wrong pivot position.")
	}

	zs := []int{8, 7, 6, 5, 4, 3, 2, 1}

	pivotPos = Partition(zs, 0)

	if pivotPos != 7 {
		t.Errorf("Partition a jumbled array gives wrong pivot position.")
	}

	zs = []int{8, 7, 6, 5, 4, 3, 2, 1}

	pivotPos = Partition(zs, 3)

	if pivotPos != 4 {
		t.Errorf("Partition a jumbled array gives wrong pivot position.")
	}

	zs = []int{8, 7, 6, 5, 4, 3, 2, 1}

	pivotPos = Partition(zs, 5)

	if pivotPos != 2 {
		t.Errorf("Partition a jumbled array gives wrong pivot position. 2")
	}

	zs = []int{8, 7, 6, 5, 4, 3, 2, 1}

	pivotPos = Partition(zs, 2)

	if pivotPos != 5 {
		t.Errorf("Partition a jumbled array gives wrong pivot position. 2")
	}
}
