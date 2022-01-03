package main

import (
	solution0303 "LearnGo/leetcode/0303.Range-Sum-Query-Immutable"
	"fmt"
)

func main() {
	numArray := solution0303.Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2) == 1)
	fmt.Println(numArray.SumRange(2, 5) == -1)
	fmt.Println(numArray.SumRange(0, 5) == -3)
}
