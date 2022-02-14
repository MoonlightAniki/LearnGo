package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*func msgGen() <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("message %d", i)
			i++
		}
	}()
	return c
}

func main() {
	m := msgGen()
	for {
		fmt.Println(<-m)
		//m <- "abc"
	}
}

func main() {
	m1, m2 := msgGen(), msgGen()
	for {
		fmt.Println(<-m1)
		fmt.Println(<-m2)
	}
}*/

/*func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func main() {
	m1, m2 := msgGen("service1"), msgGen("service2")

	for {
		fmt.Println(<-m1)
		fmt.Println(<-m2)
	}
}
*/

/*func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

//func fanIn(c1, c2 chan string) chan string {
//	c := make(chan string)
//	go func() {
//		for {
//			c <- <-c1
//		}
//	}()
//	go func() {
//		for {
//			c <- <-c2
//		}
//	}()
//	return c
//}

//func fanIn(chs ...chan string) chan string {
//	c := make(chan string)
//	for _, ch := range chs {
//		chCopy := ch
//		go func() {
//			for {
//				c <- <-chCopy
//			}
//		}()
//	}
//	return c
//}

// 适合不知道 channel 的数量的情况
func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

// 适合 知道 channel 数量的情况
func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()
	return c
}

func main() {
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m3 := msgGen("service3")
	//m := fanIn(m1, m2)
	//m := fanInBySelect(m1, m2)
	m := fanIn(m1, m2, m3)
	for {
		fmt.Println(<-m)
	}
}*/

//func msgGen(name string) chan string {
//	c := make(chan string)
//	go func() {
//		i := 0
//		for {
//			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
//			c <- fmt.Sprintf("service %s: message %d", name, i)
//			i++
//		}
//	}()
//	return c
//}

//func nonBlockingWait(c chan string) (string, bool) {
//	select {
//	case s := <-c:
//		return s, true
//	default: //非阻塞
//		return "", false
//	}
//}

//func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
//	select {
//	case s := <-c:
//		return s, true
//	case <-time.After(timeout):
//		return "", false
//	}
//}

//func main() {
//	m1 := msgGen("service1")
//	m2 := msgGen("service2")
//
//	for {
//		fmt.Println(<-m1)
//		if s, ok := nonBlockingWait(m2); ok {
//			fmt.Println(s)
//		} else {
//			fmt.Println("No message from service2")
//		}
//	}
//}

//func main() {
//	m := msgGen("service1")
//
//	for {
//		if s, ok := timeoutWait(m, time.Second); ok {
//			fmt.Println(s)
//		} else {
//			fmt.Println("timeout")
//		}
//	}
//}

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleanup end")
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case s := <-c:
		return s, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})
	m := msgGen("service1", done)

	for i := 0; i < 5; i++ {
		if s, ok := timeoutWait(m, time.Second); ok {
			fmt.Println(s)
		} else {
			fmt.Println("timeout")
		}
	}
	done <- struct{}{}
	<-done
}
