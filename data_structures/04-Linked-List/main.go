package main

import (
	"LearnGo/data_structures/03-Stacks-and-Queues/02-Array-Stack/arraystack"
	"LearnGo/data_structures/03-Stacks-and-Queues/stack"
	"LearnGo/data_structures/04-Linked-List/linkedlist"
	"LearnGo/data_structures/04-Linked-List/linkedlistqueue"
	"LearnGo/data_structures/04-Linked-List/linkedliststack"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//test1()

	//test2()

	test3()

	//compareStacks()
}

func compareStacks() {
	opCount := 1000000
	arrayStack := arraystack.NewArrayStack()
	fmt.Printf("ArrayStack, opCount:%d, cost time:%d ms\n", opCount, testStack(arrayStack, opCount))

	linkedListStack := linkedliststack.NewLinkedListStack()
	fmt.Printf("LinkedListStack, opCount:%d, cost time:%d ms\n", opCount, testStack(linkedListStack, opCount))
}

func testStack(stack stack.Stack, opCount int) int64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		stack.Push(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		stack.Pop()
	}
	return time.Now().Sub(startTime).Milliseconds()
}

func test1() {
	list := linkedlist.NewLinkedList()
	for i := 0; i < 5; i++ {
		list.AddFirst(i)
		fmt.Println(list)
	}

	list.Add(2, 666)
	fmt.Println(list)

	list.Remove(1)
	fmt.Println(list)

	list.RemoveFirst()
	fmt.Println(list)

	list.RemoveLast()
	fmt.Println(list)
}

func test2() {
	stack := linkedliststack.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("%v\n", stack.Pop())
		fmt.Println(stack)
	}
}

func test3() {
	queue := linkedlistqueue.NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		queue.Offer(i)
		fmt.Println(queue)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(queue.Poll(), queue)
	}
}
