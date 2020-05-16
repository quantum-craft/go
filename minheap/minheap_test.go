package minheap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := make([]Node, 0, 0)
	lastEmpty := 0
	minheap := MinHeap{
		heap:      &heap,
		lastEmpty: &lastEmpty,
	}

	Insert(minheap, Node{Key: nil, Value: 88})
	Insert(minheap, Node{Key: nil, Value: 32})
	Insert(minheap, Node{Key: nil, Value: 50})
	Insert(minheap, Node{Key: nil, Value: 90})
	Insert(minheap, Node{Key: nil, Value: 23})
	Insert(minheap, Node{Key: nil, Value: 74})
	Insert(minheap, Node{Key: nil, Value: 2})
}
