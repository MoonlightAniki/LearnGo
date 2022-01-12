package bst

import (
	"LearnGo/data_structures/utils"
)

type node struct {
	value interface{}
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func NewBST() *BST {
	return &BST{}
}

/*返回二分搜索树的节点数*/
func (bst *BST) GetSize() int {
	return bst.size
}

/*判断二分搜索树是否为空*/
func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

/*向二分搜索树添加元素, 如果已经有相同的元素，则不添加*/
func (bst *BST) Add(e interface{}) {
	bst.root = bst.add(bst.root, e)
}

func (bst *BST) add(root *node, e interface{}) *node {
	if root == nil {
		bst.size++
		return &node{value: e}
	}
	if utils.Compare(e, root.value) < 0 {
		root.left = bst.add(root.left, e)
	} else if utils.Compare(e, root.value) > 0 {
		root.right = bst.add(root.right, e)
	} else {
		// e == root.value, ignore
	}
	return root
}

/*判断二分搜索树中是否包含元素e*/
func (bst *BST) Contains(e interface{}) bool {
	return bst.contains(bst.root, e)
}

func (bst *BST) contains(root *node, e interface{}) bool {
	if root == nil {
		return false
	}
	if utils.Compare(e, root.value) < 0 {
		return bst.contains(root.left, e)
	} else if utils.Compare(e, root.value) > 0 {
		return bst.contains(root.right, e)
	} else {
		return true
	}
}

/*前序遍历，递归实现*/
func (bst *BST) PreOrder() []interface{} {
	res := make([]interface{}, 0)
	bst.preorder(bst.root, &res)
	return res
}

func (bst *BST) preorder(root *node, res *[]interface{}) {
	if root == nil {
		return
	}
	*res = append(*res, root.value)
	bst.preorder(root.left, res)
	bst.preorder(root.right, res)
}

/*中序遍历，递归实现*/
func (bst *BST) InOrder() []interface{} {
	res := make([]interface{}, 0)
	bst.inorder(bst.root, &res)
	return res
}

func (bst *BST) inorder(root *node, res *[]interface{}) {
	if root == nil {
		return
	}
	bst.inorder(root.left, res)
	*res = append(*res, root.value)
	bst.inorder(root.right, res)
}

/*后序遍历，递归实现*/
func (bst *BST) PostOrder() []interface{} {
	res := make([]interface{}, 0)
	bst.postorder(bst.root, &res)
	return res
}

func (bst *BST) postorder(root *node, res *[]interface{}) {
	if root == nil {
		return
	}
	bst.postorder(root.left, res)
	bst.postorder(root.right, res)
	*res = append(*res, root.value)
}

/*层序遍历*/
func (bst *BST) LevelOrder() [][]interface{} {
	res := make([][]interface{}, 0)
	if bst.root == nil {
		return res
	}
	queue := []*node{bst.root}
	for len(queue) > 0 {
		count := len(queue)
		level := make([]interface{}, 0)
		for i := 0; i < count; i++ {
			front := queue[i]
			level = append(level, front.value)
			if front.left != nil {
				queue = append(queue, front.left)
			}
			if front.right != nil {
				queue = append(queue, front.right)
			}
		}
		queue = queue[count:]
		res = append(res, level)
	}
	return res
}

/*前序遍历，非递归实现*/
func (bst *BST) PreOrderNR() []interface{} {
	res := make([]interface{}, 0)
	if bst.root == nil {
		return res
	}
	stack := []*node{bst.root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.left == nil && top.right == nil {
			res = append(res, top.value)
		} else {
			if top.right != nil {
				stack = append(stack, top.right)
			}
			if top.left != nil {
				stack = append(stack, top.left)
			}
			stack = append(stack, &node{value: top.value})
		}
	}
	return res
}

/*中序遍历，非递归实现*/
func (bst *BST) InOrderNR() []interface{} {
	res := make([]interface{}, 0)
	if bst.root == nil {
		return res
	}
	stack := []*node{bst.root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.left == nil && top.right == nil {
			res = append(res, top.value)
		} else {
			if top.right != nil {
				stack = append(stack, top.right)
			}
			stack = append(stack, &node{value: top.value})
			if top.left != nil {
				stack = append(stack, top.left)
			}
		}
	}
	return res
}

/*后序遍历，非递归实现*/
func (bst *BST) PostOrderNR() []interface{} {
	res := make([]interface{}, 0)
	if bst.root == nil {
		return res
	}
	stack := []*node{bst.root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.left == nil && top.right == nil {
			res = append(res, top.value)
		} else {
			stack = append(stack, &node{value: top.value})
			if top.right != nil {
				stack = append(stack, top.right)
			}
			if top.left != nil {
				stack = append(stack, top.left)
			}
		}
	}
	return res
}

/* 寻找二分搜索树的最大元素 */
func (bst *BST) Maximum() interface{} {
	if bst.IsEmpty() {
		panic("Maximum failed, bst is empty.")
	}
	return bst.maximum(bst.root).value
}

func (bst *BST) maximum(root *node) *node {
	if root.right == nil {
		return root
	} else {
		return bst.maximum(root.right)
	}
}

/*寻找二分搜索树的最小元素*/
func (bst *BST) Minimum() interface{} {
	if bst.IsEmpty() {
		panic("Minimum failed, bst is empty.")
	}
	return bst.minimum(bst.root).value
}

func (bst *BST) minimum(root *node) *node {
	if root.left == nil {
		return root
	} else {
		return bst.minimum(root.left)
	}
}

/*删除二分搜索树最大元素所在的节点，并返回最大元素*/
func (bst *BST) RemoveMax() interface{} {
	if bst.IsEmpty() {
		panic("RemoveMax failed, bst is empty.")
	}
	max := bst.Maximum()
	bst.root = bst.removeMax(bst.root)
	return max
}

/*从以root为根节点的二分搜索树中删除最大元素所在的节点，并返回新的根节点*/
func (bst *BST) removeMax(root *node) *node {
	if root.right == nil {
		left := root.left
		root.left = nil
		bst.size--
		return left
	}
	root.right = bst.removeMax(root.right)
	return root
}

/*删除二分搜索树最小元素所在的节点，并返回最小元素*/
func (bst *BST) RemoveMin() interface{} {
	if bst.IsEmpty() {
		panic("RemoveMin failed, bst is empty.")
	}
	min := bst.Minimum()
	bst.root = bst.removeMin(bst.root)
	return min
}

/*从以root为根节点的二分搜索树中删除最小元素所在的节点，并返回新的根节点*/
func (bst *BST) removeMin(root *node) *node {
	if root.left == nil {
		right := root.right
		root.right = nil
		bst.size--
		return right
	}
	root.left = bst.removeMin(root.left)
	return root
}

/*二分搜索树中如果存在等于e的节点，删除并返回true，否则返回false*/
func (bst *BST) RemoveElement(e interface{}) bool {
	if bst.Contains(e) {
		bst.root = bst.removeElement(bst.root, e)
		return true
	} else {
		return false
	}
}

/*从以root为根节点的二分搜索树中删除值等于e的节点，并返回新的根节点*/
func (bst *BST) removeElement(root *node, e interface{}) *node {
	if root == nil {
		return nil
	}
	if utils.Compare(e, root.value) > 0 {
		root.right = bst.removeElement(root.right, e)
		return root
	} else if utils.Compare(e, root.value) < 0 {
		root.left = bst.removeElement(root.left, e)
		return root
	} else {
		leftChild, rightChild := root.left, root.right
		root.left, root.right = nil, nil
		if rightChild == nil {
			bst.size--
			return leftChild
		} else {
			// 将右子树的最小节点作为新的根节点
			successor := bst.minimum(rightChild)
			successor.right = bst.removeMin(rightChild)
			successor.left = leftChild
			return successor
		}
	}
}
