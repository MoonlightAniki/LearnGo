package solution0268

// sum 有越界的风险
/*func missingNumber(nums []int) int {
	sum := 0
	for i := 1; i <= len(nums); i++ {
		sum += i
	}
	for _, num := range nums {
		sum -= num
	}
	return sum
}*/

/*func missingNumber(nums []int) int {
	sum := 0
	sum += len(nums)
	for i, num := range nums {
		sum += i
		sum -= num
	}
	return sum
}*/

func missingNumber(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		res ^= i
		res ^= num
	}
	return res
}
