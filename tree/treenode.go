package main

import "fmt"

// go语言支持封装，不支持继承和多态
type treeNode struct {
	value       int
	left, right *treeNode
}

// 定义成员函数
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 等价于
//func print(node treeNode) {
//	fmt.Println(node.value)
//}

// 值传递，无法修改node的value
//func (node treeNode) setValue(value int) {
//	node.value = value
//}

// 指针版本
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node, ignored!")
		return
	}
	node.value = value
}

// go语言中没有构造函数，使用工厂函数代替
func createNode(value int) *treeNode {
	return &treeNode{value: value} //返回一个局部变量的地址，在C++是不合法的！
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Println(node.value)
	node.right.traverse()
}

func main() {
	var root treeNode // zero value is {0 <nil> <nil>}
	fmt.Println(root)

	root = treeNode{value: 3} // 第一种创建方式
	fmt.Println(root)

	/*
	      3
	    /  \
	   0    5
	       /  \
	     0     6
	*/
	root.left = &treeNode{} // 第二种创建方式，所有成员都使用默认初始值
	root.right = &treeNode{value: 5}
	root.right.left = new(treeNode) // 第三种创建方式与第二种创建方式等价，不过new返回的为结构体的指针
	fmt.Println(root.right.left)
	root.right.right = createNode(6) // 第四种创建方式：使用工厂方法

	// 调用成员函数
	root.print()
	fmt.Println()

	// 调用setValue之后root的value没有变为100，值传递
	root.setValue(100)
	root.print()
	fmt.Println()

	// 创建 treeNode 切片
	nodes := []treeNode{
		{value: 1}, // treeNode{value: 1} 省略 treeNode
		{3, nil, nil},
		{5, nil, &root},
	}
	fmt.Println(nodes)

	// nil可以调用函数
	var pRoot *treeNode
	println(pRoot)
	pRoot.setValue(100)
	pRoot = &treeNode{value: 3}
	pRoot.print()
	fmt.Println()

	root.traverse() // 0, 100, 0, 5, 6

}
