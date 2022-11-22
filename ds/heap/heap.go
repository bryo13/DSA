package main

import (
	"container/heap"
	"fmt"
)

// a min heap of integers
type HeapOfInts []int

func (hi HeapOfInts) Len() int { return len(hi) }

func (hi HeapOfInts) Less(i, j int) bool { return hi[i] < hi[j] }

func (hi HeapOfInts) Swap(i, j int) { hi[i], hi[j] = hi[j], hi[i] }

func (hi *HeapOfInts) Push(k any) {
	*hi = append(*hi, k.(int))
}

func (hi *HeapOfInts) Pop() any {
	old := *hi
	n := len(old)
	x := old[n-1]
	*hi = old[0 : n-1]
	return x
}

func ExampleHeap() {
	h := &HeapOfInts{54, 34, 23, 1, 0}
	heap.Init(h)
	heap.Push(h, 3)

	fmt.Println(h)
}

func main() {
	ExampleHeap()
}
