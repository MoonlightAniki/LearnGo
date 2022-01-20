package solution0049

import (
	"fmt"
	"strconv"
)

/*
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]


Constraints:
1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 执行用时：40 ms, 在所有 Go 提交中击败了 8.27% 的用户
// 时间复杂度 O(nklogk), n为数组的长度，k为字符串的最大长度
/*func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		an := getAnagram(str)
		if m[an] == nil {
			m[an] = []string{}
		}
		m[an] = append(m[an], str)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getAnagram(s string) string {
	var ss []string
	for _, c := range s {
		ss = append(ss, string(c))
	}
	sort.Strings(ss)
	return strings.Join(ss, "")
}*/

// 执行用时：40 ms, 在所有 Go 提交中击败了 8.27% 的用户
// 时间复杂度 O(n * k)
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		key := getKey(s)
		m[key] = append(m[key], s)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getKey(s string) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[int(c-'a')]++
	}
	key := ""
	for _, count := range freq {
		key += strconv.Itoa(count) + "#"
	}
	return key
}

func Test() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
