package solution0274

func hIndex(nums []int) int {
	n := len(nums)
	freq := make([]int, n+1)
	for _, num := range nums {
		if num > n {
			freq[n]++
		} else {
			freq[num]++
		}
	}
	total := 0
	for i := len(freq) - 1; i >= 0; i-- {
		total += freq[i]
		if total >= i {
			return i
		}
	}
	return 0
}
