package solution0134

// 执行用时：68 ms, 在所有 Go 提交中击败了 50.17% 的用户
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	start := 0
	end := 0
	curGas := 0
	count := 0
	for start < n && count < n {
		count++
		curGas = curGas + gas[end] - cost[end]
		for curGas < 0 && start < n {
			count--
			curGas = curGas - gas[start] + cost[start]
			start++
		}
		end = (end + 1) % n
	}
	if curGas >= 0 && count == n {
		return start
	} else {
		return -1
	}
}
