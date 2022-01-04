package main

import "fmt"

func main() {
	var arr [10]int
	fmt.Println(arr, len(arr), cap(arr))

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr, len(arr), cap(arr))

	scores := [...]int{100, 99, 66}
	for i := 0; i < len(scores); i++ {
		fmt.Println(scores[i])
	}

	for _, score := range scores {
		fmt.Println(score)
	}

	scores[0] = 101
	fmt.Println(scores)
}
