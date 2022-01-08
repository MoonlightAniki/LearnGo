package main

import (
	"LearnGo/data_structures/03-Stacks-and-Queues/02-Array-Stack/arraystack"
	"fmt"
)

func main() {
	//stack := arrayStack.NewArrayStack()
	//fmt.Println(stack == nil)
	//fmt.Println(stack)
	var stack *arraystack.ArrayStack
	fmt.Println(stack, stack == nil)
	stack = arraystack.NewArrayStack()
	fmt.Println(stack, stack == nil)
	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}
	fmt.Println(stack)
	fmt.Println(stack.GetSize(), stack.IsEmpty())

	fmt.Println(stack.Pop(), stack.GetSize(), stack.Peek(), stack.GetSize())
	for i := 0; i < 4; i++ {
		fmt.Println(stack.Pop())
	}
}
