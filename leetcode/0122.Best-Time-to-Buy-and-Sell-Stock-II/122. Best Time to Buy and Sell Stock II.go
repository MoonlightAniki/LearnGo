package solution0122

// 122. Best Time to Buy and Sell Stock II
/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.


Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
Total profit is 4 + 3 = 7.


Example 2:
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Total profit is 4.


Example 3:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.


Constraints:
1 <= prices.length <= 3 * 10^4
0 <= prices[i] <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 执行用时: 4 ms, 在所有 Go 提交中击败了89.87% 的用户
/* func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	peekAndValley := []int{0}
	i := 1
	for i < len(prices) {
		// 先寻找下一个高点
		for i < len(prices) && prices[i] >= prices[i-1] {
			i++
		}
		if i-1 != peekAndValley[len(peekAndValley)-1] {
			peekAndValley = append(peekAndValley, i-1)
		}
		// 寻找下一个低点
		for i < len(prices) && prices[i] <= prices[i-1] {
			i++
		}
		if i-1 != peekAndValley[len(peekAndValley)-1] {
			peekAndValley = append(peekAndValley, i-1)
		}
	}
	res := 0
	for i := 1; i < len(peekAndValley); i++ {
		if prices[peekAndValley[i]] > prices[peekAndValley[i-1]] {
			res += prices[peekAndValley[i]] - prices[peekAndValley[i-1]]
		}
	}
	return res
} */

// 执行用时: 4 ms, 在所有 Go 提交中击败了89.87% 的用户
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
