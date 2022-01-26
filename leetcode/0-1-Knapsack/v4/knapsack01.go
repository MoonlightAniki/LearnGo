package v4

import "fmt"

func knapsack01(w []int, v []int, C int) int {
	n := len(w)
	dp := make([]int, C+1)
	for j := 0; j <= C; j++ {
		if j >= w[0] {
			dp[j] = v[0]
		}
	}
	for i := 1; i < n; i++ {
		for j := C; j >= w[i]; j-- {
			dp[j] = maxOf(dp[j], v[i]+dp[j-w[i]])
		}
	}
	return dp[C]
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
