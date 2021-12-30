package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数可以有多个返回值
//func div(a, b int) (int, int) {
//	return a / b, a % b
//}

// 可以给函数返回值命名，提高代码可读性
// q - quotient 商
// r - remainder 余数
//func div(a int, b int) (q int, r int) {
//	return a / b, a % b
//}

// 第三种写法
// 可读性较差，不建议使用这种写法
func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	//case "/":
	//	return a / b
	case "/":
		// 不需要使用 r，使用 _ 将其省略
		q, _ := div(a, b)
		return q
	default:
		panic("Unsupported operation: " + op)
	}
}

// go语言风格的写法
func eval2(a, b int, op string) (res int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("Unsupported operation: %s\n", op)
	}
}

// 使用函数式编程实现eval函数
func eval3(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数
func sum2(nums ...int) int {
	sum := 0
	//for _, v := range nums {
	//	sum += v
	//}
	// 或者
	for i, _ := range nums {
		sum += nums[i]
	}
	return sum
}

// 交换两个数的值

// go语言中只有值传递
func swap(a, b int) {
	a, b = b, a
}

// 使用指针交换两个数
func swap2(a, b *int) {
	// 交换 a, b 的地址
	*a, *b = *b, *a
}

// 使用函数多返回值实现
func swap3(a, b int) (int, int) {
	return b, a
}

func main() {
	//fmt.Println(div(13, 3))

	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(
		eval(3, 4, "+"),
		eval(3, 4, "-"),
		eval(3, 4, "*"),
		eval(3, 4, "/"),
		//eval(3, 4, "//"),
	)

	if res, err := eval2(3, 4, "+"); err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	res := eval3(pow, 3, 4)
	fmt.Println(res)

	// 定义匿名函数
	fmt.Println(
		eval3(
			// 匿名函数
			func(a, b int) int {
				return int(math.Sqrt(float64(a*a + b*b)))
			},
			3,
			4,
		),
	)

	fmt.Println(sum2(1, 2, 3, 4, 5))

	a, b := 3, 4
	//fmt.Printf("Before swapping, a = %d, b = %d\n", a, b)
	//// 值传递无法实现交换变量的值
	//swap(a, b)
	//fmt.Printf("After swapping, a = %d, b = %d\n", a, b)

	//fmt.Printf("Before swapping, a = %d, b = %d\n", a, b)
	//swap2(&a, &b)
	//fmt.Printf("After swapping, a = %d, b = %d\n", a, b)

	fmt.Printf("Before swapping, a = %d, b = %d\n", a, b)
	a, b = swap3(a, b)
	fmt.Printf("After swapping, a = %d, b = %d\n", a, b)

}
