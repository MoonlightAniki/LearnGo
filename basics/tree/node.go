package tree

import "fmt"

// go语言支持封装，不支持继承和多态
type Node struct {
	Value       int
	Left, Right *Node
}

// 定义成员函数
func (node Node) Print() {
	fmt.Println(node.Value)
}

// 等价于
//func Print(node Node) {
//	fmt.Println(node.Value)
//}

// 值传递，无法修改node的value
//func (node Node) SetValue(Value int) {
//	node.Value = Value
//}

// 指针版本
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node, ignored!")
		return
	}
	node.Value = value
}

// go语言中没有构造函数，使用工厂函数代替
func CreateNode(value int) *Node {
	return &Node{Value: value} //返回一个局部变量的地址，在C++是不合法的！
}
