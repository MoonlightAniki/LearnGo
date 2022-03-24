package solution0187

/*
The DNA sequence is composed of a series of nucleotides abbreviated as 'A', 'C', 'G', and 'T'.

For example, "ACGAATTCCG" is a DNA sequence.
When studying DNA, it is useful to identify repeated sequences within the DNA.

Given a string s that represents a DNA sequence, return all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in any order.


Example 1:
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]


Example 2:
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]


Constraints:
1 <= s.length <= 10^5
s[i] is either 'A', 'C', 'G', or 'T'.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/repeated-dna-sequences
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findRepeatedDnaSequences(s string) []string {
	freq := make(map[string]int)
	res := make([]string, 0)
	for i := 0; i+10 <= len(s); i++ {
		sub := s[i : i+10]
		freq[sub]++
		if freq[sub] == 2 {
			res = append(res, sub)
		}
	}
	return res
}
