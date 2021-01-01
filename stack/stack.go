package stack

// DataElement is used to store data in stack
type DataElement interface {
}

// Stack is the data keeper
type Stack struct {
	data   []DataElement
	topIdx int
}

// NewStack returns an empty stack
func NewStack() Stack {
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
func (stack *Stack) Pop() DataElement {
	if stack.topIdx < 0 {
		return nil
	}

	ret := (stack.data)[stack.topIdx]
	stack.topIdx--

	return ret
}

// Peek returns the top element without poping it
func (stack *Stack) Peek() DataElement {
	if stack.topIdx < 0 {
		return nil
	}

	return stack.data[stack.topIdx]
}

// Peek2ndTop returns the second top element without poping it
func (stack *Stack) Peek2ndTop() DataElement {
	if stack.topIdx <= 0 {
		return nil
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
