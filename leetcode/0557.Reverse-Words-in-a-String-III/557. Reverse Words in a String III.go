package solution0557

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverseWord(word)
	}
	return strings.Join(words, " ")
}

func reverseWord(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}
