package structures

import (
	"testing"
)

func TestHeap(t *testing.T) {
	// Create a new heap
	h := Heap{}

	// Add some values to the heap
	h.Insert(3)
	h.Insert(2)
	h.Insert(1)
	h.Insert(4)

	// Test the Remove method
	if h.Remove() != 4 {
		t.Errorf("Expected %d, got %d", 4, h.Remove())
	}

	// Test the bubbleUp method
	h.Insert(5)
	if h.array[0] != 5 {
		t.Errorf("Expected %d, got %d", 5, h.array[0])
	}

	// Test the bubbleDown method
	h.Remove()
	if h.array[0] != 3 {
		t.Errorf("Expected %d, got %d", 3, h.array[0])
	}

	// Test the Remove method on an empty heap
	h.array = []int{}
	if h.Remove() != -1 {
		t.Errorf("Expected %d, got %d", -1, h.Remove())
	}
}
