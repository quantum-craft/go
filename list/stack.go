package list

// MaxUint is MaxUint
const MaxUint = ^uint(0)

// MinUint is MinUint
const MinUint = uint(0)

// MaxInt is MaxInt
const MaxInt = int(MaxUint >> 1)

// MinInt is MinInt
const MinInt = -MaxInt - 1

// Stack is the data keeper
type Stack struct {
	data   []int
	topIdx int
}

// NewStack returns an empty stack
func NewStack() Stack {
	return Stack{
		topIdx: -1,
	}
}

// Push is used to push data into stack
func (stack *Stack) Push(d int) {
	if stack.topIdx == len(stack.data)-1 {
		stack.data = append(stack.data, d)
		stack.topIdx++
	} else {
		stack.topIdx++
		(stack.data)[stack.topIdx] = d
	}
}

// Pop returns the top element
func (stack *Stack) Pop() int {
	if stack.topIdx < 0 {
		return MinInt
	}

	ret := (stack.data)[stack.topIdx]
	stack.topIdx--

	return ret
}

// Peek returns the top element without poping it
func (stack *Stack) Peek() int {
	if stack.topIdx < 0 {
		return MinInt
	}

	return stack.data[stack.topIdx]
}

// PeekSecond returns the second top element without poping it
func (stack *Stack) PeekSecond() int {
	if stack.topIdx <= 0 {
		return MinInt
	}

	return stack.data[stack.topIdx-1]
}

// Empty tell whether the stack is empty
func (stack *Stack) Empty() bool {
	if stack.topIdx < 0 {
		return true
	}

	return false
}

// Size gives the size of the stack
func (stack *Stack) Size() int {
	return stack.topIdx + 1
}
