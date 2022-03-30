package solution0299

import "fmt"

func getHint(secret string, guess string) string {
	sMap := make(map[byte]int)
	gMap := make(map[byte]int)
	sBytes, gBytes := []byte(secret), []byte(guess)
	A, B := 0, 0
	for i := range sBytes {
		if sBytes[i] == gBytes[i] {
			A++
		} else {
			sMap[sBytes[i]]++
			gMap[gBytes[i]]++
		}
	}
	for k, v := range sMap {
		B += minOf(v, gMap[k])
	}
	return fmt.Sprintf("%dA%dB", A, B)
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
