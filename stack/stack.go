package stack

const maxUint = ^uint(0)         // 1111...1
const minUint = uint(0)          // 0000...0
const maxInt = int(maxUint >> 1) // 0111...1
const minInt = -maxInt - 1       // 1000...0

// DataElement is used to store data in stack
type DataElement interface {
}

// Stack is the data keeper
type Stack struct {
	data   []DataElement
	topIdx int
}

func (stack *Stack) GetData() []DataElement {
	return stack.data
}

// MakeStack returns an empty stack
func MakeStack() Stack {
	return Stack{
		topIdx: -1,
	}
}

// Push is used to push data into stack
func (stack *Stack) Push(d DataElement) {
	if stack.topIdx == len(stack.data)-1 {
		stack.data = append(stack.data, d)
		stack.topIdx++
	} else {
		stack.topIdx++
		(stack.data)[stack.topIdx] = d
	}
}

// Pop returns the top element
// Notice that data being maxInt means empty
func (stack *Stack) Pop() DataElement {
	if stack.topIdx < 0 {
		return maxInt // data being maxInt means empty
	}

	ret := (stack.data)[stack.topIdx]
	stack.topIdx--

	return ret
}

// PeekTop returns the top element without poping it
func (stack *Stack) PeekTop() DataElement {
	if stack.topIdx < 0 {
		return maxInt // data being maxInt means empty
	}

	return stack.data[stack.topIdx]
}

// Peek2ndTop returns the second top element without poping it
func (stack *Stack) Peek2ndTop() DataElement {
	if stack.topIdx <= 0 {
		return maxInt // data being maxInt means empty
	}

	return stack.data[stack.topIdx-1]
}

// IsEmpty tell whether the stack is empty
func (stack *Stack) IsEmpty() bool {
	if stack.PeekTop() != maxInt {
		return false
	}

	return true
}
