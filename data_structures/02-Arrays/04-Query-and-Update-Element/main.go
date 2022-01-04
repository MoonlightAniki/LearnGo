package main

import (
	"LearnGo/data_structures/02-Arrays/04-Query-and-Update-Element/array"
	"fmt"
)

func main() {
	arr := array.New(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)
}
