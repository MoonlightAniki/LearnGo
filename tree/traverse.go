package tree

import "fmt"

// 将二叉树的遍历方法统一放到这个文件中

func (node *Node) InOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.InOrderTraverse()
	fmt.Println(node.Value)
	node.Right.InOrderTraverse()
}

func (node *Node) PreOrderTraverse() {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	node.Left.PreOrderTraverse()
	node.Right.PreOrderTraverse()
}

func (node *Node) PostOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.PostOrderTraverse()
	node.Right.PostOrderTraverse()
	fmt.Println(node.Value)
}
