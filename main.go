package main

import (
  "os"
  "bufio"
  "fmt"
  "container/heap"
  "strconv"
)

func addNumber(x int, lowers *MaxHeap, highers *MinHeap) {
  if lowers.Len() == 0 || x < lowers.Peek() {
    heap.Push(lowers, x)
  } else {
    heap.Push(highers, x)
  }
}

func rebalance(lowers *MaxHeap, highers *MinHeap) {
  if lowers.Len() < highers.Len() {
    if highers.Len() - lowers.Len() >= 2 {
      heap.Push(lowers, heap.Pop(highers))
    }
  } else {
    if lowers.Len() - highers.Len() >= 2 {
      heap.Push(highers, heap.Pop(lowers))
    }
  }
}

func findMedian(lowers *MaxHeap, highers *MinHeap) float64 {
  if lowers.Len() == highers.Len() {
    return (float64(lowers.Peek()) + float64(highers.Peek())) / 2.0
  }
  if lowers.Len() < highers.Len() {
    return float64(highers.Peek())
  } else {
    return float64(lowers.Peek())
  }
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())

  arr  := make([]int, n)

  lowers  := &MaxHeap{}
  highers := &MinHeap{}
  heap.Init(lowers)
  heap.Init(highers)

  // Collecting our input
  for i, _ := range(arr) {
    scanner.Scan()
    x, _ := strconv.Atoi(scanner.Text())
    arr[i] = x
  }

  // Printing our medians
  for _, x := range(arr) {
    addNumber(x, lowers, highers)
    rebalance(lowers, highers)
    median := findMedian(lowers, highers)
    fmt.Printf("%.2f\n", median)
  }
}
