package solution0011

import "fmt"

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.


Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1


Constraints:
n == height.length
2 <= n <= 10^5
0 <= height[i] <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		area := (r - l) * minOf(height[l], height[r])
		res = maxOf(res, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
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

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(
		maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49,
		maxArea([]int{1, 1}) == 1,
	)
}
