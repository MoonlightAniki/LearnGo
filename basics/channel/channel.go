package main

import (
	"fmt"
	"time"
)

/*func channelDemo() {
	//var c chan int // c == nil
	c := make(chan int)
	// 向 channel 发送数据
	c <- 1
	c <- 2

	// 从 channel 接收数据
	n := <-c
	fmt.Println(n)
}*/

/*func channelDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}*/

/*func worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func channelDemo() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Millisecond)
}*/

/*func worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Printf("goroutine %d received %c\n", id, n)
	}
}

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}*/

/*func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}*/

/*func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}*/

/*func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}*/

func worker(id int, c chan int) {
	/*	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d received %d\n", id, n)
	}*/

	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	//channelDemo()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
