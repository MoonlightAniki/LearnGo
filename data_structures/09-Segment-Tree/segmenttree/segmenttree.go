package segmenttree

type node struct {
	start int
	end   int
	value int
	left  *node
	right *node
}

type SegmentTree struct {
	root   *node
	data   []int
	merger func(a int, b int) int
}

func NewSegmentTree(data []int, merger func(a int, b int) int) *SegmentTree {
	tree := &SegmentTree{}
	tree.root = buildSegmentTree(data, 0, len(data)-1, merger)
	tree.data = data
	tree.merger = merger
	return tree
}

func (st *SegmentTree) Query(queryL int, queryR int) int {
	if queryL > queryR || queryL < 0 || queryL >= len(st.data) || queryR < 0 || queryR >= len(st.data) {
		panic("Query failed, illegal arguments.")
	}
	return st.query(st.root, queryL, queryR)
}

func (st *SegmentTree) Update(index int, value int) {
	if index < 0 || index >= len(st.data) {
		panic("Update failed, illegal arguments.")
	}
	st.update(st.root, index, value)
}

func buildSegmentTree(
	data []int,
	l int,
	r int,
	merger func(a int, b int) int) *node {
	if l > r {
		return nil
	}
	if l == r {
		return &node{start: l, end: r, value: data[l]}
	}
	mid := l + (r-l)/2
	root := &node{}
	root.start = l
	root.end = r
	root.left = buildSegmentTree(data, l, mid, merger)
	root.right = buildSegmentTree(data, mid+1, r, merger)
	root.value = merger(root.left.value, root.right.value)
	return root
}

func (st *SegmentTree) query(n *node, queryL int, queryR int) int {
	if queryL == n.start && queryR == n.end {
		return n.value
	}
	mid := n.start + (n.end-n.start)/2
	if queryR <= mid {
		return st.query(n.left, queryL, queryR)
	} else if queryL >= mid+1 {
		return st.query(n.right, queryL, queryR)
	} else {
		resultL := st.query(n.left, queryL, mid)
		resultR := st.query(n.right, mid+1, queryR)
		return st.merger(resultL, resultR)
	}
}

func (st *SegmentTree) update(n *node, index int, value int) {
	if n.start == n.end && n.start == index {
		n.value = value
		st.data[index] = value
		return
	}
	mid := n.start + (n.end-n.start)/2
	if index <= mid {
		st.update(n.left, index, value)
	} else {
		st.update(n.right, index, value)
	}
	n.value = st.merger(n.left.value, n.right.value)
}
