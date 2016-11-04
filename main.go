package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func addNumber(x int, lowers *MaxHeap, highers *MinHeap) {
	if lowers.Len() == 0 {
		heap.Push(lowers, x)
		return
	}

	highestFromLowers := heap.Pop(lowers).(int)
	heap.Push(lowers, highestFromLowers)
	if x < highestFromLowers {
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
		low := heap.Pop(lowers).(int)
		high := heap.Pop(highers).(int)
		heap.Push(lowers, low)
		heap.Push(highers, high)
		return (float64(low) + float64(high)) / 2.0
	}
	if lowers.Len() < highers.Len() {
		high := heap.Pop(highers).(int)
		heap.Push(highers, high)
		return float64(highers.Peek())
	} else {
		low := heap.Pop(lowers).(int)
		heap.Push(lowers, low)
		return float64(lowers.Peek())
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
		fmt.Printf("%.2f\n", median)
	}
}
