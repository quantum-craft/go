package stack

// Stack is the data keeper
type Stack struct {
	data   []int
	topIdx int
}

// MakeStack returns an empty stack
func MakeStack() Stack {
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
// Notice that data being -1 means empty
func (stack *Stack) Pop() int {
	if stack.topIdx < 0 {
		return -1 // data being -1 means empty
	}

	ret := (stack.data)[stack.topIdx]
	stack.topIdx--

	return ret
}

// PeekTop returns the top element without poping it
func (stack *Stack) PeekTop() int {
	if stack.topIdx < 0 {
		return -1 // data being -1 means empty
	}

	return stack.data[stack.topIdx]
}
