package solution0303

type NumArray struct {
	// sums[i] 代表 nums[0,i)范围内元素的和
	// SumRange(left, right) 等于 sums(right+1) - sums(left)
	sums []int
}

func Constructor(nums []int) NumArray {
	var sums []int
	sums = append(sums, 0)
	for i, num := range nums {
		sums = append(sums, sums[i]+num)
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}
