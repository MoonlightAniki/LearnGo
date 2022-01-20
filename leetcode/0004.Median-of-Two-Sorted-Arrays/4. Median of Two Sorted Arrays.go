package solution0004

import "fmt"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Example 2:
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


Constraints:
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 使用归并排序 merge 的思路，时间复杂度为 O(m+n)
// 执行用时：12 ms, 在所有 Go 提交中击败了 88.20% 的用户
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	// 奇数个元素，等于第 (m+n)/2 个元素，偶数个元素，等于 (m+n)/2 - 1 和 (m+n)/2 两个元素的和
	sum := 0
	for k := 0; k <= (m+n)/2; k++ {
		var num int
		if i >= m {
			num = nums2[j]
			j++
		} else if j >= n {
			num = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			num = nums1[i]
			i++
		} else {
			num = nums2[j]
			j++
		}
		if k == (m+n)/2-1 && (m+n)%2 == 0 {
			sum += num
		} else if k == (m+n)/2 {
			sum += num
		}
	}
	if (m+n)%2 == 0 {
		// 偶数个元素的情况
		return float64(sum) / 2
	} else {
		// 奇数个元素的情况
		return float64(sum)
	}
}

func Test() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}
