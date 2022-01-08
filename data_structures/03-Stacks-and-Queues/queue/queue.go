package queue

type Queue interface {
	Offer(e interface{})
	Poll() interface{}
	Peek() interface{}
	IsEmpty() bool
	GetSize() int
}
