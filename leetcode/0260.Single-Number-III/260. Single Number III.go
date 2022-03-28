package solution0260

func singleNumber(nums []int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	res := make([]int, 0)
	for num, count := range freq {
		if count == 1 {
			res = append(res, num)
		}
	}
	return res
}
