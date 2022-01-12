package set

type Set interface {
	GetSize() int
	IsEmpty() bool
	Add(e interface{})
	Remove(e interface{}) bool
	Contains(e interface{}) bool
}
