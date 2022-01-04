package main

import (
	"LearnGo/basics/queue"
	"fmt"
)

func main() {
	q := queue.Queue{}
	fmt.Println(q.String(), q.IsEmpty())

	q = queue.Queue{1}
	fmt.Println(q.String(), q.IsEmpty())

	q.Push(2)
	q.Push(3)
	fmt.Println(q.String(), q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
