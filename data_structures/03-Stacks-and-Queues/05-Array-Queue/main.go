package main

import (
	"LearnGo/data_structures/03-Stacks-and-Queues/05-Array-Queue/arrayqueue"
	"fmt"
)

func main() {
	queue := &arrayqueue.ArrayQueue{}
	fmt.Println(queue)
	queue.Offer(1)
	queue.Offer(2)
	queue.Offer(3)
	queue.Offer(4)
	fmt.Println(queue)
	fmt.Println(queue.Poll())
	fmt.Println(queue)
}
