// https://leetcode-cn.com/problems/hamming-distance/

package solution0461

import "fmt"

func hammingDistance(x int, y int) int {
	xor := x ^ y
	res := 0
	for xor > 0 {
		res += xor % 2
		xor /= 2
	}
	return res
}

func Test() {
	fmt.Println(hammingDistance(1, 4) == 2)
	fmt.Println(hammingDistance(1, 3) == 1)
}
