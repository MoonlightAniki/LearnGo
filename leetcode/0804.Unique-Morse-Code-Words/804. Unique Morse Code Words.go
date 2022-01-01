package solution0804

var morseCode = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]bool)
	for _, word := range words {
		m[stringToMorseCode(word)] = true
	}
	return len(m)
}

func stringToMorseCode(word string) string {
	res := ""
	for _, ch := range []byte(word) {
		res += morseCode[ch-'a']
	}
	return res
}
