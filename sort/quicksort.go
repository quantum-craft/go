package sort

// Data is the aggregation of data (ex: []int)
type Data interface {
	LessThan(i, j int) bool
	Len() int
	Range(i, j int) Data // Range is just like [i, j), j is not included
	Swap(i, j int)
	Get(i int) interface{}
	Set(i int, data interface{})
}

// QuickSort sorts array in-place with randomized choices of pivot
func QuickSort(xs Data) {
	if xs.Len() <= 1 {
		return
	}

	pivotPos := Partition(xs, r.Intn(xs.Len()))
	QuickSort(xs.Range(0, pivotPos))
	QuickSort(xs.Range(pivotPos+1, xs.Len()))
}

// Partition func will split the input array into two parts:
// one part contains elements which are smaller than the pivot,
// the other part contains elements which are larger than the pivot.
func Partition(xs Data, pivotIdx int) int {
	if xs.Len() <= 1 {
		return 0
	}

	xs.Swap(0, pivotIdx)

	i := 0
	for j := 1; j < xs.Len(); j++ {
		if xs.LessThan(j, 0) {
			xs.Swap(i+1, j)
			i++
		}
	}

	xs.Swap(0, i)

	return i
}
