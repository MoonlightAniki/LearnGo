package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(s, len(s))

	// 打印每个byte
	for i, b := range []byte(s) {
		fmt.Printf("i=%d, b=%X\n", i, b)
	}

	// 打印每个rune
	for i, ch := range []rune(s) {
		fmt.Println("i =", i, ", ch =", ch)
	}

	// 打印每个rune
	fmt.Println()
	for i, ch := range s {
		fmt.Println("i =", i, ", ch =", ch)
	}

	count := utf8.RuneCountInString(s)
	fmt.Println("rune count:", count)

	fmt.Println("=====================Decode Rune====================")
	bytes := []byte(s)
	for len(bytes) > 0 {
		// size 为解码出来的这个rune的字节数
		r, size := utf8.DecodeRune(bytes)
		fmt.Println(r, size, string(r))
		bytes = bytes[size:]
	}

}
