package main

// import 也可以写在一个组内
import (
	"fmt"
	"math"
	"math/cmplx"
)

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
	// 变量类型可以省略
	var a, b, c, d = 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	// := 变量声明和赋值
	a, b, c, d := 3, 4, true, "abc"
	fmt.Println(a, b, c, d)
}

// 验证欧拉公式 pow(e, i * PI) + 1 = 0
func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // 输出 (0+1.2246467991473515e-16i)
	//fmt.Println(cmplx.Exp(1i*math.Pi) + 1) // 输出 (0+1.2246467991473515e-16i)
	fmt.Printf("%3f\n", cmplx.Exp(1i*math.Pi)+1) // (0.000000+0.000000i)
}

// go语言中类型转换是强制的
func triangle() {
	//var a int = 3
	//var b int = 4
	a, b := 3, 4
	var c int
	// a * a + b * b 不能自动转换成 float64, float64 也不能自动转成 int
	//c = math.Sqrt(a * a + b * b)
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()

	fmt.Println(aa, bb, ss)
}
