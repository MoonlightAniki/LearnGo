package solution1002

import "strings"

/*func commonChars(words []string) []string {
	var freqs []map[byte]int
	for _, word := range words {
		freqs = append(freqs, getFreq(word))
	}
	var res []string
	for b := 'a'; b <= 'z'; b++ {
		minSeq := math.MaxInt
		for _, freq := range freqs {
			minSeq = minOf(minSeq, freq[byte(b)])
		}
		for i := 0; i < minSeq; i++ {
			res = append(res, string(b))
		}
	}
	return res
}

func getFreq(word string) map[byte]int {
	res := make(map[byte]int)
	for _, b := range []byte(word) {
		res[b]++
	}
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}*/

func commonChars(words []string) []string {
	var res []string
	if len(words) == 0 {
		return res
	}
	for b := 'a'; b <= 'z'; b++ {
		sub := string(b)
		count := strings.Count(words[0], sub)
		for i := 1; i < len(words); i++ {
			count = minOf(count, strings.Count(words[i], sub))
		}
		for count > 0 {
			res = append(res, sub)
			count--
		}
	}
	return res
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
