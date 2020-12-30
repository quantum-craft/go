package queue

import (
	"testing"
)

type node struct {
	val       int
	testOrder int
}

func TestQueue(t *testing.T) {
	queue := NewQueue()

	queue.PushFront(&node{
		val:       100,
		testOrder: 182,
	})

	f, ok := queue.PopFront().(*node)
	if !ok {
		t.Error("TestQueue error !")
	}

	if f.testOrder != 182 || f.val != 100 {
		t.Error("TestQueue error !")
	}

	queue.PushFront(3)
	queue.PushFront(4)

	if queue.PopFront() != 4 {
		t.Error("TestQueue error !")
	}

	if queue.PopFront() != 3 {
		t.Error("TestQueue error !")
	}

	queue.PushBack(3)
	queue.PushBack(4)

	if queue.PopBack() != 4 {
		t.Error("TestQueue error !")
	}

	if queue.PopBack() != 3 {
		t.Error("TestQueue error !")
	}
}
