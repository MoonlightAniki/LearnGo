package main

import (
	"LearnGo/data_structures/03-Stacks-and-Queues/05-Array-Queue/arrayqueue"
	"LearnGo/data_structures/03-Stacks-and-Queues/06-Loop-Queue/loopqueue"
	"LearnGo/data_structures/03-Stacks-and-Queues/queue"
	"LearnGo/data_structures/04-Linked-List/linkedlistqueue"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	opCount := 100000
	aq := arrayqueue.NewArrayQueue()
	fmt.Printf("ArrayQueue: opCount:%d, cost time:%d ms.\n", opCount, testQueue(aq, opCount))

	lq := loopqueue.NewLoopQueue()
	fmt.Printf("LoopQueue: opCount:%d, cost time:%d ms.\n", opCount, testQueue(lq, opCount))

	llq := linkedlistqueue.NewLinkedListQueue()
	fmt.Printf("LinkedListQueue: opCount:%d, cost time:%d ms.\n", opCount, testQueue(llq, opCount))
}

func testQueue(queue queue.Queue, opCount int) int64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		queue.Offer(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		queue.Poll()
	}
	return time.Now().Sub(startTime).Milliseconds()
}
