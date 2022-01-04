package solution0232

// region stack
type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

// endregion

// region MyQueue
type MyQueue struct {
	s1 stack
	s2 stack
}

func Constructor() MyQueue {
	return MyQueue{
		s1: stack{},
		s2: stack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.push(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.isEmpty() {
		for !this.s1.isEmpty() {
			this.s2.push(this.s1.pop())
		}
	}
	return this.s2.pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.isEmpty() {
		for !this.s1.isEmpty() {
			this.s2.push(this.s1.pop())
		}
	}
	return this.s2.peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.isEmpty() && this.s2.isEmpty()
}

// endregion
