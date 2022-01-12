package main

import (
	"LearnGo/data_structures/06-Binary-Search-Tree/bst"
	"fmt"
)

func main() {
	/*
	          5
	       /    \
	      3      7
	    /  \    /  \
	   2    4  6    8

	*/

	tree := bst.NewBST()
	tree.Add(5)
	tree.Add(3)
	tree.Add(7)
	tree.Add(2)
	tree.Add(4)
	tree.Add(6)
	tree.Add(8)

	fmt.Println("PreOrder:")
	fmt.Println(tree.PreOrder())
	fmt.Println("InOrder:")
	fmt.Println(tree.InOrder())
	fmt.Println("PostOrder:")
	fmt.Println(tree.PostOrder())
	fmt.Println("LevelOrder:")
	fmt.Println(tree.LevelOrder())
	fmt.Println("PreOrderNR:")
	fmt.Println(tree.PreOrderNR())
	fmt.Println("InOrderNR:")
	fmt.Println(tree.InOrderNR())
	fmt.Println("PostOrderNR:")
	fmt.Println(tree.PostOrderNR())

	//fmt.Println("RemoveMax:")
	//fmt.Println(tree.RemoveMax())
	//fmt.Println(tree.InOrderNR())
	//fmt.Println(tree.RemoveMax())
	//fmt.Println(tree.InOrderNR())
	//
	//fmt.Println("RemoveMin:")
	//fmt.Println(tree.RemoveMin())
	//fmt.Println(tree.InOrderNR())
	//fmt.Println(tree.RemoveMin())
	//fmt.Println(tree.InOrderNR())
	fmt.Println("RemoveElement:")
	fmt.Println(tree.GetSize())
	fmt.Println(tree.RemoveElement(5))
	fmt.Println(tree.GetSize())
	fmt.Println(tree.LevelOrder())
}
