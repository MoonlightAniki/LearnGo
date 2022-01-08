package sortTestHelper

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintArray(arr []int) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// 拷贝int数组
func CopyIntArray(arr []int) []int {
	dst := make([]int, len(arr))
	copy(dst, arr)
	return dst
}

// 生成有 n 个元素的随机数组，每个元素的取值返回 [rangeL...rangeR]
func GenerateRandomArray(n int, rangeL int, rangeR int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rangeL + rand.Intn(rangeR-rangeL+1)
	}
	return arr
}

// 生成一个近乎有序的数组
func GenerateNearlyOrderedArray(n int, swapTimes int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	for i := 0; i < swapTimes; i++ {
		x := rand.Intn(n)
		y := rand.Intn(n)
		arr[x], arr[y] = arr[y], arr[x]
	}
	return arr
}

func isSorted(arr []int) bool {
	for i := 0; i+1 < len(arr); i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func TestSort(sortName string, sortFunc func([]int), arr []int) {
	startTime := time.Now().UnixMilli()
	sortFunc(arr)
	endTime := time.Now().UnixMilli()
	fmt.Printf("%s : %d ms\n", sortName, endTime-startTime)

	if !isSorted(arr) {
		panic("arr is not sorted.")
	}
}
