package main

import (
	solution0297 "LearnGo/leetcode/0297.Serialize-and-Deserialize-Binary-Tree"
	"LearnGo/leetcode/kit"
	"fmt"
)

func main() {
	root := &kit.TreeNode{Val: 3}
	root.Left = &kit.TreeNode{Val: 2}
	root.Right = &kit.TreeNode{Val: 4}
	root.Left.Left = &kit.TreeNode{Val: 3}
	codec := solution0297.Codec{}
	data := codec.Serialize(root)
	fmt.Println(data)
	node := codec.Deserialize("nil")
	fmt.Println(node)
}
