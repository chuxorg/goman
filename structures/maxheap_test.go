package structures

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	// Create a new MaxHeap
	h := MaxHeap{}

	// Push some elements onto the heap
	h.Push(5)
	h.Push(7)
	h.Push(3)
	h.Push(2)
	h.Push(8)

	// Verify that the Peek method returns the maximum element without removing it
	expectedMax := 8
	actualMax := h.Peek()
	if actualMax != expectedMax {
		t.Errorf("Peek: expected %d, but got %d", expectedMax, actualMax)
	}

	// Verify that the Pop method removes and returns the maximum element
	expectedPop := 8
	actualPop := h.Pop()
	if actualPop != expectedPop {
		t.Errorf("Pop: expected %d, but got %d", expectedPop, actualPop)
	}

	// Verify that the Pop method maintains the MaxHeap property
	expectedSlice := []int{7, 5, 3, 2}
	actualSlice := h.slice
	if fmt.Sprintf("%v", actualSlice) != fmt.Sprintf("%v", expectedSlice) {
		t.Errorf("Pop: expected %v, but got %v", expectedSlice, actualSlice)
	}

	// Push some more elements onto the heap
	h.Push(10)
	h.Push(1)

	// Verify that the Pop method maintains the MaxHeap property
	expectedSlice = []int{10, 7, 3, 2, 5, 1}
	actualSlice = h.slice
	if fmt.Sprintf("%v", actualSlice) != fmt.Sprintf("%v", expectedSlice) {
		t.Errorf("Push and Pop: expected %v, but got %v", expectedSlice, actualSlice)
	}
}
