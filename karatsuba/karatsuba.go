package karatsuba

// NormalMultiplication will give you the multiplication
// of two numbers, by means of a smart recursive algorithm
func NormalMultiplication(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{0}
	}

	if len(a) == 1 && len(b) == 1 {
		ans := a[0] * b[0]
		ret := []int{}

		if ans/10 != 0 {
			ret = []int{ans / 10, ans - ans/10*10}
		} else {
			ret = []int{ans}
		}

		return ret
	}

	k := min(len(a)/2, len(b)/2)
	if k == 0 {
		k = 1
	}

	high := k * 2

	A := a[0 : len(a)-k]
	B := a[len(a)-k:]
	C := b[0 : len(b)-k]
	D := b[len(b)-k:]

	AC := NormalMultiplication(A, C)
	BD := NormalMultiplication(B, D)
	AD := NormalMultiplication(A, D)
	BC := NormalMultiplication(B, C)

	return add(padZeros(AC, high), add(padZeros(add(AD, BC), k), BD))
}

// Multiplication will give you the multiplication
// of two numbers, by means of Karatsuba algorithm
func Multiplication(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{0}
	}

	if len(a) == 1 && len(b) == 1 {
		ans := a[0] * b[0]
		ret := []int{}

		if ans/10 != 0 {
			ret = []int{ans / 10, ans - ans/10*10}
		} else {
			ret = []int{ans}
		}

		return ret
	}

	k := min(len(a)/2, len(b)/2)
	if k == 0 {
		k = 1
	}

	high := k * 2

	A := a[0 : len(a)-k]
	B := a[len(a)-k:]
	C := b[0 : len(b)-k]
	D := b[len(b)-k:]

	AC := Multiplication(A, C)
	BD := Multiplication(B, D)
	ApBCpD := Multiplication(add(A, B), add(C, D))
	ADpBC := sub(sub(ApBCpD, AC), BD)

	return add(padZeros(AC, high), add(padZeros(ADpBC, k), BD))
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func addSingle(a int, b int, c int) (int, int) {
	ans := a + b + c
	return ans / 10, ans - ans/10*10
}

func add(a []int, b []int) []int {
	if len(a) > len(b) {
		b = prePadZeros(b, len(a)-len(b))
	}

	if len(b) > len(a) {
		a = prePadZeros(a, len(b)-len(a))
	}

	carry := 0
	least := 0
	ans := make([]int, len(a), len(a))

	for i := len(a) - 1; i >= 0; i-- {
		carry, least = addSingle(carry, a[i], b[i])
		ans[i] = least
	}

	if carry != 0 {
		ans = append([]int{carry}, ans...)
	}

	return ans
}

func larger(a []int, b []int) bool {
	if len(a) > len(b) {
		b = prePadZeros(b, len(a)-len(b))
	}

	if len(b) > len(a) {
		a = prePadZeros(a, len(b)-len(a))
	}

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return false
		}

		if a[i] == b[i] {
			continue
		}

		if a[i] > b[i] {
			return true
		}
	}

	return false
}

func sub(a []int, b []int) []int {
	if larger(b, a) {
		return []int{-1}
	}

	if len(a) > len(b) {
		b = prePadZeros(b, len(a)-len(b))
	}

	if len(b) > len(a) {
		a = prePadZeros(a, len(b)-len(a))
	}

	carry := 0
	ans := make([]int, len(a), len(a))

	for i := len(a) - 1; i >= 0; i-- {
		digit := a[i] - b[i] + carry
		if digit < 0 {
			digit = 10 + digit
			carry = -1
		} else {
			carry = 0
		}

		ans[i] = digit
	}

	return ans
}

func prePadZeros(a []int, n int) []int {
	for i := 0; i < n; i++ {
		a = append([]int{0}, a...)
	}

	return a
}

func padZeros(a []int, n int) []int {
	for i := 0; i < n; i++ {
		a = append(a, 0)
	}

	return a
}
