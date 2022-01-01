package solution0709

import "strings"

/*func isUpperCase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func toLowerCase(s string) string {
	bytes := []byte(s)
	for i, b := range bytes {
		if isUpperCase(b) {
			bytes[i] = b + ('a' - 'A')
		}
	}
	return string(bytes)
}
*/

func toLowerCase(s string) string {
	return strings.ToLower(s)
}
