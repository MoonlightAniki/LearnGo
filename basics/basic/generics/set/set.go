package set

import (
	"bytes"
	"fmt"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elements ...T) *Set[T] {
	s := Set[T]{}
	for _, e := range elements {
		s.Add(e)
	}
	return &s
}

func (this *Set[T]) Len() int {
	return len(*this)
}

func (this *Set[T]) IsEmpty() bool {
	return this.Len() == 0
}

func (this *Set[T]) Add(elements ...T) {
	for _, e := range elements {
		(*this)[e] = struct{}{}
	}
}

func (this *Set[T]) Remove(elements ...T) {
	for _, e := range elements {
		delete(*this, e)
	}
}

func (this *Set[T]) Contains(e T) bool {
	_, ok := (*this)[e]
	return ok
}

func (this *Set[T]) ToSlice() []T {
	ret := make([]T, 0)
	for k := range *this {
		ret = append(ret, k)
	}
	return ret
}

func (this *Set[T]) String() string {
	buffer := bytes.Buffer{}
	index := 0
	for k := range *this {
		if index != 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString(fmt.Sprintf("%v", k))
		index++
	}
	return buffer.String()
}
