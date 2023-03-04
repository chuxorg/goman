package structures

// A MaxHeap is a binary heap data structure that maintains the MaxHeap property:
// each parent node is greater than or equal to its children nodes.
// The MaxHeap supports insertion, removal of the maximum element, and peeking at the maximum element.
// This is the type used to represent a MaxHeap.
type MaxHeap struct {
	slice []int
}

// instantiates and initializes a new MaxHeap
func NewMaxHeap() *MaxHeap {
	heap := &MaxHeap{}
	return heap
}

// Push adds a new element to the MaxHeap and adjusts the heap to maintain the MaxHeap property.
// It adds the new element to the end of the heap, and then repeatedly swaps it with its parent node
// until it is in the correct position in the heap.
func (h *MaxHeap) Push(val int) {
	// Append the new element to the end of the heap
	h.slice = append(h.slice, val)

	// Find the index of the new element
	index := len(h.slice) - 1

	// Repeatedly swap the new element with its parent node until it is in the correct position
	// in the heap, until it is either the root of the heap and its value is greater than its parent's value.
	for index > 0 && h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Pop removes and returns the maximum element from the MaxHeap and adjusts the heap to maintain the MaxHeap property.
// It removes the root node of the heap (which contains the maximum element), replaces it with the last element in the heap,
// and then repeatedly swaps the new root node with its larger child nodes until it is in the correct position in the heap.
func (h *MaxHeap) Pop() int {
	// Get the index of the last element in the heap
	lastIndex := len(h.slice) - 1

	// Swap the root node (which contains the maximum element) with the last element in the heap
	rootVal := h.slice[0]
	h.slice[0] = h.slice[lastIndex]
	h.slice = h.slice[:lastIndex]

	// Adjust the heap by repeatedly swapping the new root node with its larger child nodes until it is in the correct position
	// in the heap, as long as it has one or more child nodes that are greater than it.
	index := 0
	for childIndex := leftChild(index); childIndex < lastIndex; childIndex = leftChild(index) {
		if rightChild := childIndex + 1; rightChild < lastIndex && h.slice[childIndex] < h.slice[rightChild] {
			childIndex = rightChild
		}
		if h.slice[index] >= h.slice[childIndex] {
			break
		}
		h.swap(index, childIndex)
		index = childIndex
	}

	return rootVal
}

// Peek returns the maximum element in the MaxHeap without removing it.
// It simply returns the root node of the heap, which contains the maximum element.
func (h MaxHeap) Peek() int {
	return h.slice[0]
}

// Returns the length of the Heap
func (h MaxHeap) Len() int {
	return len(h.slice)
}

// This method simply swaps the values at indices i and j in the MaxHeap slice by using Go's multiple assignment syntax (heap[i], heap[j] = heap[j], heap[i]).
// The purpose of the Swap method is to allow the MaxHeap to be manipulated while maintaining the MaxHeap property.
// When a new element is added to the MaxHeap, the Push method appends the element to the end of the slice and
// then percolates it up the tree by repeatedly swapping it with its parent until the MaxHeap property is satisfied.
// Similarly, when the largest element is removed from the MaxHeap, the Pop method replaces it with the last element
// in the slice and then percolates it down the tree by repeatedly swapping it with its larger child until the MaxHeap property is satisfied.
// The Swap method is used in both cases to swap the values at the appropriate indices in the MaxHeap slice.
func (h *MaxHeap) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

// Returns the index of a parent in relation to the target index.
// parent returns the index of the parent node of the element at the given index.
func parent(index int) int {
	return (index - 1) / 2
}

// leftChild returns the index of the left child node of the element at the given index.
func leftChild(index int) int {
	return 2*index + 1
}

// // rightChild returns the index of the right child node of the element at the given index.
// func rightChild(index int) int {
// 	return 2*index + 2
// }
