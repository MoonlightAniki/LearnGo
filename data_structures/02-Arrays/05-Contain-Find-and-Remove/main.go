package main

import (
	"LearnGo/data_structures/02-Arrays/05-Contain-Find-and-Remove/array"
	"fmt"
)

func main() {
	a := array.New(20)
	for i := 0; i < 10; i++ {
		a.AddLast(i)
	}
	for i := 0; i < 10; i++ {
		a.AddFirst(i)
	}
	fmt.Println(a)

	fmt.Println(a.Contains(5), a.Contains(100))

	fmt.Println(a.Find(8), a.FindAll(8), a.Find(10), a.FindAll(10))

	fmt.Println(a.RemoveElement(1))
	fmt.Println(a)

	fmt.Println(a.RemoveAllElements(0))
	fmt.Println(a)
}
