package main

/*func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				// Printf 是 IO 操作
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}*/

// 协程 Coroutine
// 轻量级"线程"
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器、解释器、虚拟机层面的多任务
// 多个协程可能在一个或多个线程上运行

/*func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// 无法交出控制权
				a[i]++

				// 手动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}*/

// go run -race goroutine.go
// 检测数据访问冲突
/*func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() { // race condition!
			for {
				a[i]++
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}*/

// goroutine 可能的切换点
// IO / select
// channel
// 等待锁
// 函数调用(有时)
// runtime.Gosched()

/*func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}*/
