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
}
