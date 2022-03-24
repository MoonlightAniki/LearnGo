package solution0179

import (
	"sort"
	"strconv"
)

/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.



Example 1:
Input: nums = [10,2]
Output: "210"


Example 2:
Input: nums = [3,30,34,5,9]
Output: "9534330"


Constraints:
1 <= nums.length <= 100
0 <= nums[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 执行用时：4 ms, 在所有 Go 提交中击败了 32.32% 的用户
/*func largestNumber(nums []int) string {
	numbers := Numbers{nums}
	sort.Sort(numbers)
	if numbers.data[0] == 0 {
		return "0"
	}
	res := ""
	for _, num := range numbers.data {
		res += strconv.Itoa(num)
	}
	return res
}

type Numbers struct {
	data []int
}

func (n Numbers) Len() int {
	return len(n.data)
}

func (n Numbers) Less(i, j int) bool {
	s1, s2 := strconv.Itoa(n.data[i]), strconv.Itoa(n.data[j])
	a, _ := strconv.Atoi(s1 + s2)
	b, _ := strconv.Atoi(s2 + s1)
	return a > b
}

func (n Numbers) Swap(i, j int) {
	n.data[i], n.data[j] = n.data[j], n.data[i]
}*/

// 执行用时：4 ms, 在所有 Go 提交中击败了 32.32% 的用户
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		s1, s2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		sum1, _ := strconv.Atoi(s1 + s2)
		sum2, _ := strconv.Atoi(s2 + s1)
		return sum1 > sum2
	})
	if nums[0] == 0 {
		return "0"
	}
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}
