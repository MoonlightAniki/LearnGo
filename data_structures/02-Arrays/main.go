package main

import (
	"LearnGo/data_structures/02-Arrays/array"
	"fmt"
)

func main() {
	arr := array.NewArray(10)
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

	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	arr.RemoveFirst()
	fmt.Println(arr)
	arr.RemoveFirst()
	fmt.Println(arr)

}
