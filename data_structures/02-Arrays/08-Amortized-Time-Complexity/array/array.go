package array

import (
	"LearnGo/data_structures/utils"
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{} //interface{} 相当于Java中的 Object
	size int
}

// 构造函数
func New(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
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
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("Add failed, index is out of range.")
	}
	if a.size == len(a.data) {
		a.resize(a.size * 2)
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = e
	a.size++
}

// 向数组头部添加元素 e
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

// 向数组尾部添加元素 e
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

// 获取数组索引 index 位置处的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed, index is out of range.")
	}
	return a.data[index]
}

// 更新索引 index 位置处元素的值为 e
func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed, index is out of range.")
	}
	a.data[index] = e
}

// 判断数组中是否包含元素 e
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			return true
		}
	}
	return false
}

// 返回数组中元素 e 的索引，如果数组中不包含元素 e, 则返回-1
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

// 返回数组中元素 e 所有索引组成的切片，如果不包含元素 e，则返回空切片
func (a *Array) FindAll(e interface{}) []int {
	var indexes []int
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// 删除数组中 index 位置处的元素，并返回被删除的元素
func (a *Array) Remove(index int) interface{} {
	if a.size == 0 {
		panic("Remove failed, array is empty.")
	}
	if index < 0 || index >= a.size {
		panic("Remove failed, index is out of range.")
	}
	e := a.data[index]
	for i := index; i+1 < a.size; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	a.data[a.size] = nil
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除一个元素 e
func (a *Array) RemoveElement(e interface{}) bool {
	if !a.Contains(e) {
		return false
	}
	if index := a.Find(e); index != -1 {
		a.Remove(index)
	}
	return true
}

// 从数组中删除所有元素 e
func (a *Array) RemoveAllElements(e interface{}) bool {
	if !a.Contains(e) {
		return false
	}
	if indexes := a.FindAll(e); len(indexes) != 0 {
		for i := len(indexes) - 1; i >= 0; i-- {
			a.Remove(indexes[i])
		}
	}
	return true
}

// 为数组扩容
func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

// 重写String()方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size=%d, capacity=%d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
