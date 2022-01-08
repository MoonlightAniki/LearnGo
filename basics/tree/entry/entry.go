package main

import (
	"LearnGo/basics/tree"
	"fmt"
)

// 扩展系统类型或者别人的类型的方法有两种：
// 1.定义别名
// 2.使用组合

// 演示如何使用组合扩展别人的类型：
// 假设 tree.Node 是别人定义的类型，并且只实现了中序遍历，使用组合的方式为 tree.Node 增加前序遍历和后续遍历
type myTreeNode struct {
	node *tree.Node
}

func (node *myTreeNode) postOrder() {
	if node == nil || node.node == nil {
		return
	}
	left := myTreeNode{node.node.Left}
	left.postOrder()
	right := myTreeNode{node.node.Right}
	right.postOrder()
	node.node.Print()
}

func main() {
	//test001()

	root := tree.CreateNode(1)
	root.Left = tree.CreateNode(2)
	root.Right = tree.CreateNode(3)
	root.Left.Left = tree.CreateNode(4)
	root.Left.Right = tree.CreateNode(5)
	root.Right.Left = tree.CreateNode(6)
	root.Right.Right = tree.CreateNode(7)
	root.InOrderTraverse()

	fmt.Println()

	myRoot := myTreeNode{root}
	myRoot.postOrder()
}

func test001() {
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
