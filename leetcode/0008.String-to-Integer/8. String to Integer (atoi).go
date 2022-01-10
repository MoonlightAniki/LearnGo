package solution0008

import (
	"fmt"
	"math"
	"strings"
)

var maxInt32 = 1<<31 - 1
var minInt32 = -1 << 31

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	bytes := []byte(s)
	if len(bytes) == 0 {
		return 0
	}
	i := 0
	sign := 1
	if bytes[i] == '+' {
		sign = 1
		i++
	} else if bytes[i] == '-' {
		sign = -1
		i++
	}
	var res int64 = 0
	for ; i < len(bytes); i++ {
		if isDigit(bytes[i]) {
			res = res*10 + int64(bytes[i]-'0')
			if int64(sign)*res > int64(maxInt32) {
				return maxInt32
			} else if int64(sign)*res < int64(minInt32) {
				return minInt32
			}
		} else {
			break
		}
	}
	return sign * int(res)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func Test() {
	fmt.Println(
		myAtoi("42"),
		myAtoi("    -42"),
		myAtoi("4193 with words"),
		myAtoi(fmt.Sprintf("%d", math.MaxInt64)),
	)
}
