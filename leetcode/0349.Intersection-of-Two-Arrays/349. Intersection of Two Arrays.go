package solution0349

/*func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)

	for _, num := range nums1 {
		set[num] = true
	}

	var res []int
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			delete(set, num)
			res = append(res, num)
		}
	}
	return res
}*/

func intersection(nums1 []int, nums2 []int) []int {
	a := getInts(nums1)
	b := getInts(nums2)

	var res []int
	for num, _ := range a {
		if b[num] {
			res = append(res, num)
		}
	}
	return res
}

func getInts(nums []int) map[int]bool {
	res := make(map[int]bool)
	for _, num := range nums {
		res[num] = true
	}
	return res
}
