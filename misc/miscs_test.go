package misc

import "testing"

func TestFind2ndLargest(t *testing.T) {
	xs := []int{5, 4, 2, 3, 1, 8, 7, 6}

	ans := Find2ndLargest(xs)

	if ans != 7 {
		t.Error("Find2ndLargest error !")
	}
}
