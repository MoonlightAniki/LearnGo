package main

import (
	"LearnGo/data_structures/02-Arrays/02-Create-Our-Own-Array/array"
	"fmt"
)

func main() {
	arr := array.New(5)
	fmt.Println(arr, arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())
}
