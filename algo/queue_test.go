package algo

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()
	if !queue.IsEmpty() {
		t.Error("Queue should be empty")
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.Size() != 3 {
		t.Errorf("Expected size 3, got %d", queue.Size())
	}
	peek, ok := queue.Peek()
	if !ok || peek != 1 {
		t.Error("Peek failed")
	}
	dequeue, ok := queue.Dequeue()
	if !ok || dequeue != 1 {
		t.Error("Dequeue failed")
	}
	if queue.Size() != 2 {
		t.Error("Size after dequeue failed")
	}
	queue.Clear()
	if !queue.IsEmpty() {
		t.Error("Clear failed")
	}
}
