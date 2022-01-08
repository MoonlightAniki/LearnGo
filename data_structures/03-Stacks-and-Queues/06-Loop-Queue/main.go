package main

import (
	"LearnGo/data_structures/03-Stacks-and-Queues/06-Loop-Queue/loopqueue"
	"fmt"
)

func main() {
	queue := loopqueue.NewLoopQueue()
	fmt.Println(queue)

	for i := 0; i < 30; i++ {
		queue.Offer(i)
		fmt.Println(queue)
	}

	for i := 0; i < 25; i++ {
		fmt.Println(queue.Poll())
		fmt.Println(queue)
	}
}
