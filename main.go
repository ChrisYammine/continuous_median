package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Peek() int {
	heap := *h
	return heap[0]
}

// Create our MaxHeap by embedding MinHeap
type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

func addNumber(x int, lowers *MaxHeap, highers *MinHeap) {
	if lowers.Len() == 0 {
		heap.Push(lowers, x)
		return
	}

	if x < lowers.Peek() {
		heap.Push(lowers, x)
	} else {
		heap.Push(highers, x)
	}
}

func rebalance(lowers *MaxHeap, highers *MinHeap) {
	if lowers.Len() < highers.Len() {
		if highers.Len()-lowers.Len() >= 2 {
			heap.Push(lowers, heap.Pop(highers))
		}
	} else {
		if lowers.Len()-highers.Len() >= 2 {
			heap.Push(highers, heap.Pop(lowers))
		}
	}
}

func findMedian(lowers *MaxHeap, highers *MinHeap) float64 {
	if lowers.Len() == highers.Len() {
    low := lowers.Peek()
    high := highers.Peek()
		return (float64(low) + float64(high)) / 2.0
	}
	if lowers.Len() < highers.Len() {
    high := highers.Peek()
		return float64(high)
	} else {
    low := lowers.Peek()
		return float64(low)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, n)

	lowers := &MaxHeap{}
	highers := &MinHeap{}
	heap.Init(lowers)
	heap.Init(highers)

	// Collecting our input
	for i, _ := range arr {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		arr[i] = x
	}

	// Printing our medians
	for _, x := range arr {
		addNumber(x, lowers, highers)
		rebalance(lowers, highers)
		median := findMedian(lowers, highers)
		fmt.Printf("%.1f\n", median)
	}
}
