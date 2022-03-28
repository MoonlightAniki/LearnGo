package solution0229

// 执行用时：8 ms, 在所有 Go 提交中击败了 92.08% 的用户
func majorityElement(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res := make([]int, 0)
	for num, count := range freq {
		if count > len(nums)/3 {
			res = append(res, num)
		}
	}
	return res
}
