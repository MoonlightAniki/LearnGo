package solution0307

import "LearnGo/data_structures/09-Segment-Tree/segmenttree"

type NumArray struct {
	tree *segmenttree.SegmentTree
}

func Constructor(nums []int) NumArray {
	numArry := NumArray{}
	numArry.tree = segmenttree.NewSegmentTree(nums, func(a, b int) int { return a + b })
	return numArry
}

func (this *NumArray) Update(index int, val int) {
	this.tree.Update(index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.tree.Query(left, right)
}
