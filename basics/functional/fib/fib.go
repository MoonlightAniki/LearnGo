package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 为函数实现 Reader 接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 将next转成字符串，代理
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	/*	fmt.Println(f()) // 1
		fmt.Println(f()) // 1
		fmt.Println(f()) // 2
		fmt.Println(f()) // 3
		fmt.Println(f()) // 5
		fmt.Println(f()) // 8*/

	printFileContents(f)
}
