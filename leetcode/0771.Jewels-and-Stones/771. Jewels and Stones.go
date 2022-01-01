package solution0771

import "fmt"

// https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(jewels string, stones string) int {
	jewelMap := make(map[rune]int)
	for _, ch := range jewels {
		jewelMap[ch]++
	}

	res := 0
	for _, ch := range stones {
		if jewelMap[ch] > 0 {
			res++
		}
	}
	return res
}

func Test() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb") == 3)
}
