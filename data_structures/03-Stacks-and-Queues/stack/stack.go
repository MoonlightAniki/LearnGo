package stack

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	GetSize() int
}
