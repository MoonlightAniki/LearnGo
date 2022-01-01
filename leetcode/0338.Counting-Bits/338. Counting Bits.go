// https://leetcode-cn.com/problems/counting-bits/

package solution0338

import "fmt"

/*func countBits(n int) []int {
	res := make([]int, n+1)
	for i, _ := range res {
		res[i] = -1
	}
	res[0] = 0
	for i := 1; i < len(res); i++ {
		if res[i] == -1 {
			res[i] = countBit(i)
			for j := 2 * i; j < len(res); j *= 2 {
				res[j] = res[i]
			}
		}
	}
	return res
}

func countBit(n int) int {
	res := 0
	for n > 0 {
		res += n & 1
		n = n >> 1
	}
	return res
}*/

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i < len(res); i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

func Test() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
