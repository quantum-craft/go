package list

type node struct {
	val  int
	next *node
	prev *node
}

type dlist struct {
	head *node
	tail *node
}

func newList() *dlist {
	head := node{}
	tail := node{}
	head.next = &tail
	tail.prev = &head

	return &dlist{head: &head, tail: &tail}
}

func (l *dlist) empty() bool {
	if l.head.next == l.tail {
		return true
	}

	return false
}

func (l *dlist) front() *node {
	if !l.empty() {
		return l.head.next
	}

	return nil
}

func (l *dlist) back() *node {
	if !l.empty() {
		return l.tail.prev
	}

	return nil
}

func (l *dlist) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev

	n.prev = nil
	n.next = nil
}

func (l *dlist) pushFront(n *node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next = n
	n.next.prev = n
}

func (l *dlist) pushBack(n *node) {
	n.next = l.tail
	n.prev = l.tail.prev
	l.tail.prev = n
	n.prev.next = n
}
