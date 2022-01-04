package main

import (
	"LearnGo/data_structures/02-Arrays/07-Dynamic-Array/array"
	"fmt"
)

func main() {
	arr := array.New(10)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)

	arr.RemoveFirst()
	arr.RemoveFirst()
	fmt.Println(arr)

}
