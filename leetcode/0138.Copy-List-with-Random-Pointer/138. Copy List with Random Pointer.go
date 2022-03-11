package solution0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 执行用时：0 ms , 在所有 Go 提交中击败了 100.00% 的用户
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node)
	copyMap := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		nodeMap[cur] = cur
		copyMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	for node := range nodeMap {
		if node.Next != nil {
			copyMap[node].Next = copyMap[node.Next]
		}
		if node.Random != nil {
			copyMap[node].Random = copyMap[node.Random]
		}
	}
	return copyMap[head]
}
