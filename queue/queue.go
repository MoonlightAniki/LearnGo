package queue

import "strconv"

// 使用别名的方式扩展已有类型
// 使用Queue作为[]int的别名
type Queue []int

func (q *Queue) String() string {
	if q == nil {
		return "nil"
	}
	res := ""
	res += "font: ["
	for i, v := range *q {
		res += strconv.Itoa(v)
		if i != len(*q)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
