package main

import (
	"LearnGo/data_structures/09-Segment-Tree/segmenttree"
	"fmt"
)

func main() {
	tree := segmenttree.NewSegmentTree([]int{1, 3, 5, 2, 4, 6}, func(a int, b int) int { return a + b })
	fmt.Println(tree)
	fmt.Println(tree.Query(1, 2))
	tree.Update(0, 100)
	tree.Update(1, 300)
	fmt.Println(tree.Query(0, 1))
	fmt.Println(tree.Query(1, 2))
	fmt.Println(tree.Query(1, 4))
}
