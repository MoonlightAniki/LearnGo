package array

import (
	"bytes"
	"fmt"
	"strconv"
)

type Array struct {
	data []int
	size int
}

// 构造函数
func New(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
	}
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获取数组中元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 判断数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 向数组中 index 索引位置添加元素 e
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

// 向数组头部添加元素 e
func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

// 向数组尾部添加元素 e
func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}

// 获取数组索引 index 位置处的元素
func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("Get failed, index is out of range.")
	}
	return a.data[index]
}

// 更新索引 index 位置处元素的值为 e
func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("Set failed, index is out of range.")
	}
	a.data[index] = e
}

// 重写String()方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size=%d, capacity=%d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(strconv.Itoa(a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
