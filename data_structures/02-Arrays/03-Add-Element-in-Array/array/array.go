package array

type Array struct {
	data []int
	size int
}

func New(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
	}
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) Add(index int, e int) {
	if a.size == len(a.data) {
		panic("Add failed, array is full.")
	}
	if index < 0 || index > a.size {
		panic("Add failed, index is out of range.")
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = e
	a.size++
}

func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}
