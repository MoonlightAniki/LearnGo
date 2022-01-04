package main

import "fmt"

// Slice 本身是没有数据的，是底层数组的一个 View 视图
func updateSlice(s []int) {
	s[0] = 100
}

func printArray(arr []int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)

	fmt.Println(arr[2:6]) //切片，左闭右开
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Printf("Before updateSlice, s1 = %v, s2 = %v\n", s1, s2)
	updateSlice(s1)
	updateSlice(s2)
	// 调用 updateSlice 之后 s1 和 s2 的值更新了，说明切片不是值类型
	// 数组是值类型
	fmt.Printf("After updateSlice, s1 = %v, s2 = %v\n", s1, s2)
	// 更新切片的值之后，数组的值也跟着改变了。
	fmt.Printf("After updateSlice, arr = %v\n", arr)

	arr2 := [5]int{1, 3, 5, 7, 9}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr2, arr3)

	printArray(arr2[:])
	printArray(arr3[:])
	fmt.Println(arr2, arr3)

	// Reslice
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	extendingSlice()
}

func extendingSlice() {
	// Extending Slice
	fmt.Println("Extending slice")

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr =", arr)

	s1 := arr[2:6] // [2, 3, 4, 5], len=4, cap=6
	// s1 中有4个元素，s2 对 s1 向后扩展了一个元素
	s2 := s1[3:5] // [5, 6], len=2, cap=3
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	// s[i], i 需满足 i < len(s)
	//fmt.Println(s1[4])

	// 向后扩展不能超过 cap(s)
	//fmt.Println(s1[3:7])

	s3 := append(s2, 10) // [5, 6, 10], len=3, cap=3
	s4 := append(s3, 20) // [5, 6, 10, 20], len=4, cap=6, cap扩容为原来的2倍
	s5 := append(s4, 30) // [5, 6, 10, 20, 30], len=5, cap=6
	fmt.Println("s3, s4, s5 =", sliceToString(s3), sliceToString(s4), sliceToString(s5))
}

func sliceToString(s []int) string {
	return fmt.Sprintf("%v, len=%d, cap=%d", s, len(s), cap(s))
}
