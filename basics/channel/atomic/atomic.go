package main

import (
	"fmt"
	"sync"
	"time"
)

/*type atomicInt int

func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

func main() {
	var aInt atomicInt = 0
	aInt.increment()
	go func() {
		aInt.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(aInt)
}
*/

type atomicInt struct {
	value int
	mutex sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("atomicInt increment")
	func() {
		a.mutex.Lock()
		defer a.mutex.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.value
}

func main() {
	aInt := atomicInt{
		value: 0,
		mutex: sync.Mutex{},
	}
	aInt.increment()
	go func() {
		aInt.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(aInt.get())
}
