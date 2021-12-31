package main

import (
	"LearnGo/tree"
	"fmt"
)

func main() {
	// 不在一个包内，需要加上包名
	var root tree.Node // zero Value is {0 <nil> <nil>}
	fmt.Println(root)

	root = tree.Node{Value: 3} // 第一种创建方式
	fmt.Println(root)

	/*
	      3
	    /  \
	   0    5
	       /  \
	     0     6
	*/
	root.Left = &tree.Node{} // 第二种创建方式，所有成员都使用默认初始值
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node) // 第三种创建方式与第二种创建方式等价，不过new返回的为结构体的指针
	fmt.Println(root.Right.Left)
	root.Right.Right = tree.CreateNode(6) // 第四种创建方式：使用工厂方法

	// 调用成员函数
	root.Print()
	fmt.Println()

	// 调用setValue之后root的value没有变为100，值传递
	root.SetValue(100)
	root.Print()
	fmt.Println()

	// 创建 TreeNode 切片
	nodes := []tree.Node{
		{Value: 1}, // TreeNode{Value: 1} 省略 TreeNode
		{3, nil, nil},
		{5, nil, &root},
	}
	fmt.Println(nodes)

	// nil可以调用函数
	var pRoot *tree.Node
	println(pRoot)
	pRoot.SetValue(100)
	pRoot = &tree.Node{Value: 3}
	pRoot.Print()
	fmt.Println()

	root.InOrderTraverse() // 0, 100, 0, 5, 6

}
