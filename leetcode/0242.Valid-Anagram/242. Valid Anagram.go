package solution0242

/*func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq1, freq2 := getFreq(s), getFreq(t)
	if len(freq1) != len(freq2) {
		return false
	}
	for k, v := range freq1 {
		if freq2[k] != v {
			return false
		}
	}
	return true
}

func getFreq(s string) map[byte]int {
	freq := make(map[byte]int)
	for _, b := range []byte(s) {
		freq[b]++
	}
	return freq
}*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := []byte(s)
	tt := []byte(t)
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[ss[i]]++
		freq[tt[i]]--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
