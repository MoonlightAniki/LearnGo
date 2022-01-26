package v3

import "fmt"

func knapsack01(w []int, v []int, C int) int {
	n := len(w)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, C+1)
	}
	for j := 0; j <= C; j++ {
		if j >= w[0] {
			dp[0][j] = v[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= C; j++ {
			// 1 不将 i 位置的物品放入背包
			maxValue := dp[(i-1)%2][j]
			// 2 将 i 位置的物品放入背包
			if w[i] <= j {
				maxValue = maxOf(maxValue, v[i]+dp[(i-1)%2][j-w[i]])
			}
			dp[i%2][j] = maxValue
		}
	}
	return dp[(n-1)%2][C]
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
