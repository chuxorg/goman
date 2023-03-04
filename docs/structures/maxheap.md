# MaxHeap

## What is a MaxHeap?

A MaxHeap is a type of binary tree where every parent node is greater than or equal to its children. In other words, the root node of a MaxHeap is always the maximum element in the tree. MaxHeaps are used in many applications where fast access to the maximum element is important, such as implementing priority queues or sorting algorithms.

## How does a MaxHeap work?

A MaxHeap is implemented as a binary tree, where each node has at most two child nodes. The tree is structured so that the value of each parent node is greater than or equal to the value of its children. This means that the maximum element is always at the root node of the tree.

To maintain the MaxHeap property, there are two operations that can be performed on the tree:

- **Insertion:** To insert a new element into the MaxHeap, it is added as a new leaf node at the end of the tree, and then "bubbled up" by repeatedly swapping it with its parent node until the MaxHeap property is satisfied.

- **Extraction:** To extract the maximum element from the MaxHeap, the root node is removed and replaced with the last leaf node in the tree. Then, the new root node is "bubbled down" by repeatedly swapping it with its largest child node until the MaxHeap property is satisfied.

## Use cases for MaxHeap

### Priority Queue

One common use case for a MaxHeap is implementing a priority queue. A priority queue is a data structure that allows elements to be inserted and extracted in order of priority. The element with the highest priority is always at the front of the queue.

Here's an example of using a MaxHeap to implement a priority queue in Go:

```go
// Create a new MaxHeap
h := MaxHeap{}

// Push some elements onto the heap
h.Push(5)
h.Push(7)
h.Push(3)
h.Push(2)
h.Push(8)

// Extract the maximum element (8) and print it
fmt.Println(h.Pop()) // Output: 8

// Extract the next maximum element (7) and print it
fmt.Println(h.Pop()) // Output: 7
```

### Sorting Algorithm

Another use case for a MaxHeap is implementing a sorting algorithm. The heapsort algorithm is a comparison-based sorting algorithm that works by building a MaxHeap from the elements to be sorted, and then repeatedly extracting the maximum element from the heap and placing it at the end of the sorted list.

Here's an example of using a MaxHeap to implement the heapsort algorithm in Python

```python
def heapsort(arr):
    # Build a MaxHeap from the array
    heapify(arr)

    # Extract the maximum element from the heap and place it at the end of the array
    for i in range(len(arr) - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        sift_down(arr, 0, i - 1)

def heapify(arr):
    # Convert the array into a MaxHeap
    n = len(arr)
    for i in range(n // 2 - 1, -1, -1):
        sift_down(arr, i, n - 1)

def sift_down(arr, start, end):
    # Move the element at the given start index down the MaxHeap until the MaxHeap property is satisfied
    root = start
    while root * 2 + 1 <= end:
        child = root * 2 + 1
        if child + 1 <= end and arr[child] < arr[child + 1]:
            child += 1
        if arr[root] < arr[child]:
            arr[root], arr[child] = arr[child], arr[root]
            root = child
        else:
            return
```
In this example, the heapsort() function takes an array as input and sorts it in ascending order using a MaxHeap. The heapify() function converts the input array into a MaxHeap, and the sift_down() function moves an element down the MaxHeap until the MaxHeap property is satisfied.





