package solution0297

import (
	"LearnGo/leetcode/kit"
	"strconv"
	"strings"
)

/*
Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

Clarification: The input/output format is the same as how LeetCode serializes a binary tree. You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.



Example 1:
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]


Example 2:
Input: root = []
Output: []


Constraints:
The number of nodes in the tree is in the range [0, 10^4].
-1000 <= Node.val <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode = kit.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// 执行用时： 12 ms , 在所有 Go 提交中击败了 47.06% 的用户
/*// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(fmt.Sprintf("%d,", node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, err := strconv.Atoi(sp[0])
		if err != nil {
			panic(err)
		}
		sp = sp[1:]
		return &TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}*/

func (this *Codec) serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	if root == nil {
		sb.WriteString("nil")
		return sb.String()
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			sb.WriteString("nil")
		} else {
			sb.WriteString(strconv.Itoa(cur.Val))
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
		sb.WriteString(",")
	}
	return sb.String()
}

func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	if arr[0] == "nil" {
		return nil
	}
	rootVal, _ := strconv.Atoi(arr[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == nil {
			continue
		}
		if arr[i] != "nil" {
			val, _ := strconv.Atoi(arr[i])
			cur.Left = &TreeNode{Val: val}
		} else {
			cur.Left = nil
		}
		if arr[i+1] != "nil" {
			val, _ := strconv.Atoi(arr[i+1])
			cur.Right = &TreeNode{Val: val}
		} else {
			cur.Right = nil
		}
		i += 2
		queue = append(queue, cur.Left, cur.Right)
	}
	return root
}

func (this *Codec) Serialize(root *TreeNode) string {
	return this.serialize(root)
}

func (this *Codec) Deserialize(data string) *TreeNode {
	return this.deserialize(data)
}
