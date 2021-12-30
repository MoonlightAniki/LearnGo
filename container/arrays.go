package main

import "fmt"

func main() {
	// 元素个数写在元素类型签名
	var arr1 [5]int
	fmt.Println(arr1)

	// 需要指定数组内元素个数
	arr2 := [4]int{1, 2, 3, 4}
	fmt.Println(arr2)

	// 使用 ... 自动获取元素个数
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr3)

	// [4]int 与 [5]int 是不同的类型
	//arr3 = [4]int{1, 3, 5, 7}

	// 二维数组
	var grid [4][5]int
	fmt.Println(grid)

	// 遍历一维数组
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("i = %d, v = %d\n", i, arr2[i])
	}

	// 遍历二维数组
	fmt.Println("遍历二维数组")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("i = %d, j = %d, v = %d\n", i, j, grid[i][j])
		}
	}

	fmt.Println("使用range遍历一维数组")
	for i, v := range arr3 {
		fmt.Printf("i = %d, v = %d\n", i, v)
	}

	fmt.Println("使用range遍历二维数组")
	for i, row := range grid {
		for j, v := range row {
			fmt.Printf("i = %d, j = %d, v = %d\n", i, j, v)
		}
	}

	printArray(arr1)
	// 调用 printArray 之后 arr1 中第一个元素还是0，说明数组是值传递
	fmt.Println(arr1)
	// [4]int 与 [5]int 是不同的类型
	//printArray(arr2)
	printArray(arr3)
	fmt.Println(arr3)

	// 可以使用数组指针来修改数组中的元素
	// 但是go语言中一般不直接使用数组，而是使用切片
	printArray2(&arr1)
	printArray2(&arr3)
	fmt.Println(arr1, arr3)
}

// []int 表示一个数组切片, [5]int表示一个元素数量为5的int数组
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func printArray2(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
