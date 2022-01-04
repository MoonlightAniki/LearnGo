package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")

	var s []int // Zero value for slice is nil
	fmt.Println(s, s == nil)

	for i := 0; i < 100; i++ {
		s = append(s, i*2+1)
		printSlice(s)
	}

	// slice 的初始化
	fmt.Println("slice的初始化")
	s1 := []int{1, 2, 3, 4}
	printSlice(s1)

	fmt.Println("创建一个长度为16的slice")
	s2 := make([]int, 16)
	printSlice(s2)

	fmt.Println("创建一个长度为16，容量为32的slice")
	s3 := make([]int, 16, 32)
	printSlice(s3)

	fmt.Println("slice的拷贝")
	copy(s2, s1)
	printSlice(s1)
	printSlice(s2)

	fmt.Println("删除slice中的元素")
	// 删除s2中第3个元素, s2[4:]... 表示将 s2[4:] 展开
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("删除s2头尾的元素")
	front := s2[0]
	s2 = s2[1:]
	fmt.Printf("front = %d\n", front)
	printSlice(s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Printf("tail = %d\n", tail)
	printSlice(s2)

}
