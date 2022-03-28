package solution0275

func hIndex(nums []int) int {
	i := len(nums) - 1
	h := 0
	for i >= 0 && h < nums[i] {
		i--
		h++
	}
	return h
}
