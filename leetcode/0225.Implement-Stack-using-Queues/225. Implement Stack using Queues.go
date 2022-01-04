package solution0225

// region queue
type queue []int

func (q *queue) offer(x int) {
	*q = append(*q, x)
}

func (q *queue) poll() int {
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func (q *queue) peek() int {
	return (*q)[0]
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

// endregion

// region MyStack
type MyStack struct {
	a queue // 输入队列
	b queue // 输出队列
}

func Constructor() MyStack {
	return MyStack{
		a: queue{},
		b: queue{},
	}
}

func (this *MyStack) Push(x int) {
	this.a.offer(x)
	for !this.b.isEmpty() {
		this.a.offer(this.b.poll())
	}
	this.a, this.b = this.b, this.a
}

func (this *MyStack) Pop() int {
	return this.b.poll()
}

func (this *MyStack) Top() int {
	return this.b.peek()
}

func (this *MyStack) Empty() bool {
	return this.b.isEmpty()
}

// endregion
