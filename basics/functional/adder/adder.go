package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		// sum 自由变量
		sum += v
		return sum
	}
}

// 传统函数式编程，不使用变量、不保留状态
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	//testAdder()

	//testAdder2()

}

func testAdder2() {
	a := adder2(0)
	for i := 1; i <= 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("1 + 2 + ... + %d = %d\n", i, s)
	}
}

func testAdder() {
	a := adder()
	fmt.Printf("%T %v\n", a, a)
	for i := 1; i <= 10; i++ {
		fmt.Println(a(i))
	}
}
