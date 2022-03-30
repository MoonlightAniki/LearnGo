package solution0284

type PeekingIterator struct {
	iter         *Iterator
	_nextElement int
	_hasNext     bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, iter.next(), iter.hasNext()}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	ret := this._nextElement
	this._hasNext = this.iter.hasNext()
	if this._hasNext {
		this._nextElement = this.iter.next()
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	return this._nextElement
}
