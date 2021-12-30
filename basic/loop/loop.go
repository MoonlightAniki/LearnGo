package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sum(n int) int {
	sum := 0
	// go 语言不支持 ++i ？
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func convertToBinaryString(n int) string {
	if n < 0 {
		panic(fmt.Sprintf("Illegal argument: %d", n))
	}
	if n == 0 {
		return "0"
	}
	res := ""
	// 初始条件可以省略
	for ; n > 0; n /= 2 {
		r := n % 2
		// strconv.Itoa int 转 string
		res = strconv.Itoa(r) + res
	}
	return res
}

// 逐行打印文件内容
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	// 无循环初始条件和循环递增条件时可以省略分号
	// 相当于java中的while语句，go语言中没有while关键字
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 死循环
func forever() {
	// true 可以省略
	for {
		fmt.Println("abc")
	}
}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(sum(100))

	fmt.Println(
		convertToBinaryString(0),  // 0
		convertToBinaryString(3),  // 11
		convertToBinaryString(11), // 1011
		convertToBinaryString(13), // 1101
		convertToBinaryString(32), // 100000
	)

	printFile("abc.txt")

	//forever()

	// 多行字符串
	s := `Hello golang
aa
bb
  cc
    dd`
	// 将字符串转成reader
	printFileContent(strings.NewReader(s))
}
