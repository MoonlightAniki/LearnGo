package main

import "fmt"

// 函数外部不能使用 := 定义变量
//aa := 3

//var aa = 3
//var bb = true
//var ss = "kkk"
// 可将变量定义在一个组内
var (
	aa = 3
	bb = true
	ss = "kkk"
)

func variableZeroValue() {
	// int 初始值 0， string 初始值 ""
	var a int
	var s string
	//fmt.Println(a, s)
	//fmt.Printf("%d, %s\n", a, s)
	// %q 将字符串的引号打印出来
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeDeduction() {
	//var a, b int = 3, 4
	//var bb bool = true
	//var s string = "abc"
	//fmt.Println(a, b, bb, s)
	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, bb, ss)
}
