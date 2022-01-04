package main

import (
	"LearnGo/data_structures/02-Arrays/06-Generic-Data-Structures/array"
	"LearnGo/data_structures/02-Arrays/06-Generic-Data-Structures/student"
	"fmt"
)

func main() {
	// 测试 student 类型
	students := array.New(10)
	students.AddLast(student.New("Alice", 100))
	students.AddLast(student.New("Bob", 66))
	students.AddLast(student.New("Charlie", 88))
	fmt.Println(students)
}
