---
id: ljglm
title: Feature/structure
file_version: 1.1.2
app_version: 1.3.8
---


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 cmd/main.go
```go
1      // Package main implements the main function for goman
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 docs/structures/maxheap.md
```markdown
1      # MaxHeap
2      
3      ## What is a MaxHeap?
4      
5      A MaxHeap is a type of binary tree where every parent node is greater than or equal to its children. In other words, the root node of a MaxHeap is always the maximum element in the tree. MaxHeaps are used in many applications where fast access to the maximum element is important, such as implementing priority queues or sorting algorithms.
6      
7      ## How does a MaxHeap work?
8      
9      A MaxHeap is implemented as a binary tree, where each node has at most two child nodes. The tree is structured so that the value of each parent node is greater than or equal to the value of its children. This means that the maximum element is always at the root node of the tree.
10     
11     To maintain the MaxHeap property, there are two operations that can be performed on the tree:
12     
13     - **Insertion:** To insert a new element into the MaxHeap, it is added as a new leaf node at the end of the tree, and then "bubbled up" by repeatedly swapping it with its parent node until the MaxHeap property is satisfied.
14     
15     - **Extraction:** To extract the maximum element from the MaxHeap, the root node is removed and replaced with the last leaf node in the tree. Then, the new root node is "bubbled down" by repeatedly swapping it with its largest child node until the MaxHeap property is satisfied.
16     
17     ## Use cases for MaxHeap
18     
19     ### Priority Queue
20     
21     One common use case for a MaxHeap is implementing a priority queue. A priority queue is a data structure that allows elements to be inserted and extracted in order of priority. The element with the highest priority is always at the front of the queue.
22     
23     Here's an example of using a MaxHeap to implement a priority queue in Go:
24     
25     ```go
26     // Create a new MaxHeap
27     h := MaxHeap{}
28     
29     // Push some elements onto the heap
30     h.Push(5)
31     h.Push(7)
32     h.Push(3)
33     h.Push(2)
34     h.Push(8)
35     
36     // Extract the maximum element (8) and print it
37     fmt.Println(h.Pop()) // Output: 8
38     
39     // Extract the next maximum element (7) and print it
40     fmt.Println(h.Pop()) // Output: 7
41     ```
42     
43     ### Sorting Algorithm
44     
45     Another use case for a MaxHeap is implementing a sorting algorithm. The heapsort algorithm is a comparison-based sorting algorithm that works by building a MaxHeap from the elements to be sorted, and then repeatedly extracting the maximum element from the heap and placing it at the end of the sorted list.
46     
47     Here's an example of using a MaxHeap to implement the heapsort algorithm in Python
48     
49     ```python
50     def heapsort(arr):
51         # Build a MaxHeap from the array
52         heapify(arr)
53     
54         # Extract the maximum element from the heap and place it at the end of the array
55         for i in range(len(arr) - 1, 0, -1):
56             arr[i], arr[0] = arr[0], arr[i]
57             sift_down(arr, 0, i - 1)
58     
59     def heapify(arr):
60         # Convert the array into a MaxHeap
61         n = len(arr)
62         for i in range(n // 2 - 1, -1, -1):
63             sift_down(arr, i, n - 1)
64     
65     def sift_down(arr, start, end):
66         # Move the element at the given start index down the MaxHeap until the MaxHeap property is satisfied
67         root = start
68         while root * 2 + 1 <= end:
69             child = root * 2 + 1
70             if child + 1 <= end and arr[child] < arr[child + 1]:
71                 child += 1
72             if arr[root] < arr[child]:
73                 arr[root], arr[child] = arr[child], arr[root]
74                 root = child
75             else:
76                 return
77     ```
78     In this example, the heapsort() function takes an array as input and sorts it in ascending order using a MaxHeap. The heapify() function converts the input array into a MaxHeap, and the sift_down() function moves an element down the MaxHeap until the MaxHeap property is satisfied.
79     
80     
81     
82     
83     
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/maxheap.go
```go
12     func NewMaxHeap() *MaxHeap {
13     	heap := &MaxHeap{}
```

<br/>

Here is a simple builder pattern to construct a new `MaxHeap`<swm-token data-swm-token=":docs/structures/maxheap.md:3:8:8:`## What is a MaxHeap?`"/>
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/maxheap.go
```go
12     func NewMaxHeap() *MaxHeap {
13     	heap := &MaxHeap{}
14     	return heap
15     }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/README.md
```markdown
1      # Introduction to Data Structures
2      
3      Data structures are a fundamental part of computer science and programming. They are used to store and organize data in a way that makes it easy to access and manipulate.
4      
5      There are many different types of data structures, each with its own strengths and weaknesses. Some of the most common data structures include arrays, linked lists, stacks, queues, trees, and graphs.
6      
7      ## Table of Contents
8      
9      - [Introduction to Data Structures](#introduction-to-data-structures)
10       - [Table of Contents](#table-of-contents)
11       - [Arrays](#arrays)
12       - [Linked Lists](#linked-lists)
13       - [Stacks](#stacks)
14       - [Queues](#queues)
15       - [Trees](#trees)
16         - [Binary Trees](#binary-trees)
17         - [Binary Search Trees](#binary-search-trees)
18         - [Heap](#heap)
19           - [MaxHeap](#maxheap)
20     
21     ## Arrays
22     
23     An array is a collection of elements of the same type, stored in contiguous memory locations. Arrays are a simple and efficient way to store and access data, but they have a fixed size and cannot be easily resized or modified.
24     
25     ## Linked Lists
26     
27     A linked list is a collection of elements, each of which contains a reference to the next element in the list. Linked lists can be used to store and access data in a flexible and dynamic way, but they can be less efficient than arrays for certain operations.
28     
29     ## Stacks
30     
31     A stack is a Last-In-First-Out (LIFO) data structure, where elements are added and removed from the top of the stack. Stacks can be used to implement recursive algorithms, evaluate expressions, and keep track of function calls.
32     
33     ## Queues
34     
35     A queue is a First-In-First-Out (FIFO) data structure, where elements are added to the back of the queue and removed from the front. Queues can be used to implement breadth-first search algorithms, simulate waiting lines, and handle requests in a web server.
36     
37     ## Trees
38     
39     A tree is a collection of nodes, each of which contains a value and references to its child nodes. Trees can be used to represent hierarchical relationships, and they are the basis for many common algorithms and data structures.
40     
41     ### Binary Trees
42     
43     A binary tree is a tree in which each node has at most two child nodes. Binary trees can be used to represent sorting algorithms, decision trees, and expression trees.
44     
45     ### Binary Search Trees
46     
47     A binary search tree is a binary tree in which the left child node is always less than the parent node, and the right child node is always greater than the parent node. Binary search trees can be used to implement search algorithms, and they have efficient average-case performance for many operations.
48     
49     ### Heap
50     
51     A heap is a binary tree in which each parent node is greater than or equal to its child nodes (in a MaxHeap) or less than or equal to its child nodes (in a MinHeap). Heaps can be used to implement priority queues, sorting algorithms, and graph algorithms.
52     
53     #### MaxHeap
54     
55     A MaxHeap is a heap in which each parent node is greater than or equal to its child nodes. The maximum element is always at the root node of the heap. A MaxHeap can be used to implement a priority queue, where the element with the highest priority is always at the front of the queue.
56     
57     MaxHeap is a useful data structure for solving many programming problems that require efficient management of priorities. With its simple and efficient operations, it can be a powerful tool in any programmer's toolkit.
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/heap.go
```go
1      package structures
2      
3      // A heap is a data structure that represents a binary tree where each node is
4      // either greater than or equal to (in a max heap) or less than or equal to (in a min heap) its children.
5      // A max heap is a type of heap where each node is greater than or equal to its children.
6      // The main difference between a heap and a max heap is that a heap can be either a max heap or a min heap,
7      // whereas a max heap is specifically a type of heap where the largest element is always stored at the root.
8      // In other words, a max heap is a specialized form of a heap where the largest value is always the root node,
9      // and its child nodes are smaller than it.
10     // In a max heap, the largest element can be retrieved in constant time (O(1)), because it is always stored at the root.
11     // In contrast, retrieving the largest element from a heap in general requires traversing the tree, which takes
12     // logarithmic time (O(log n)), where n is the number of nodes in the heap.
13     // In summary, while both heaps and max heaps are binary trees with a specific structure, a max heap is a specific type of heap
14     // where the largest element is always stored at the root.
15     
16     // Heap is a data structure that represents a binary max heap.
17     type Heap struct {
18     	array []int // slice that holds the values in the heap
19     }
20     
21     // Insert adds a value to the heap.
22     func (h *Heap) Insert(value int) {
23     	// Add the value to the end of the slice
24     	h.array = append(h.array, value)
25     
26     	// Bubble up the value to its correct position in the heap
27     	h.bubbleUp(len(h.array) - 1)
28     }
29     
30     // bubbleUp moves a value up the heap until it is in the correct position.
31     func (h *Heap) bubbleUp(index int) {
32     	// Calculate the index of the parent node
33     	parent := (index - 1) / 2
34     
35     	// If the value at the current index is greater than its parent, swap them
36     	if parent >= 0 && h.array[index] > h.array[parent] {
37     		h.array[index], h.array[parent] = h.array[parent], h.array[index]
38     
39     		// Recursively bubble up the parent node until the value is in the correct position
40     		h.bubbleUp(parent)
41     	}
42     }
43     
44     // Remove removes and returns the maximum value from the heap.
45     func (h *Heap) Remove() int {
46     	// If the heap is empty, return -1
47     	if len(h.array) == 0 {
48     		return -1
49     	}
50     
51     	// Save the root value and replace it with the last value in the heap
52     	root := h.array[0]
53     	h.array[0] = h.array[len(h.array)-1]
54     	h.array = h.array[:len(h.array)-1]
55     
56     	// Bubble down the new root value to its correct position in the heap
57     	h.bubbleDown(0)
58     
59     	// Return the saved root value
60     	return root
61     }
62     
63     // bubbleDown moves a value down the heap until it is in the correct position.
64     func (h *Heap) bubbleDown(index int) {
65     	// Calculate the indices of the left and right child nodes
66     	left := index*2 + 1
67     	right := index*2 + 2
68     
69     	// Find the index of the largest value among the current node and its children
70     	largest := index
71     	if left < len(h.array) && h.array[left] > h.array[largest] {
72     		largest = left
73     	}
74     	if right < len(h.array) && h.array[right] > h.array[largest] {
75     		largest = right
76     	}
77     
78     	// If the largest value is not the current node, swap them and recursively bubble down the new node
79     	if largest != index {
80     		h.array[index], h.array[largest] = h.array[largest], h.array[index]
81     		h.bubbleDown(largest)
82     	}
83     }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/heap_test.go
```go
1      package structures
2      
3      import (
4      	"testing"
5      )
6      
7      func TestHeap(t *testing.T) {
8      	// Create a new heap
9      	h := Heap{}
10     
11     	// Add some values to the heap
12     	h.Insert(3)
13     	h.Insert(2)
14     	h.Insert(1)
15     	h.Insert(4)
16     
17     	// Test the Remove method
18     	if h.Remove() != 4 {
19     		t.Errorf("Expected %d, got %d", 4, h.Remove())
20     	}
21     
22     	// Test the bubbleUp method
23     	h.Insert(5)
24     	if h.array[0] != 5 {
25     		t.Errorf("Expected %d, got %d", 5, h.array[0])
26     	}
27     
28     	// Test the bubbleDown method
29     	h.Remove()
30     	if h.array[0] != 3 {
31     		t.Errorf("Expected %d, got %d", 3, h.array[0])
32     	}
33     
34     	// Test the Remove method on an empty heap
35     	h.array = []int{}
36     	if h.Remove() != -1 {
37     		t.Errorf("Expected %d, got %d", -1, h.Remove())
38     	}
39     }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/maxheap.go
```go
1      package structures
2      
3      // A MaxHeap is a binary heap data structure that maintains the MaxHeap property:
4      // each parent node is greater than or equal to its children nodes.
5      // The MaxHeap supports insertion, removal of the maximum element, and peeking at the maximum element.
6      // This is the type used to represent a MaxHeap.
7      type MaxHeap struct {
8      	slice []int
9      }
10     
11     // instantiates and initializes a new MaxHeap
12     func NewMaxHeap() *MaxHeap {
13     	heap := &MaxHeap{}
14     	return heap
15     }
16     
17     // Push adds a new element to the MaxHeap and adjusts the heap to maintain the MaxHeap property.
18     // It adds the new element to the end of the heap, and then repeatedly swaps it with its parent node
19     // until it is in the correct position in the heap.
20     func (h *MaxHeap) Push(val int) {
21     	// Append the new element to the end of the heap
22     	h.slice = append(h.slice, val)
23     
24     	// Find the index of the new element
25     	index := len(h.slice) - 1
26     
27     	// Repeatedly swap the new element with its parent node until it is in the correct position
28     	// in the heap, until it is either the root of the heap and its value is greater than its parent's value.
29     	for index > 0 && h.slice[parent(index)] < h.slice[index] {
30     		h.swap(parent(index), index)
31     		index = parent(index)
32     	}
33     }
34     
35     // Pop removes and returns the maximum element from the MaxHeap and adjusts the heap to maintain the MaxHeap property.
36     // It removes the root node of the heap (which contains the maximum element), replaces it with the last element in the heap,
37     // and then repeatedly swaps the new root node with its larger child nodes until it is in the correct position in the heap.
38     func (h *MaxHeap) Pop() int {
39     	// Get the index of the last element in the heap
40     	lastIndex := len(h.slice) - 1
41     
42     	// Swap the root node (which contains the maximum element) with the last element in the heap
43     	rootVal := h.slice[0]
44     	h.slice[0] = h.slice[lastIndex]
45     	h.slice = h.slice[:lastIndex]
46     
47     	// Adjust the heap by repeatedly swapping the new root node with its larger child nodes until it is in the correct position
48     	// in the heap, as long as it has one or more child nodes that are greater than it.
49     	index := 0
50     	for childIndex := leftChild(index); childIndex < lastIndex; childIndex = leftChild(index) {
51     		if rightChild := childIndex + 1; rightChild < lastIndex && h.slice[childIndex] < h.slice[rightChild] {
52     			childIndex = rightChild
53     		}
54     		if h.slice[index] >= h.slice[childIndex] {
55     			break
56     		}
57     		h.swap(index, childIndex)
58     		index = childIndex
59     	}
60     
61     	return rootVal
62     }
63     
64     // Peek returns the maximum element in the MaxHeap without removing it.
65     // It simply returns the root node of the heap, which contains the maximum element.
66     func (h MaxHeap) Peek() int {
67     	return h.slice[0]
68     }
69     
70     // Returns the length of the Heap
71     func (h MaxHeap) Len() int {
72     	return len(h.slice)
73     }
74     
75     // This method simply swaps the values at indices i and j in the MaxHeap slice by using Go's multiple assignment syntax (heap[i], heap[j] = heap[j], heap[i]).
76     // The purpose of the Swap method is to allow the MaxHeap to be manipulated while maintaining the MaxHeap property.
77     // When a new element is added to the MaxHeap, the Push method appends the element to the end of the slice and
78     // then percolates it up the tree by repeatedly swapping it with its parent until the MaxHeap property is satisfied.
79     // Similarly, when the largest element is removed from the MaxHeap, the Pop method replaces it with the last element
80     // in the slice and then percolates it down the tree by repeatedly swapping it with its larger child until the MaxHeap property is satisfied.
81     // The Swap method is used in both cases to swap the values at the appropriate indices in the MaxHeap slice.
82     func (h *MaxHeap) swap(i, j int) {
83     	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
84     }
85     
86     // Returns the index of a parent in relation to the target index.
87     // parent returns the index of the parent node of the element at the given index.
88     func parent(index int) int {
89     	return (index - 1) / 2
90     }
91     
92     // leftChild returns the index of the left child node of the element at the given index.
93     func leftChild(index int) int {
94     	return 2*index + 1
95     }
96     
97     // // rightChild returns the index of the right child node of the element at the given index.
98     // func rightChild(index int) int {
99     // 	return 2*index + 2
100    // }
```

<br/>


<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 structures/maxheap_test.go
```go
1      package structures
2      
3      import (
4      	"fmt"
5      	"testing"
6      )
7      
8      func TestMaxHeap(t *testing.T) {
9      	// Create a new MaxHeap
10     	h := MaxHeap{}
11     
12     	// Push some elements onto the heap
13     	h.Push(5)
14     	h.Push(7)
15     	h.Push(3)
16     	h.Push(2)
17     	h.Push(8)
18     
19     	// Verify that the Peek method returns the maximum element without removing it
20     	expectedMax := 8
21     	actualMax := h.Peek()
22     	if actualMax != expectedMax {
23     		t.Errorf("Peek: expected %d, but got %d", expectedMax, actualMax)
24     	}
25     
26     	// Verify that the Pop method removes and returns the maximum element
27     	expectedPop := 8
28     	actualPop := h.Pop()
29     	if actualPop != expectedPop {
30     		t.Errorf("Pop: expected %d, but got %d", expectedPop, actualPop)
31     	}
32     
33     	// Verify that the Pop method maintains the MaxHeap property
34     	expectedSlice := []int{7, 5, 3, 2}
35     	actualSlice := h.slice
36     	if fmt.Sprintf("%v", actualSlice) != fmt.Sprintf("%v", expectedSlice) {
37     		t.Errorf("Pop: expected %v, but got %v", expectedSlice, actualSlice)
38     	}
39     
40     	// Push some more elements onto the heap
41     	h.Push(10)
42     	h.Push(1)
43     
44     	// Verify that the Pop method maintains the MaxHeap property
45     	expectedSlice = []int{10, 7, 3, 2, 5, 1}
46     	actualSlice = h.slice
47     	if fmt.Sprintf("%v", actualSlice) != fmt.Sprintf("%v", expectedSlice) {
48     		t.Errorf("Push and Pop: expected %v, but got %v", expectedSlice, actualSlice)
49     	}
50     }
```

<br/>

This file was generated by Swimm. [Click here to view it in the app](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBZ29tYW4lM0ElM0Fjc2FpbGVy/docs/ljglm).
