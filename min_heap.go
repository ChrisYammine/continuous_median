package main

// min-heap of ints
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Uses a pointer receiver because it modifies the slice's length & not just the contents
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Uses a pointer receiver because it modifies the slice's length & not just the contents
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Peek() int {
	heap := *h
	n := len(*h)
	return heap[n-1]
}
