package main

// max-heap of ints
type MaxHeap []int

func (h MaxHeap) Len() int {
  return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
  return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
  h[i], h[j] = h[j], h[i]
}

// Uses a pointer receiver because it modifies the slice's length & not just the contents
func (h *MaxHeap) Push(x interface{}) {
  *h = append(*h, x.(int))
}

// Uses a pointer receiver because it modifies the slice's length & not just the contents
func (h *MaxHeap) Pop() interface{} {
  old := *h
  n   := len(old)
  x   := old[n-1]
  *h   = old[:n-1]
  return x
}

func (h *MaxHeap) Peek() int {
  heap := *h
  n := len(*h)
  return heap[n-1]
}
