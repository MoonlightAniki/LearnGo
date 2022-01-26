package v1

import "fmt"

var memo [][]int

// 重量 w, 价值 v， 容积 C
// 取n件物品装入容积为C的背包中，使得物品的价值最大
func knapsack01(w []int, v []int, C int) int {
	memo = make([][]int, len(w))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			memo[i][j] = -1
		}
	}
	return bestValue(w, v, len(w)-1, C)
}

// 从 [0...index] 范围内取n件物品，装入容积为c的背包中
func bestValue(w []int, v []int, index int, c int) int {
	if index < 0 || c <= 0 {
		return 0
	}
	if memo[index][c] != -1 {
		return memo[index][c]
	}
	// 两种情况，
	// 1. 不将 index 装入背包
	res := bestValue(w, v, index-1, c)
	// 2. 将 index 装入背包
	if w[index] <= c {
		res = maxOf(res, v[index]+bestValue(w, v, index-1, c-w[index]))
	}
	memo[index][c] = res
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(knapsack01([]int{1, 2, 3}, []int{6, 10, 12}, 5))
}
