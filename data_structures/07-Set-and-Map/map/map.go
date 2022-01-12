package mymap

type Map interface {
	GetSize() int

	IsEmpty() bool

	Add(key interface{}, value interface{})

	Set(key interface{}, value interface{})

	Get(key interface{}) interface{}

	Remove(key interface{})

	Contains(key interface{}) bool
}
