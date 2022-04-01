package solution0386

// 执行用时：88 ms, 在所有 Go 提交中击败了 5.76% 的用户
/*func lexicalOrder(n int) []int {

	res := make([]int, n)
	for i := 1; i <= n; i++ {
		res[i-1] = i
	}

	sort.Slice(res, func(i, j int) bool {
		return strconv.Itoa(res[i]) < strconv.Itoa(res[j])
	})

	return res
}*/

// 执行用时： 8 ms, 在所有 Go 提交中击败了 87.96% 的用户
func lexicalOrder(n int) []int {
	// 10叉树前序遍历
	res := make([]int, 0, n)
	for i := 1; i <= 9; i++ {
		dfs(i, n, &res)
	}
	return res
}

func dfs(cur int, n int, res *[]int) {
	if cur > n {
		return
	}
	*res = append(*res, cur)
	for i := 0; i <= 9; i++ {
		dfs(cur*10+i, n, res)
	}
}
