package solution0494

/*
You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".
Return the number of different expressions that you can build, which evaluates to target.


Example 1:
Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3


Example 2:
Input: nums = [1], target = 1
Output: 1


Constraints:
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/

/*func findTargetSumWays(nums []int, target int) int {
	// 添加 - 号的元素和为 a, 添加 + 号的元素和为 b 则有 -a + b = target, a + b = sum
	// 则有 2 * b = target + sum
	// 转化为 01 背包问题，从 nums 中取出若干个元素，使得这些元素的和等于 (target+sum)/2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)%2 == 1 || sum+target < 0 {
		return 0
	}
	return subsetSum(nums, (sum+target)/2)
}

func subsetSum(nums []int, sum int) int {
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for j := sum; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[sum]
}*/

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	s := sum + target
	if s%2 == 1 || s < 0 {
		return 0
	}
	C := s / 2
	memo = make([][]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, C+1)
		for j := 0; j <= C; j++ {
			memo[i][j] = -1
		}
	}
	return subsetSum(nums, len(nums)-1, C)
}

var memo [][]int

func subsetSum(nums []int, index int, sum int) int {
	if index < 0 {
		if sum == 0 {
			return 1
		} else {
			return 0
		}
	}
	if sum < 0 {
		return 0
	}
	if memo[index][sum] != -1 {
		return memo[index][sum]
	}
	memo[index][sum] = subsetSum(nums, index-1, sum) + subsetSum(nums, index-1, sum-nums[index])
	return memo[index][sum]
}
