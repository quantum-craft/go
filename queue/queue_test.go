package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := MakeQueue()

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
