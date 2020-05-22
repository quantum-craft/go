package list

// Node is the element of a list
type Node struct {
	Data int
	Next *Node
}

// Append will append a new Node with data d at the end
func Append(head *Node, d int) *Node { // returns the new head
	if head == nil {
		return &Node{Data: d, Next: nil}
	}

	n := head
	for n.Next != nil {
		n = n.Next
	}

	n.Next = &Node{Data: d, Next: nil}

	return head
}

// DeleteNode will delete the Node with data d and manage pointers well
func DeleteNode(head *Node, d int) *Node { // returns the new head
	if head == nil {
		return nil
	}

	if head.Data == d {
		return head.Next
	}

	n := head
	for n.Next != nil {
		if n.Next.Data == d {
			n.Next = n.Next.Next
			return head
		}

		n = n.Next
	}

	return head
}
