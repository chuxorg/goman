package structures

// A heap is a data structure that represents a binary tree where each node is
// either greater than or equal to (in a max heap) or less than or equal to (in a min heap) its children.
// A max heap is a type of heap where each node is greater than or equal to its children.
// The main difference between a heap and a max heap is that a heap can be either a max heap or a min heap,
// whereas a max heap is specifically a type of heap where the largest element is always stored at the root.
// In other words, a max heap is a specialized form of a heap where the largest value is always the root node,
// and its child nodes are smaller than it.
// In a max heap, the largest element can be retrieved in constant time (O(1)), because it is always stored at the root.
// In contrast, retrieving the largest element from a heap in general requires traversing the tree, which takes
// logarithmic time (O(log n)), where n is the number of nodes in the heap.
// In summary, while both heaps and max heaps are binary trees with a specific structure, a max heap is a specific type of heap
// where the largest element is always stored at the root.

// Heap is a data structure that represents a binary max heap.
type Heap struct {
	array []int // slice that holds the values in the heap
}

// Insert adds a value to the heap.
func (h *Heap) Insert(value int) {
	// Add the value to the end of the slice
	h.array = append(h.array, value)

	// Bubble up the value to its correct position in the heap
	h.bubbleUp(len(h.array) - 1)
}

// bubbleUp moves a value up the heap until it is in the correct position.
func (h *Heap) bubbleUp(index int) {
	// Calculate the index of the parent node
	parent := (index - 1) / 2

	// If the value at the current index is greater than its parent, swap them
	if parent >= 0 && h.array[index] > h.array[parent] {
		h.array[index], h.array[parent] = h.array[parent], h.array[index]

		// Recursively bubble up the parent node until the value is in the correct position
		h.bubbleUp(parent)
	}
}

// Remove removes and returns the maximum value from the heap.
func (h *Heap) Remove() int {
	// If the heap is empty, return -1
	if len(h.array) == 0 {
		return -1
	}

	// Save the root value and replace it with the last value in the heap
	root := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]

	// Bubble down the new root value to its correct position in the heap
	h.bubbleDown(0)

	// Return the saved root value
	return root
}

// bubbleDown moves a value down the heap until it is in the correct position.
func (h *Heap) bubbleDown(index int) {
	// Calculate the indices of the left and right child nodes
	left := index*2 + 1
	right := index*2 + 2

	// Find the index of the largest value among the current node and its children
	largest := index
	if left < len(h.array) && h.array[left] > h.array[largest] {
		largest = left
	}
	if right < len(h.array) && h.array[right] > h.array[largest] {
		largest = right
	}

	// If the largest value is not the current node, swap them and recursively bubble down the new node
	if largest != index {
		h.array[index], h.array[largest] = h.array[largest], h.array[index]
		h.bubbleDown(largest)
	}
}
