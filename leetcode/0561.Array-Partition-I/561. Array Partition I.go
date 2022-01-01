// https://leetcode-cn.com/problems/array-partition-i/

package solution0561

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}

func Test() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}) == 4)
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}) == 9)
}
