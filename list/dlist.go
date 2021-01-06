package list

// Node is Node
type Node struct {
	val  int
	next *Node
	prev *Node
}

// Dlist is Dlist
type Dlist struct {
	head *Node
	tail *Node
}

// NewList is NewList
func NewList() *Dlist {
	head := Node{}
	tail := Node{}
	head.next = &tail
	tail.prev = &head

	return &Dlist{head: &head, tail: &tail}
}

// Empty is Empty
func (l *Dlist) Empty() bool {
	if l.head.next == l.tail {
		return true
	}

	return false
}

// Front is Front
func (l *Dlist) Front() *Node {
	if !l.Empty() {
		return l.head.next
	}

	return nil
}

// Back is Back
func (l *Dlist) Back() *Node {
	if !l.Empty() {
		return l.tail.prev
	}

	return nil
}

// Remove is Remove
func (l *Dlist) Remove(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev

	n.prev = nil
	n.next = nil
}

// PushFront is PushFront
func (l *Dlist) PushFront(n *Node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next = n
	n.next.prev = n
}

// PushBack is PushBack
func (l *Dlist) PushBack(n *Node) {
	n.next = l.tail
	n.prev = l.tail.prev
	l.tail.prev = n
	n.prev.next = n
}
