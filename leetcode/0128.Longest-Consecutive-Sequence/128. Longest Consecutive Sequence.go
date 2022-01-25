package solution0128

import (
	"fmt"
)

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:
Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.


Example 2:
Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9


Constraints:
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Out of Memory
/*var n = 1000000000
var freq = make([]int, 2*n+1)

func longestConsecutive(nums []int) int {
	for i := 0; i < len(freq); i++ {
		freq[i] = 0
	}
	offset := n
	for _, num := range nums {
		freq[num+offset]++
	}
	res := 0
	consecutive := 0
	for _, count := range freq {
		if count == 0 {
			consecutive = 0
		} else {
			consecutive++
			res = maxOf(res, consecutive)
		}
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

// 执行用时：32 ms, 在所有 Go 提交中击败了 97.36% 的用户
// 时间复杂度 O(nlogn)
/*func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	count := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			count = 1
		} else {
			if nums[i] == nums[i-1] {
				continue
			} else if nums[i] == nums[i-1]+1 {
				count++
			} else {
				count = 1
			}
		}
		res = maxOf(res, count)
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}*/

// 时间复杂度 O(n)
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	res := 0
	for num, _ := range set {
		if !set[num-1] {
			count := 1
			i := num + 1
			for set[i] {
				count++
				i++
			}
			res = maxOf(res, count)
		}
	}
	return res
}

func maxOf(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{0, 0, 0, 1, 1, 2, 2}))
}
