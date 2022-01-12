package bstset

import "LearnGo/data_structures/06-Binary-Search-Tree/bst"

type BSTSet struct {
	bst *bst.BST
}

func NewBSTSet() *BSTSet {
	return &BSTSet{bst: bst.NewBST()}
}

func (set *BSTSet) GetSize() int {
	return set.bst.GetSize()
}

func (set *BSTSet) IsEmpty() bool {
	return set.bst.IsEmpty()
}

func (set *BSTSet) Add(e interface{}) {
	set.bst.Add(e)
}

func (set *BSTSet) Remove(e interface{}) bool {
	return set.bst.RemoveElement(e)
}

func (set *BSTSet) Contains(e interface{}) bool {
	return set.bst.Contains(e)
}
