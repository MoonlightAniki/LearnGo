package solution0042

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.


Example 1:
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
Input: height = [4,2,0,3,2,5]
Output: 9


Constraints:
n == height.length
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 暴力解法，时间复杂度 O(n^2)
/*func trap(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		lmax, rmax := height[i], height[i]
		for l := i - 1; l >= 0; l-- {
			lmax = maxOf(lmax, height[l])
		}
		for r := i + 1; r < len(height); r++ {
			rmax = maxOf(rmax, height[r])
		}
		res += minOf(lmax, rmax) - height[i]
	}
	return res
}*/

// 记忆化搜索, 时间复杂度 O(n)
/*func trap(height []int) int {
	ldp := make([]int, len(height))
	getL(len(height)-1, height, ldp)

	rdp := make([]int, len(height))
	getR(0, height, rdp)

	res := 0
	for i := 0; i < len(height); i++ {
		res += minOf(ldp[i], rdp[i]) - height[i]
	}
	return res
}

func getL(index int, height []int, ldp []int) int {
	if index == 0 {
		ldp[index] = height[index]
	} else {
		ldp[index] = maxOf(height[index], getL(index-1, height, ldp))
	}
	return ldp[index]
}

func getR(index int, height []int, rdp []int) int {
	if index == len(height)-1 {
		rdp[index] = height[index]
	} else {
		rdp[index] = maxOf(height[index], getR(index+1, height, rdp))
	}
	return rdp[index]
}*/

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	ldp := make([]int, len(height))
	ldp[0] = height[0]
	for i := 1; i < len(height); i++ {
		ldp[i] = maxOf(height[i], ldp[i-1])
	}

	rdp := make([]int, len(height))
	rdp[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rdp[i] = maxOf(height[i], rdp[i+1])
	}

	res := 0
	for i := 0; i < len(height); i++ {
		res += minOf(ldp[i], rdp[i]) - height[i]
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
