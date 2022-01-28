package solution0315

/*
You are given an integer array nums and you have to return a new counts array. The counts array has the property where
counts[i] is the number of smaller elements to the right of nums[i].


Example 1:
Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.


Example 2:
Input: nums = [-1]
Output: [0]


Example 3:
Input: nums = [-1,-1]
Output: [0,0]


Constraints:
1 <= nums.length <= 10^5
-104 <= nums[i] <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 时间超出限制, 时间复杂度 O(n^2)
/*func countSmaller(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}*/

// 超出时间限制，时间复杂度 O(nlogn)
type node struct {
	value int
	size  int
	left  *node
	right *node
}

func (n *node) add(value int) *node {
	if n == nil {
		return &node{value: value, size: 1}
	}
	if value <= n.value {
		n.left = n.left.add(value)
	} else {
		n.right = n.right.add(value)
	}
	n.size++
	return n
}

func (n *node) smaller(value int) int {
	if n == nil {
		return 0
	}
	if value <= n.value {
		return n.left.smaller(value)
	} else {
		res := 0
		res += 1
		if n.left != nil {
			res += n.left.size
		}
		res += n.right.smaller(value)
		return res
	}
}

func countSmaller(nums []int) []int {
	var root *node
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = root.smaller(nums[i])
		root = root.add(nums[i])
	}
	return res
}
