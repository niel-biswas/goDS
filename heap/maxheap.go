package maxheap

import "fmt"

type MaxHeap struct {
	array []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{array: make([]int, 0)}
}

func (h *MaxHeap) Push(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Pop() int {
	if len(h.array) == 0 {
		fmt.Println("Cannot pop element, reached empty array")
		return -1
	}
	result := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.maxHeapifyDown(h.array[0])
	return result
}

func (h *MaxHeap) Print() {
	fmt.Println("heap:", h.array)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftChildIndex, rightChildIndex := leftChild(index), rightChild(index)
	childIndexToCompare := 0
	for leftChildIndex <= lastIndex {
		if leftChildIndex == lastIndex {
			childIndexToCompare = leftChildIndex
		} else if leftChildIndex > lastIndex {
			childIndexToCompare = leftChildIndex
		} else {
			childIndexToCompare = rightChildIndex
		}
		if h.array[index] < h.array[childIndexToCompare] {
			h.swap(index, childIndexToCompare)
			index = childIndexToCompare
			leftChildIndex, rightChildIndex = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

// get the parent given the current index
func parent(currentIndex int) int {
	return (currentIndex - 1) / 2
}

// get the left child given the parent index
func leftChild(parentIndex int) int {
	return (parentIndex * 2) + 1
}

// get the right child given the parent index
func rightChild(parentIndex int) int {
	return (parentIndex * 2) + 2
}

// swap array elements based on the given indices
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *MaxHeap) Extract() int {
	return 0
}

/*func main() {
	myHeap := NewMaxHeap()
	myHeap.Push(10)
	myHeap.Print()
	myHeap.Push(20)
	myHeap.Print()
	myHeap.Push(30)
	myHeap.Print()
	fmt.Println("Pop items from the heap begins...")
	maxItem := myHeap.Pop()
	fmt.Println("Pop maxItem:", maxItem)
	myHeap.Print()
	maxItem = myHeap.Pop()
	fmt.Println("Pop maxItem:", maxItem)
	myHeap.Print()
}*/

// Part of Go standard library, heap is already implemented in container/heap interface. You need to define your own struct for heap
// data structure and implement the methods required. The benefit of using such approach is that Go standard library takes care of
// heapify process (Up & Down) reducing the complexity. This heap implementation is just to understand the concept of heap in general.
