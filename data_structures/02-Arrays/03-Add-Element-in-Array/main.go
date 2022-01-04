package main

import (
	"LearnGo/data_structures/02-Arrays/03-Add-Element-in-Array/array"
	"fmt"
)

func main() {
	arr := array.New(5)
	arr.AddLast(1)
	arr.AddLast(2)
	arr.AddLast(3)
	fmt.Println(arr)
	arr.AddFirst(4)
	//arr.AddFirst(5)
	fmt.Println(arr)
	arr.Add(1, 100)
	fmt.Println(arr)
}
