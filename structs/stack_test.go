package structs

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack[int]()
	if !stack.IsEmpty() {
		t.Error("Stack should be empty")
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("Expected size 3, got %d", stack.Size())
	}
	peek, ok := stack.Peek()
	if !ok || peek != 3 {
		t.Error("Peek failed")
	}
	pop, ok := stack.Pop()
	if !ok || pop != 3 {
		t.Error("Pop failed")
	}
	if stack.Size() != 2 {
		t.Error("Size after pop failed")
	}
	stack.Clear()
	if !stack.IsEmpty() {
		t.Error("Clear failed")
	}
}
