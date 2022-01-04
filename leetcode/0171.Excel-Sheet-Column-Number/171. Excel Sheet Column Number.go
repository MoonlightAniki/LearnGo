package solution0171

import "fmt"

func titleToNumber(columnTitle string) int {
	res := 0
	for _, b := range []byte(columnTitle) {
		res = res*26 + int(b-'A'+1)
	}
	return res
}

func Test() {
	fmt.Println(titleToNumber("A") == 1)
	fmt.Println(titleToNumber("AB") == 28)
	fmt.Println(titleToNumber("ZY") == 701)
}
