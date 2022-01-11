package main

import (
	"LearnGo/algorithm/04-Heap/heap"
	"fmt"
	"math/rand"
)

func main() {
	maxHeap := heap.NewMaxHeap(100)
	for i := 0; i < 100; i++ {
		maxHeap.Insert(rand.Intn(100))
	}
	for !maxHeap.IsEmpty() {
		fmt.Printf("%v ", maxHeap.ExtractMax())
	}
}
